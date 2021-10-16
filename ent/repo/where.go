// Code generated by entc, DO NOT EDIT.

package repo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gitploy-io/gitploy/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNamespace), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// ConfigPath applies equality check predicate on the "config_path" field. It's identical to ConfigPathEQ.
func ConfigPath(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfigPath), v))
	})
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// WebhookID applies equality check predicate on the "webhook_id" field. It's identical to WebhookIDEQ.
func WebhookID(v int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWebhookID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// LatestDeployedAt applies equality check predicate on the "latest_deployed_at" field. It's identical to LatestDeployedAtEQ.
func LatestDeployedAt(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatestDeployedAt), v))
	})
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNamespace), v))
	})
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNamespace), v))
	})
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNamespace), v...))
	})
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNamespace), v...))
	})
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNamespace), v))
	})
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNamespace), v))
	})
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNamespace), v))
	})
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNamespace), v))
	})
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNamespace), v))
	})
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNamespace), v))
	})
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNamespace), v))
	})
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNamespace), v))
	})
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNamespace), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// ConfigPathEQ applies the EQ predicate on the "config_path" field.
func ConfigPathEQ(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfigPath), v))
	})
}

// ConfigPathNEQ applies the NEQ predicate on the "config_path" field.
func ConfigPathNEQ(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfigPath), v))
	})
}

// ConfigPathIn applies the In predicate on the "config_path" field.
func ConfigPathIn(vs ...string) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConfigPath), v...))
	})
}

// ConfigPathNotIn applies the NotIn predicate on the "config_path" field.
func ConfigPathNotIn(vs ...string) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConfigPath), v...))
	})
}

// ConfigPathGT applies the GT predicate on the "config_path" field.
func ConfigPathGT(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConfigPath), v))
	})
}

// ConfigPathGTE applies the GTE predicate on the "config_path" field.
func ConfigPathGTE(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConfigPath), v))
	})
}

// ConfigPathLT applies the LT predicate on the "config_path" field.
func ConfigPathLT(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConfigPath), v))
	})
}

// ConfigPathLTE applies the LTE predicate on the "config_path" field.
func ConfigPathLTE(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConfigPath), v))
	})
}

// ConfigPathContains applies the Contains predicate on the "config_path" field.
func ConfigPathContains(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldConfigPath), v))
	})
}

// ConfigPathHasPrefix applies the HasPrefix predicate on the "config_path" field.
func ConfigPathHasPrefix(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldConfigPath), v))
	})
}

// ConfigPathHasSuffix applies the HasSuffix predicate on the "config_path" field.
func ConfigPathHasSuffix(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldConfigPath), v))
	})
}

// ConfigPathEqualFold applies the EqualFold predicate on the "config_path" field.
func ConfigPathEqualFold(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldConfigPath), v))
	})
}

// ConfigPathContainsFold applies the ContainsFold predicate on the "config_path" field.
func ConfigPathContainsFold(v string) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldConfigPath), v))
	})
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActive), v))
	})
}

// WebhookIDEQ applies the EQ predicate on the "webhook_id" field.
func WebhookIDEQ(v int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWebhookID), v))
	})
}

// WebhookIDNEQ applies the NEQ predicate on the "webhook_id" field.
func WebhookIDNEQ(v int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWebhookID), v))
	})
}

// WebhookIDIn applies the In predicate on the "webhook_id" field.
func WebhookIDIn(vs ...int64) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWebhookID), v...))
	})
}

// WebhookIDNotIn applies the NotIn predicate on the "webhook_id" field.
func WebhookIDNotIn(vs ...int64) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWebhookID), v...))
	})
}

// WebhookIDGT applies the GT predicate on the "webhook_id" field.
func WebhookIDGT(v int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWebhookID), v))
	})
}

// WebhookIDGTE applies the GTE predicate on the "webhook_id" field.
func WebhookIDGTE(v int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWebhookID), v))
	})
}

// WebhookIDLT applies the LT predicate on the "webhook_id" field.
func WebhookIDLT(v int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWebhookID), v))
	})
}

// WebhookIDLTE applies the LTE predicate on the "webhook_id" field.
func WebhookIDLTE(v int64) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWebhookID), v))
	})
}

