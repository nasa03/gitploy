package store

import (
	"context"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/chatuser"
	"github.com/hanjunlee/gitploy/ent/user"
)

func (s *Store) FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error) {
	return s.c.ChatUser.Get(ctx, id)
}

func (s *Store) FindUserWithChatUserByChatUserID(ctx context.Context, id string) (*ent.User, error) {
	return s.c.User.
		Query().
		Where(
			user.HasChatUserWith(chatuser.IDEQ(id)),
		).
		WithChatUser().
		First(ctx)
}

func (s *Store) CreateChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error) {
	return s.c.ChatUser.
		Create().
		SetID(cu.ID).
		SetToken(cu.Token).
		SetBotToken(cu.BotToken).
		SetRefresh(cu.Refresh).
		SetExpiry(cu.Expiry).
		SetUserID(u.ID).
		Save(ctx)
}

func (s *Store) UpdateChatUser(ctx context.Context, u *ent.User, cu *ent.ChatUser) (*ent.ChatUser, error) {
	return s.c.ChatUser.
		UpdateOneID(cu.ID).
		SetToken(cu.Token).
		SetBotToken(cu.BotToken).
		SetRefresh(cu.Refresh).
		SetExpiry(cu.Expiry).
		SetUserID(u.ID).
		Save(ctx)
}