package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TenantUser merupakan tabel pivot yang menghubungkan User ke Tenant
// beserta role global pada tenant tersebut.
type TenantUser struct {
	ent.Schema
}

// Annotations mengatur tabel ini berada di skema public.
func (TenantUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Schema("public"),
	}
}

// Fields mendefinisikan kolom-kolom tabel tenant_users.
func (TenantUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable().
			Comment("Primary key UUID"),

		field.UUID("tenant_id", uuid.UUID{}).
			Comment("Foreign key ke tabel tenants"),

		field.UUID("user_id", uuid.UUID{}).
			Comment("Foreign key ke tabel users"),

		field.String("role").
			NotEmpty().
			Comment("Role global pada tenant, terintegrasi dengan Casbin"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Waktu pembuatan record"),
	}
}

// Edges mendefinisikan relasi dengan entitas lain.
func (TenantUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("tenant_users").
			Field("tenant_id").
			Required().
			Unique(),
		edge.From("user", User.Type).
			Ref("tenant_users").
			Field("user_id").
			Required().
			Unique(),
	}
}
