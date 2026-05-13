# Dokumentasi Makefile

Proyek ini menggunakan **Makefile** untuk menyederhanakan alur kerja pengembangan, dari menjalankan environment lokal hingga build, testing, dan deployment. Terdapat tiga lapis Makefile utama:

1. **Root Makefile** (`/Makefile`): Orkestrasi layanan penuh dengan Docker Compose dan akses global.
2. **Backend Makefile** (`/backend/Makefile`): Perintah khusus aplikasi Go (Echo, Atlas, Ent).
3. **Frontend Makefile** (`/frontend/Makefile`): Perintah khusus untuk environment Turborepo (Next.js, React Native, PNPM).

---

## 1. Root Makefile (`/Makefile`)

Berfungsi sebagai _entry point_ utama untuk mengontrol keseluruhan _stack_. Jika Anda berada di root direktori, Anda dapat menjalankan perintah-perintah berikut:

### Bantuan

- `make help` : Tampilkan daftar command yang tersedia.

### Docker Compose — Full Stack

- `make up` : Menjalankan semua layanan di latar belakang (detached) dan otomatis mem-build image.
- `make down` : Menghentikan dan menghapus semua container dari layanan.
- `make restart` : Menghentikan layanan (`down`) kemudian menjalankannya kembali (`up`).
- `make ps` : Menampilkan status dari semua container.
- `make logs` : Menampilkan output logs dari seluruh layanan (mode follow).
- `make logs-backend` : Menampilkan logs spesifik untuk container backend dan worker.
- `make logs-frontend` : Menampilkan logs spesifik untuk container frontend.
- `make logs-infra` : Menampilkan logs dari layanan infrastruktur (PostgreSQL, Redis, Zitadel, MinIO, Caddy).
- `make build` : Mem-build ulang semua Docker image.
- `make pull` : Mengunduh versi terbaru untuk base image.

### Docker Compose — Layanan Selektif

- `make up-infra` : Hanya menjalankan infrastruktur utama (PostgreSQL, Redis, Zitadel, MinIO).
- `make up-app` : Hanya menjalankan service aplikasi (Backend, Worker, Frontend).
- `make up-monitoring`: Hanya menjalankan _stack_ pemantauan (Prometheus, Grafana, OpenTelemetry).
- `make up-backend` : Menjalankan backend, worker, beserta infrastruktur pendukungnya.
- `make up-frontend` : Menjalankan frontend, backend, worker, dan infrastruktur pendukung.

### Development Lokal (Aplikasi tanpa Docker)

- `make dev` : Menjalankan infrastruktur via Docker, lalu memberikan instruksi cara menjalankan backend & frontend secara manual/lokal di terminal terpisah.
- `make dev-backend` : Menjalankan infrastruktur via Docker, lalu secara otomatis menjalankan _development server_ backend lokal.
- `make dev-frontend` : Menjalankan infrastruktur (asumsi backend sudah jalan), kemudian menjalankan _development server_ frontend web.

### Delegasi ke Sub-Makefile

- `make backend-%` : Menjalankan command yang ada di backend. Contoh: `make backend-test` sama dengan menjalankan `make test` di dalam folder `/backend`.
- `make frontend-%` : Menjalankan command yang ada di frontend. Contoh: `make frontend-lint` sama dengan menjalankan `make lint` di dalam folder `/frontend`.

### Database

- `make db-migrate` : Menjalankan migrasi database via Atlas (delegasi ke backend).
- `make db-status` : Mengecek status skema migrasi database via Atlas.

### Production

- `make deploy` : Mendeploy _stack_ menggunakan `docker-compose.prod.yaml`.
- `make deploy-down` : Menghentikan deployment produksi.

### Status & Cleanup

- `make status` : Menampilkan status _container_, penggunaan disk space Docker, dan volume yang dipakai.
- `make clean` : Menghentikan container, sekaligus menghapus seluruh **data volume** Docker (membersihkan data).
- `make prune` : Membersihkan semua _resources_ (container, network, image, volume) Docker yang tidak terpakai (dangling/orphans).

---

## 2. Backend Makefile (`/backend/Makefile`)

Perintah khusus untuk berinteraksi dengan basis kode Go. Jalankan ini di dalam folder `/backend`.

### Bantuan

- `make help` : Tampilkan daftar command untuk backend.

### Development

- `make run` : Menjalankan API server secara lokal.
- `make run-worker` : Menjalankan Asynq Worker secara lokal.
- `make run-all` : Menjalankan API server dan Worker secara paralel.

### Build

