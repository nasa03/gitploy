package web

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
)

type (
	Interactor interface {
		SaveSCMUser(ctx context.Context, u *ent.User) (*ent.User, error)
		// Fetch the user information from SCM.
		// It has the id, login, avatar and so on.
		GetSCMUserByToken(ctx context.Context, token string) (*ent.User, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		SaveChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error)
	}
)
