# =============================================================================
# Atlas Configuration — Database Migration
# Menggunakan ent:// URL scheme untuk membaca skema dari Ent
# =============================================================================

env "local" {
  # Mengambil skema langsung dari Ent menggunakan ent:// URL scheme
  src = "ent://internal/ent/schema"
  
  # Dev database URL digunakan oleh Atlas untuk memvalidasi dan menghitung diff
  # Menggunakan docker:// agar Atlas otomatis membuat container sementara
  dev = "docker://postgres/18/dev?search_path=public"
  
  # Direktori tempat menyimpan file migrasi
  migration {
    dir = "file://internal/migrations"
  }
  
  # Konfigurasi format
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
