# BACKEND STRUCTURE

/backend
├── /cmd
│   ├── /api                # Entry point utama untuk menjalankan HTTP Server (Echo v5)
│   └── /worker             # Entry point terpisah untuk background jobs (Asynq)
├── /internal               # Kode spesifik aplikasi (tidak bisa di-import oleh project eksternal)
│   ├── /config             # Setup dan inisiasi konfigurasi (Env, DB, Redis, MinIO)
│   ├── /ent                # Skema dan kode ORM hasil generate dari Ent (graph-based)
│   ├── /migrations         # File migrasi database deklaratif yang dikelola oleh Atlas
│   ├── /handlers           # Layer Delivery/Controller: Handler HTTP spesifik untuk Echo
│   ├── /middlewares        # Middleware Echo (Integrasi Zitadel Auth, Casbin AuthZ, Slog)
│   ├── /services           # Layer Business Logic: Aturan bisnis inti dan multi-tenant logic
│   ├── /repositories       # Layer Data: Abstraksi query database menggunakan Ent
│   └── /jobs               # Handler dan logika untuk antrean tugas Asynq (tagihan, notifikasi)
├── /modules   
├── /pkg                    # Modul utilitas internal yang bisa dibagikan lintas service
│   ├── /logger             # Setup Slog / Zerolog
│   ├── /validator          # Setup Go-Playground Validator khusus
│   └── /response           # Standardisasi format response JSON/Error aplikasi
├── /docs                   # Hasil generate otomatis dari Swaggo (OpenAPI/Swagger)
├── .env.example            # Contoh variabel lingkungan (Environment Variables)
├── go.mod                  # Definisi module dan dependensi Go
├── go.sum                  # Checksum dependensi Go
├── Makefile                # Command lokal (contoh: make run, make migrate, make swagger)
└── Dockerfile              # Instruksi build container khusus layanan backend