# INFRA STRUCTURE

/ infra
├── /caddy                      # Reverse Proxy & SSL/TLS Management
│   ├── Caddyfile               # Aturan routing utama (proxy ke Backend, Frontend, MinIO)
│   └── check-domain.sh         # Endpoint/Skrip validasi On-Demand TLS (memastikan custom domain yayasan terdaftar di database)
├── /docker                     # Konfigurasi spesifik wadah (Container)
│   ├── docker-compose.prod.yml # Override konfigurasi khusus untuk lingkungan produksi
│   └── docker-swarm.yml        # (Opsional) Konfigurasi jika orkestrasi menggunakan Docker Swarm
├── /minio                      # Object Storage mandiri
│   └── init-buckets.sh         # Skrip otomatis membuat bucket (misal: "legal-docs", "student-photos") saat container pertama kali naik
└── /monitoring                 # Observability & Tracing Stack
    ├── grafana                # Konfigurasi provisioning dashboard visual secara otomatis
    ├── prometheus.yml          # Target scraping metrik (dari Echo backend, PostgreSQL, Node.js)
    └── otel-collector.yml      # Konfigurasi OpenTelemetry untuk melacak latensi request API
