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

// Gtk menyimpan data Guru dan Tenaga Kependidikan (GTK).
// Tabel ini berada di skema tenant (di-generate per institusi).
type Gtk struct {
	ent.Schema
}

// Annotations mengatur skema default. Akan di-override secara dinamis per tenant.
func (Gtk) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Schema("tenant"),
	}
}

// Fields mendefinisikan kolom-kolom tabel gtk.
func (Gtk) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable().
			Comment("Primary key UUID"),

		field.UUID("user_id", uuid.UUID{}).
			Unique().
			Optional().
			Nillable().
			Comment("Relasi ke public.users jika GTK memiliki akses login sistem"),

		field.String("nik").
			Unique().
			NotEmpty().
			MaxLen(16).
			Comment("Nomor Induk Kependudukan"),

		field.String("nuptk").
			Optional().
			Nillable().
			MaxLen(16).
			Comment("Opsional, mengakomodasi tenaga pendidik di lingkungan non-formal"),

		field.String("full_name").
			NotEmpty().
			Comment("Nama lengkap GTK"),

		field.String("gender").
			NotEmpty().
			Comment("Jenis kelamin"),

		field.String("birth_place").
			NotEmpty().
			Comment("Tempat lahir"),

		field.Time("birth_date").
			Comment("Tanggal lahir"),

		field.String("employment_status").
			NotEmpty().
			Comment("Status kepegawaian: Tetap, Honorer, Pengabdian"),

		field.Bool("is_active").
			Default(true).
			Comment("Status aktif GTK"),

		field.Time("joined_at").
			Comment("Tanggal bergabung"),

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
func (Gtk) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("positions", GtkPosition.Type),
	}
}
