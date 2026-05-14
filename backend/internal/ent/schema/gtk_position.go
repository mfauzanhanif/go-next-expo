package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GtkPosition menyimpan data jabatan/penugasan GTK di departemen tertentu.
// Satu GTK bisa memiliki lebih dari satu jabatan (rangkap tugas).
// Tabel ini berada di skema tenant (di-generate per institusi).
type GtkPosition struct {
	ent.Schema
}

// Annotations mengatur skema default. Akan di-override secara dinamis per tenant.
func (GtkPosition) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Schema("tenant"),
	}
}

// Fields mendefinisikan kolom-kolom tabel gtk_positions.
func (GtkPosition) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable().
			Comment("Primary key UUID"),

		field.UUID("gtk_id", uuid.UUID{}).
			Comment("Foreign key ke tabel gtk"),

		field.UUID("department_id", uuid.UUID{}).
			Comment("Foreign key ke tabel departments"),

		field.String("title").
			NotEmpty().
			Comment("Jabatan: Kepala Pondok, Operator, Musyrif"),

		field.Bool("is_primary").
			Default(false).
			Comment("Menandakan jabatan utama jika GTK merangkap tugas"),

		field.Time("start_date").
			Comment("Tanggal mulai menjabat"),

		field.Time("end_date").
			Optional().
			Nillable().
			Comment("Tanggal selesai menjabat (null jika masih aktif)"),
	}
}

// Edges mendefinisikan relasi dengan entitas lain.
func (GtkPosition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gtk", Gtk.Type).
			Ref("positions").
			Field("gtk_id").
			Required().
			Unique(),
		edge.From("department", Department.Type).
			Ref("positions").
			Field("department_id").
			Required().
			Unique(),
	}
}
