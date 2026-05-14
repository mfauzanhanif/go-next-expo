package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Department menyimpan data departemen/bagian dalam institusi.
// Tabel ini berada di skema tenant (di-generate per institusi).
type Department struct {
	ent.Schema
}

// Annotations mengatur skema default. Akan di-override secara dinamis per tenant.
func (Department) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Schema("tenant"),
	}
}

// Fields mendefinisikan kolom-kolom tabel departments.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable().
			Comment("Primary key UUID"),

		field.String("name").
			NotEmpty().
			Comment("Nama departemen: Kepesantrenan, Kurikulum, Keuangan"),

		field.Text("description").
			Optional().
			Comment("Deskripsi departemen"),
	}
}

// Edges mendefinisikan relasi dengan entitas lain.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("positions", GtkPosition.Type),
	}
}
