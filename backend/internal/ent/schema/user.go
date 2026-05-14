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

// User menyimpan data pengguna sentral yang terkoneksi dengan Zitadel.
type User struct {
	ent.Schema
}

// Annotations mengatur tabel ini berada di skema public.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Schema("public"),
	}
}

// Fields mendefinisikan kolom-kolom tabel users.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable().
			Comment("Primary key UUID, terkoneksi dengan ID dari Zitadel"),

		field.String("email").
			Unique().
			NotEmpty().
			Comment("Alamat email pengguna"),

		field.String("phone_number").
			Unique().
			Optional().
			Nillable().
			Comment("Nomor telepon pengguna"),

		field.String("full_name").
			NotEmpty().
			Comment("Nama lengkap pengguna"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Waktu pembuatan record"),
	}
}

// Edges mendefinisikan relasi dengan entitas lain.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenant_users", TenantUser.Type),
	}
}
