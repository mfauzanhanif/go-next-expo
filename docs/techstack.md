# Tech Stack Architecture

## 1. Backend
Sistem backend dirancang untuk performa tinggi, keamanan tipe data, dan skalabilitas multi-tenant yang kokoh.

| Kategori | Teknologi | Deskripsi / Alasan |
| :--- | :--- | :--- |
| **Language** | Go v1.26 | Versi stabil terbaru untuk performa optimal dan efisiensi memori. |
| **Framework** | Echo v5 | Standar industri dengan performa tinggi dan kompatibilitas penuh terhadap ekosistem net/http. |
| **Database** | PostgreSQL v18 | Mendukung penuh strategi *logical separation* (skema terpisah) untuk isolasi data multi-tenant. |
| **ORM / Toolkit** | Ent | Pendekatan *graph-based* yang ideal untuk relasi data kompleks antar modul pendidikan. |
| **Migration** | Atlas | Alat migrasi database *declarative* yang aman untuk mengelola ratusan skema institusi. |
| **Caching** | Redis v8 | Efisiensi akses data frekuensi tinggi dan manajemen *session* terdistribusi. |
| **Queue / Jobs** | Asynq | Penjadwalan tugas latar belakang seperti tagihan otomatis dan notifikasi massal. |
| **Autentikasi** | Zitadel | *Identity Provider* yang dirancang khusus untuk struktur organisasi B2B dan multi-tenant. |
| **Otorisasi (AuthZ)** | Casbin | Manajemen hak akses yang dinamis untuk mendukung model penjualan modul terpisah. |
| **Logging** | Slog / Zerolog | Standar bawaan Go terbaru untuk pengumpulan log sistem yang terstruktur. |
| **Validasi** | Go-Playground Validator | Standar industri untuk memastikan integritas data dari *request payload* klien. |
| **Dokumentasi API** | Swaggo | Pembuatan dokumentasi Swagger/OpenAPI yang tersinkronisasi otomatis dengan kode sumber. |

## 2. Frontend
Menggunakan ekosistem berbasis React untuk memastikan konsistensi antarmuka dan penggunaan ulang kode (*code reusability*) antara platform web dan mobile.

| Kategori | Teknologi | Deskripsi / Alasan |
| :--- | :--- | :--- |
| **Language** | Node v24 | Versi stabil terbaru untuk performa optimal dan efisiensi memori. |
| **Monorepo Manager** | Turborepo | Orkestrasi *build system* yang cerdas untuk mempercepat proses pengembangan tim. |
| **Package Manager** | PNPM | Manajemen dependensi yang sangat efisien dan optimal untuk arsitektur *workspace*. |
| **Shared Packages** | Internal Packages | Pusat logika UI, skema Zod, dan tipe data yang dipakai bersama oleh Web dan Mobile. |

### 2.1. Web
Menggunakan stack modern untuk pengalaman pengguna yang responsif dan SEO-friendly.

| Kategori | Teknologi | Deskripsi / Alasan |
| :--- | :--- | :--- |
| **Framework Utama** | Next.js (App Router) 16 | Framework terbaik untuk *dashboard* manajemen institusi dan optimasi performa. |
| **Library UI** | React | Dasar dari ekosistem frontend modern. |
| **Styling Engine** | Tailwind CSS | *Utility-first* CSS untuk fleksibilitas dan kecepatan pengembangan UI. |
| **Komponen UI** | Shadcn UI | Set komponen yang ringan, *accessible*, dan sepenuhnya dapat dikustomisasi. |
| **State (Client)** | Zustand | Manajemen *state* global yang minimalis dan berperforma tinggi. |
| **State (Server)** | React Query / SWR | Standarisasi *fetching*, *caching*, dan sinkronisasi data dari API backend. |
| **Form Management** | React Hook Form + Zod | Validasi data yang ketat di sisi klien untuk meminimalisir kesalahan input. |

### 2.2. Mobile
Aplikasi mobile lintas platform dengan performa mendekati *native* untuk portal guru dan wali murid.

| Kategori | Teknologi | Deskripsi / Alasan |
| :--- | :--- | :--- |
| **Framework** | React Native | Satu basis kode untuk aplikasi iOS dan Android secara simultan. |
| **Toolkit** | Expo | Ekosistem yang mempercepat siklus pengembangan dan mempermudah distribusi aplikasi. |
| **Styling** | NativeWind | Menggunakan Tailwind CSS untuk konsistensi desain antara web dan mobile. |
| **Navigasi** | Expo Router | Navigasi berbasis struktur file yang intuitif dan sinkron dengan konsep Next.js. |
| **State Management** | Zustand | Memastikan keseragaman logika manajemen *state* dengan aplikasi web. |
| **Data Fetching** | React Query | Konsistensi metode pengambilan data API di seluruh platform klien. |

## 3. Infrastructure
Fokus pada otomasi, keamanan *deployment*, dan pemantauan sistem secara *real-time*.

| Kategori | Teknologi | Deskripsi / Alasan |
| :--- | :--- | :--- |
| **Kontainerisasi** | Docker & Compose | Menjamin konsistensi lingkungan aplikasi dari laptop pengembang hingga server produksi. |
| **Reverse Proxy** | Caddy | Solusi cerdas untuk *On-Demand TLS* agar yayasan dapat menggunakan *custom domain* mereka sendiri. |
| **Task Automation** | Make (Makefile) | Standarisasi alur kerja tim untuk tugas-tugas teknis seperti *build*, *test*, dan *migrate*. |
| **Object Storage** | MinIO | Penyimpanan data mandiri (*self-hosted*) untuk berkas, foto, dan dokumen legal institusi. |
| **CI/CD Pipeline** | GitHub Actions | Otomatisasi siklus rilis fitur secara aman dan terstandardisasi. |
| **Monitoring** | Prometheus & Grafana | Dashboard visual untuk memantau beban server dan kesehatan aplikasi secara langsung. |
| **Tracing** | OpenTelemetry | Diagnosa mendalam untuk melacak latensi dan kegagalan pada setiap permintaan API. |