- `make build` : Melakukan kompilasi API server ke `/bin/api-server`.
- `make build-worker`: Melakukan kompilasi Worker ke `/bin/worker`.
- `make build-all` : Kompilasi semua binary backend.

### Quality Assurance (QA)

- `make test` : Menjalankan semua _unit test_.
- `make test-coverage`: Menjalankan _test_ dan men-generate laporan cakupan (coverage HTML) di folder `/bin`.
- `make lint` : Menjalankan `golangci-lint` untuk memeriksa kualitas kode.
- `make fmt` : Memformat kode menggunakan `go fmt` dan `goimports`.
- `make vet` : Melakukan analisis statis dengan `go vet`.
- `make check` : Menjalankan semua tahapan _check_ (fmt, vet, lint, test) sekaligus.

### Dependency & Tools

- `make tidy` : Membersihkan file `go.mod` dan `go.sum` (`go mod tidy`).
- `make vendor` : Memindahkan _dependencies_ ke dalam folder `vendor/`.
- `make swagger` : Men-generate atau memperbarui dokumentasi API Open-API (Swagger) via alat Swaggo.

### ORM (Ent)

- `make ent-generate` : Men-generate ulang _boilerplate_ ORM Ent berdasarkan skema yang ada.
- `make ent-new NAME=User`: Menginisialisasi _schema template_ baru untuk tabel (misalnya tabel User).

### Database Migration (Atlas)

- `make migrate` : Mengaplikasikan migrasi database pending.
- `make migrate-status`: Mengecek _drift_ atau perbedaan status skema.
- `make migrate-new NAME=nama_migrasi`: Men-generate skrip migrasi deklaratif baru otomatis dari perubahan skema Ent.
- `make migrate-hash` : Me-_rehash_ file _directory state_ migrasi setelah modifikasi manual file `.sql`.

### Docker

- `make docker-build` : Mem-build Docker image untuk backend.
- `make docker-run` : Menjalankan container yang baru di-build.

### Cleanup

- `make clean` : Menghapus folder `bin/` dan membersihkan Go cache.

---

## 3. Frontend Makefile (`/frontend/Makefile`)

Perintah khusus untuk berinteraksi dengan monorepo PNPM + Turborepo. Jalankan ini di dalam folder `/frontend`.

### Bantuan

- `make help` : Tampilkan daftar command untuk frontend.

### Setup

- `make install` : Menjalankan `pnpm install` untuk seluruh _workspace_.
- `make install-frozen`: Menjalankan `pnpm install --frozen-lockfile` untuk instalasi ketat (biasanya untuk CI/CD).

### Development

- `make dev` : Menjalankan Next.js Web dan React Native Mobile secara bersamaan (via Turborepo).
- `make dev-web` : Hanya menjalankan _dev server_ Next.js Web.
- `make dev-mobile` : Hanya menjalankan _dev server_ Expo Mobile.

### Build

- `make build` : Mem-build seluruh _workspace_ menggunakan Turborepo (mengambil layer cache).
- `make build-web` : Mem-build spesifik _package_ Next.js Web.
- `make build-packages`: Mem-build _shared packages_ di dalam monorepo.

### Quality Assurance (QA)

- `make lint` : Menjalankan Eslint untuk seluruh _workspace_.
- `make lint-web` : Hanya melinting untuk _app web_.
- `make lint-mobile` : Hanya melinting untuk _app mobile_.
- `make format` : Melakukan perbaikan otomatis (fix) _formatting_ kode via Prettier.
- `make format-check` : Memeriksa kesesuaian format kode dengan Prettier (tanpa mengubah isi file).
- `make type-check` : Mengecek _type safety_ TypeScript.
- `make test` : Menjalankan seluruh proses _unit test_.
- `make test-web` : Menjalankan _unit test_ spesifik pada _app web_.
- `make check` : Menjalankan seluruh validasi (`lint`, `type-check`, `test`).

### Expo (Mobile App)

- `make expo-start` : Menjalankan _dev server_ React Native (Expo).
- `make expo-ios` : Menjalankan Expo lalu membuka emulator iOS langsung.
- `make expo-android` : Menjalankan Expo lalu membuka emulator Android langsung.
- `make expo-prebuild`: Mengubah proyek Expo managed _workflow_ menjadi native iOS & Android _build artifacts_.

### Docker

- `make docker-build` : Mem-build Docker image khusus app web secara mandiri (menggunakan fitur _Turbo Prune_).

### Cleanup

- `make clean` : Menghapus seluruh instansi folder `node_modules`, `out`, `.next`, dan cache _Turbo_ dari seluruh _packages_ di dalam _workspace_.
- `make clean-cache` : Menghapus file _state_ internal Turborepo saja.
