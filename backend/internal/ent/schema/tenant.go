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

// Tenant menyimpan data institusi/organisasi penyewa sistem.
// Setiap tenant memiliki skema database terpisah (logical separation).
type Tenant struct {
	ent.Schema
}

// Annotations mengatur tabel ini berada di skema public.
func (Tenant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Schema("public"),
	}
}

// Fields mendefinisikan kolom-kolom tabel tenants.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable().
			Comment("Primary key UUID"),

		field.String("name").
			NotEmpty().
			Comment("Nama institusi, misal: Pondok Pesantren Dar Al Tauhid"),

		field.String("schema_name").
			Unique().
			NotEmpty().
			Comment("Nama skema di PostgreSQL, misal: tenant_ppdt"),

		field.String("domain").
			Unique().
			Optional().
			Nillable().
			Comment("Custom domain jika ada"),

		field.Bool("is_active").
			Default(true).
			Comment("Status aktif tenant"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Waktu pembuatan record"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Waktu terakhir update record"),
	}
}

// Edges mendefinisikan relasi dengan entitas lain.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenant_users", TenantUser.Type),
	}
}
