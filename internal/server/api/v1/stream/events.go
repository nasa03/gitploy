// Copyright 2021 Gitploy.IO Inc. All rights reserved.
// Use of this source code is governed by the Gitploy Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package stream

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/event"
	gb "github.com/gitploy-io/gitploy/internal/server/global"
	"github.com/gitploy-io/gitploy/pkg/e"
)

// GetEvents streams events of deployment, or review.
func (s *Stream) GetEvents(c *gin.Context) {
	ctx := c.Request.Context()

	v, _ := c.Get(gb.KeyUser)
	u, _ := v.(*ent.User)

	debugID := randstr()

	events := make(chan *ent.Event, 10)

	// Subscribe events
	// it'll unsubscribe after the connection is closed.
	sub := func(e *ent.Event) {

		// Deleted type is always propagated to all.
		if e.Type == event.TypeDeleted {
			events <- e
			return
		}

		if ok, err := s.hasPermForEvent(ctx, u, e); err != nil {
			s.log.Error("It has failed to check the perm.", zap.Error(err))
			return
		} else if !ok {
			s.log.Debug("Skip the event. The user has not the perm.")
			return
		}

		events <- e
	}
	if err := s.i.SubscribeEvent(sub); err != nil {
		s.log.Check(gb.GetZapLogLevel(err), "Failed to subscribe notification events").Write(zap.Error(err))
		gb.ResponseWithError(c, err)
		return
	}

	defer func() {
		if err := s.i.UnsubscribeEvent(sub); err != nil {
			s.log.Error("failed to unsubscribe notification events.")
		}

		close(events)
	}()

	w := c.Writer

L:
	for {
		select {
		case <-w.CloseNotify():
			break L
		case <-time.After(time.Hour):
			break L
		case <-time.After(time.Second * 30):
			c.Render(-1, sse.Event{
				Event: "ping",
				Data:  "ping",
			})
			w.Flush()
		case e := <-events:
			c.Render(-1, sse.Event{
				Event: "event",
				Data:  e,
			})
			w.Flush()
			s.log.Debug("server sent event.", zap.Int("event_id", e.ID), zap.String("debug_id", debugID))
		}
	}
}

// hasPermForEvent checks the user has permission for the event.
func (s *Stream) hasPermForEvent(ctx context.Context, u *ent.User, evt *ent.Event) (bool, error) {
	if evt.Kind == event.KindDeployment {
		d, err := s.i.FindDeploymentByID(ctx, evt.DeploymentID)
		if err != nil {
			return false, err
		}

		if _, err = s.i.FindPermOfRepo(ctx, d.Edges.Repo, u); e.HasErrorCode(err, e.ErrorCodeEntityNotFound) {
			return false, nil
		} else if err != nil {
			return false, err
		}

		return true, nil
	}

	if evt.Kind == event.KindReview {
		rv, err := s.i.FindReviewByID(ctx, evt.ReviewID)
		if err != nil {
			return false, err
		}

		d, err := s.i.FindDeploymentByID(ctx, rv.DeploymentID)
		if err != nil {
			return false, err
		}

		if _, err = s.i.FindPermOfRepo(ctx, d.Edges.Repo, u); e.HasErrorCode(err, e.ErrorCodeEntityNotFound) {
			return false, nil
		} else if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, fmt.Errorf("The type of event is not \"deployment\" or \"review\".")
}

func randstr() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 4)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
