// Code generated by entc, DO NOT EDIT.

package chatuser

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hanjunlee/gitploy/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// Refresh applies equality check predicate on the "refresh" field. It's identical to RefreshEQ.
func Refresh(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefresh), v))
	})
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// BotToken applies equality check predicate on the "bot_token" field. It's identical to BotTokenEQ.
func BotToken(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBotToken), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToken), v))
	})
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldToken), v...))
	})
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldToken), v...))
	})
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToken), v))
	})
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToken), v))
	})
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToken), v))
	})
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToken), v))
	})
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldToken), v))
	})
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldToken), v))
	})
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldToken), v))
	})
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldToken), v))
	})
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldToken), v))
	})
}

// RefreshEQ applies the EQ predicate on the "refresh" field.
func RefreshEQ(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRefresh), v))
	})
}

// RefreshNEQ applies the NEQ predicate on the "refresh" field.
func RefreshNEQ(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRefresh), v))
	})
}

// RefreshIn applies the In predicate on the "refresh" field.
func RefreshIn(vs ...string) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRefresh), v...))
	})
}

// RefreshNotIn applies the NotIn predicate on the "refresh" field.
func RefreshNotIn(vs ...string) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRefresh), v...))
	})
}

// RefreshGT applies the GT predicate on the "refresh" field.
func RefreshGT(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRefresh), v))
	})
}

// RefreshGTE applies the GTE predicate on the "refresh" field.
func RefreshGTE(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRefresh), v))
	})
}

// RefreshLT applies the LT predicate on the "refresh" field.
func RefreshLT(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRefresh), v))
	})
}

// RefreshLTE applies the LTE predicate on the "refresh" field.
func RefreshLTE(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRefresh), v))
	})
}

// RefreshContains applies the Contains predicate on the "refresh" field.
func RefreshContains(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRefresh), v))
	})
}

// RefreshHasPrefix applies the HasPrefix predicate on the "refresh" field.
func RefreshHasPrefix(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRefresh), v))
	})
}

// RefreshHasSuffix applies the HasSuffix predicate on the "refresh" field.
func RefreshHasSuffix(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRefresh), v))
	})
}

// RefreshEqualFold applies the EqualFold predicate on the "refresh" field.
func RefreshEqualFold(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRefresh), v))
	})
}

// RefreshContainsFold applies the ContainsFold predicate on the "refresh" field.
func RefreshContainsFold(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRefresh), v))
	})
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiry), v))
	})
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpiry), v...))
	})
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpiry), v...))
	})
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiry), v))
	})
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiry), v))
	})
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiry), v))
	})
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiry), v))
	})
}

// BotTokenEQ applies the EQ predicate on the "bot_token" field.
func BotTokenEQ(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBotToken), v))
	})
}

// BotTokenNEQ applies the NEQ predicate on the "bot_token" field.
func BotTokenNEQ(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBotToken), v))
	})
}

// BotTokenIn applies the In predicate on the "bot_token" field.
func BotTokenIn(vs ...string) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBotToken), v...))
	})
}

// BotTokenNotIn applies the NotIn predicate on the "bot_token" field.
func BotTokenNotIn(vs ...string) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBotToken), v...))
	})
}

// BotTokenGT applies the GT predicate on the "bot_token" field.
func BotTokenGT(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBotToken), v))
	})
}

// BotTokenGTE applies the GTE predicate on the "bot_token" field.
func BotTokenGTE(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBotToken), v))
	})
}

// BotTokenLT applies the LT predicate on the "bot_token" field.
func BotTokenLT(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBotToken), v))
	})
}

// BotTokenLTE applies the LTE predicate on the "bot_token" field.
func BotTokenLTE(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBotToken), v))
	})
}

// BotTokenContains applies the Contains predicate on the "bot_token" field.
func BotTokenContains(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBotToken), v))
	})
}

// BotTokenHasPrefix applies the HasPrefix predicate on the "bot_token" field.
func BotTokenHasPrefix(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBotToken), v))
	})
}

// BotTokenHasSuffix applies the HasSuffix predicate on the "bot_token" field.
func BotTokenHasSuffix(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBotToken), v))
	})
}

// BotTokenEqualFold applies the EqualFold predicate on the "bot_token" field.
func BotTokenEqualFold(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBotToken), v))
	})
}

// BotTokenContainsFold applies the ContainsFold predicate on the "bot_token" field.
func BotTokenContainsFold(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBotToken), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.ChatUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ChatUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserID), v))
	})
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserID), v))
	})
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserID), v))
	})
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserID), v))
	})
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserID), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChatUser) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChatUser) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ChatUser) predicate.ChatUser {
	return predicate.ChatUser(func(s *sql.Selector) {
		p(s.Not())
	})
}
