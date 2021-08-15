package store

import (
	"context"
	"testing"
	"time"

	"github.com/hanjunlee/gitploy/ent"
	"github.com/hanjunlee/gitploy/ent/enttest"
	"github.com/hanjunlee/gitploy/ent/migrate"
)

func TestStore_ListUsers(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithMigrateOptions(migrate.WithForeignKeys(false)),
	)
	defer client.Close()

	ctx := context.Background()

	u1 := client.User.
		Create().
		SetID("1").
		SetLogin("banana").
		SetToken("").
		SetRefresh("").
		SetExpiry(time.Time{}).
		SaveX(ctx)

	u2 := client.User.
		Create().
		SetID("2").
		SetLogin("apple").
		SetToken("").
		SetRefresh("").
		SetExpiry(time.Time{}).
		SaveX(ctx)

	t.Run("Returns users in login-sorted manager.", func(t *testing.T) {
		s := NewStore(client)

		var (
			us  []*ent.User
			err error
		)

		if us, err = s.ListUsers(ctx, "", 1, 30); err != nil {
			t.Fatalf("ListUsers returns an error: %s", err)
		}

		expected := []*ent.User{u2, u1}
		if !(len(us) == 2 && us[0].ID == expected[0].ID) {
			t.Fatalf("ListUsers = %v, wanted %v", us, expected)
		}
	})

	t.Run("Returns user filtered by login.", func(t *testing.T) {
		s := NewStore(client)

		var (
			us  []*ent.User
			err error
		)

		if us, err = s.ListUsers(ctx, "pple", 1, 30); err != nil {
			t.Fatalf("ListUsers returns an error: %s", err)
		}

		expected := u2
		if !(len(us) == 1 && us[0].ID == expected.ID) {
			t.Fatalf("ListUsers = %v, wanted %v", us, expected)
		}
	})
}