// WebhookIDIsNil applies the IsNil predicate on the "webhook_id" field.
func WebhookIDIsNil() predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWebhookID)))
	})
}

// WebhookIDNotNil applies the NotNil predicate on the "webhook_id" field.
func WebhookIDNotNil() predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWebhookID)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// LatestDeployedAtEQ applies the EQ predicate on the "latest_deployed_at" field.
func LatestDeployedAtEQ(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatestDeployedAt), v))
	})
}

// LatestDeployedAtNEQ applies the NEQ predicate on the "latest_deployed_at" field.
func LatestDeployedAtNEQ(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLatestDeployedAt), v))
	})
}

// LatestDeployedAtIn applies the In predicate on the "latest_deployed_at" field.
func LatestDeployedAtIn(vs ...time.Time) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLatestDeployedAt), v...))
	})
}

// LatestDeployedAtNotIn applies the NotIn predicate on the "latest_deployed_at" field.
func LatestDeployedAtNotIn(vs ...time.Time) predicate.Repo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLatestDeployedAt), v...))
	})
}

// LatestDeployedAtGT applies the GT predicate on the "latest_deployed_at" field.
func LatestDeployedAtGT(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLatestDeployedAt), v))
	})
}

// LatestDeployedAtGTE applies the GTE predicate on the "latest_deployed_at" field.
func LatestDeployedAtGTE(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLatestDeployedAt), v))
	})
}

// LatestDeployedAtLT applies the LT predicate on the "latest_deployed_at" field.
func LatestDeployedAtLT(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLatestDeployedAt), v))
	})
}

// LatestDeployedAtLTE applies the LTE predicate on the "latest_deployed_at" field.
func LatestDeployedAtLTE(v time.Time) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLatestDeployedAt), v))
	})
}

// LatestDeployedAtIsNil applies the IsNil predicate on the "latest_deployed_at" field.
func LatestDeployedAtIsNil() predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLatestDeployedAt)))
	})
}

// LatestDeployedAtNotNil applies the NotNil predicate on the "latest_deployed_at" field.
func LatestDeployedAtNotNil() predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLatestDeployedAt)))
	})
}

// HasPerms applies the HasEdge predicate on the "perms" edge.
func HasPerms() predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PermsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PermsTable, PermsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermsWith applies the HasEdge predicate on the "perms" edge with a given conditions (other predicates).
func HasPermsWith(preds ...predicate.Perm) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PermsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PermsTable, PermsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeployments applies the HasEdge predicate on the "deployments" edge.
func HasDeployments() predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DeploymentsTable, DeploymentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentsWith applies the HasEdge predicate on the "deployments" edge with a given conditions (other predicates).
func HasDeploymentsWith(preds ...predicate.Deployment) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DeploymentsTable, DeploymentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCallback applies the HasEdge predicate on the "callback" edge.
func HasCallback() predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CallbackTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CallbackTable, CallbackColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCallbackWith applies the HasEdge predicate on the "callback" edge with a given conditions (other predicates).
func HasCallbackWith(preds ...predicate.Callback) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CallbackInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CallbackTable, CallbackColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLocks applies the HasEdge predicate on the "locks" edge.
func HasLocks() predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocksTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LocksTable, LocksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocksWith applies the HasEdge predicate on the "locks" edge with a given conditions (other predicates).
func HasLocksWith(preds ...predicate.Lock) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocksInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LocksTable, LocksColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeploymentStatistics applies the HasEdge predicate on the "deployment_statistics" edge.
func HasDeploymentStatistics() predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentStatisticsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DeploymentStatisticsTable, DeploymentStatisticsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentStatisticsWith applies the HasEdge predicate on the "deployment_statistics" edge with a given conditions (other predicates).
func HasDeploymentStatisticsWith(preds ...predicate.DeploymentStatistics) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentStatisticsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DeploymentStatisticsTable, DeploymentStatisticsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Repo) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Repo) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
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
func Not(p predicate.Repo) predicate.Repo {
	return predicate.Repo(func(s *sql.Selector) {
		p(s.Not())
	})
}
