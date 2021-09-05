package store

import (
	"context"
	"testing"
	"time"

	"github.com/gitploy-io/gitploy/ent"
	"github.com/gitploy-io/gitploy/ent/approval"
	"github.com/gitploy-io/gitploy/ent/enttest"
	"github.com/gitploy-io/gitploy/ent/migrate"
)

func TestStore_SearchApprovals(t *testing.T) {
	client := enttest.Open(t,
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	const (
		u1 = "1"
		d1 = 1
		d2 = 2
		d3 = 3
	)

	client.Approval.
		Create().
		SetStatus(approval.StatusApproved).
		SetUserID(u1).
		SetDeploymentID(d1).
		SaveX(ctx)

	client.Approval.
		Create().
		SetStatus(approval.StatusPending).
		SetUserID(u1).
		SetDeploymentID(d2).
		SaveX(ctx)

	client.Approval.
		Create().
		SetStatus(approval.StatusPending).
		SetUserID(u1).
		SetDeploymentID(d3).
		SaveX(ctx)

	t.Run("u1 searchs requested approvals of the deployment.", func(t *testing.T) {
		const (
			owned   = false
			page    = 1
			perPage = 30
		)

		store := NewStore(client)

		res, err := store.SearchApprovals(ctx,
			&ent.User{ID: u1},
			[]approval.Status{},
			time.Now().Add(-time.Minute),
			time.Now(),
			page,
			perPage,
		)
		if err != nil {
			t.Fatalf("SearchApprovals return an error: %s", err)
			t.FailNow()
		}

		expected := 3
		if len(res) != expected {
			t.Fatalf("SearchApprovals = %v, wanted %v", res, expected)
			t.FailNow()
		}
	})

}
