package interactor

import (
	"context"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/vo"
)

type (
	Store interface {
		FindUser() (*ent.User, error)
		FindUserByHash(ctx context.Context, hash string) (*ent.User, error)
		FindRepoByNamespaceName(ctx context.Context, u *ent.User, namespace, name string) (*ent.Repo, error)
		SyncPerm(ctx context.Context, user *ent.User, perm *ent.Perm, sync time.Time) error
	}

	SCM interface {
		GetConfig(ctx context.Context, u *ent.User, r *ent.Repo) (*vo.Config, error)
		GetAllPermsWithRepo(ctx context.Context, token string) ([]*ent.Perm, error)
	}
)
