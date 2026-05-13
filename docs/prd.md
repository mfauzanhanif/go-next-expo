# Product Requirements Document (PRD): Sistem SaaS Manajemen Pendidikan Multi-Tenant

## 1. Visi & Objektif Produk

**Visi:**  
Menjadi platform SaaS (Software as a Service) pendidikan paling fleksibel dan terukur di Indonesia, yang memungkinkan yayasan dan lembaga pendidikan mengelola Administrasi, Kesiswaan, SDM, Akademik, Keuangan, dan Aset dalam satu ekosistem terintegrasi.

**Objektif:**
- Menyediakan isolasi data yang aman antar penyewa (tenant) menggunakan strategi skema database mandiri (*logical separation*).
- Menawarkan fleksibilitas komersial melalui fitur/modul yang dapat dibongkar pasang (*add-ons*).
- Memastikan skalabilitas sistem dengan memanfaatkan arsitektur microservices-ready (Go + Next.js + React Native).

---

## 2. Target Pengguna (User Personas)

Sistem ini melayani berbagai tingkat peran dengan hak akses (*Role-Based Access Control*) yang dikelola melalui Zitadel dan Casbin:

- **Super Admin Yayasan:** Mengelola master data yayasan, memantau konsolidasi keuangan dari seluruh lembaga, dan mendistribusikan modul ke lembaga di bawahnya.
- **Admin Lembaga / Unit:** Operator di tingkat sekolah/pondok yang mengelola operasional harian, pendaftaran siswa, dan pengaturan penagihan lembaga tersebut.
- **Tenaga Pendidik / GTK:** Guru atau ustadz yang berinteraksi via web/mobile untuk absensi kelas, pengisian nilai, dan jurnal mengajar.
- **Pengurus Asrama / Musyrif:** Bertanggung jawab memantau kehadiran, perizinan keluar/masuk, serta mencatat pelanggaran dan kedisiplinan.
- **Wali Murid:** Menggunakan aplikasi mobile untuk memantau tagihan bulanan, histori perizinan, dan perkembangan akademik anak.

---

## 3. Topologi Multi-Tenant & Strategi Modul (SaaS Matrix)

Infrastruktur dirancang untuk melayani berbagai skema organisasi pelanggan.

### 3.1. Skema Pelanggan
- **Tipe A (Lembaga Mandiri):** Satu sekolah mandiri berlangganan sistem untuk unitnya sendiri.
- **Tipe B (Yayasan Sentralistik):** Yayasan menggunakan sistem hanya untuk konsolidasi SDM dan Keuangan tingkat pusat.
- **Tipe C (Terpadu):** Yayasan berlangganan paket lengkap dan mendistribusikan lisensi ke beberapa lembaga/unit (misal: MI, SMP, Pondok) di bawah naungannya.

### 3.2. Distribusi Domain (Custom Domain)
Setiap tenant atau unit dapat memiliki identitas sendiri menggunakan *On-Demand TLS via Caddy*.
- Tenant dapat diakses via subdomain: `unit1.saas-pendidikan.com`
- Atau domain kustom mandiri: `smp-hebat.sch.id`

---

## 4. Kebutuhan Fungsional (Functional Requirements)

Kebutuhan fungsional dibagi berdasarkan modul. Hak akses ke modul-modul ini divalidasi oleh Casbin di layer backend sesuai dengan paket lisensi yang dibeli tenant.

### Modul 1: Core System & Master Data SDM (Wajib / Prioritas Utama)
Modul ini adalah pondasi sistem. Penyelesaian master data Guru & Tenaga Kependidikan (GTK) wajib dilakukan pertama kali sebelum memproses entitas data lainnya untuk memastikan hierarki approval berjalan benar.
- **Manajemen Organisasi:** Pendaftaran entitas Yayasan dan Lembaga bawahan.
- **Master Data GTK:** Manajemen biodata lengkap, riwayat pendidikan, dan dokumen kepegawaian.
- **Penugasan Multi-Lembaga:** Satu GTK dapat ditugaskan di lebih dari satu lembaga dalam satu yayasan dengan role yang berbeda (misal: Guru di MI, Pengurus di Pondok).
- **RBAC Engine:** Pengaturan hak akses dinamis berbasis peran (Zitadel + Casbin).

### Modul 2: Kesiswaan & Kesantrian (Add-on)
Struktur data dirancang sangat fleksibel karena status pendidikan formal tidak bersifat wajib bagi residen/santri dalam sebuah ekosistem pendidikan.
- **PPDB Terpadu:** Pendaftaran siswa baru satu pintu dengan opsi pemilihan lembaga formal atau non-formal.
- **Buku Induk Siswa/Santri:** Database profil siswa yang komprehensif.
- **Pemetaan Asrama & Kelas:** Siswa dapat di-assign ke kamar/asrama tertentu tanpa keharusan di-assign ke kelas formal, maupun sebaliknya.

### Modul 3: Kedisiplinan & Perizinan Digital (Add-on)
Sangat krusial untuk lembaga dengan sistem asrama (boarding).
- **Sistem Perizinan:** Alur approval perizinan pulang/keluar berjenjang (dari Musyrif hingga Kepala Pondok) secara digital.
- **Manajemen Poin:** Pencatatan pelanggaran (takzir) dan poin prestasi secara real-time dengan notifikasi otomatis ke wali murid.
- **Log Keamanan:** Pencatatan jam keluar/masuk asrama (bisa diintegrasikan dengan pemindai barcode/RFID kelak).

### Modul 4: Akademik & Penilaian (Add-on)
- **Manajemen Kurikulum:** Setup mata pelajaran, muatan lokal, atau kitab yang diajarkan.
- **Jadwal Mengajar:** Penyusunan roster pelajaran yang terikat pada data GTK.
- **Presensi & Jurnal:** Pengisian absensi harian siswa dan jurnal materi oleh guru via portal/aplikasi.
- **Rapor/Penilaian:** Sistem grading dinamis yang dapat disesuaikan formatnya (Kurikulum Merdeka, K13, atau standar Pesantren).

### Modul 5: Keuangan & Akuntansi (Add-on)
Mendukung konsolidasi dana institusi secara transparan.
- **Akuntansi Double-Entry:** Pencatatan jurnal keuangan, buku besar, dan neraca (menggunakan backend logic yang solid).
- **Billing / Tagihan Otomatis:** Background jobs (via Asynq) yang men-generate tagihan rutin bulanan (SPP, Syahriyah) ke setiap siswa.
- **Manajemen Tabungan:** Sistem pencatatan penitipan uang saku siswa dengan batas penarikan harian.
- **Payment Gateway:** Integrasi untuk pembayaran Virtual Account secara langsung.
- **Konsolidasi Yayasan:** Laporan arus kas terpusat bagi Super Admin Yayasan.

### Modul 6: Mobile Apps (Guru & Wali Murid)
Berbasis React Native (Expo), terhubung dengan API dari arsitektur monorepo.
- **Portal Guru:** Melihat jadwal mengajar harian, input nilai, dan menyetujui perizinan santri.
- **Portal Wali Murid:** Notifikasi push (via Expo Notifications) untuk tagihan baru, riwayat pembayaran, sisa tabungan anak, dan catatan pelanggaran/prestasi.

---

## 5. Kebutuhan Non-Fungsional (Non-Functional Requirements)

**Skalabilitas & Multi-Tenancy:**
- PostgreSQL menggunakan *schema-per-tenant* (dikelola oleh ORM Ent dan Atlas migration) untuk memastikan data antar yayasan/lembaga terisolasi 100%.

**Keamanan (Security):**
- Payload validation di layer frontend (Zod) dan backend (Go-Playground Validator).
- Token JWT/OIDC via Zitadel.
- Rate limiting pada endpoint krusial via middleware Echo.

**Performa:**
- Latensi API target di bawah 200ms.
- Penggunaan Redis v8 untuk caching query yang intensif (seperti menu navigasi dan pengaturan institusi).

**Infrastruktur & Deployment:**
- CI/CD via GitHub Actions untuk proses lint, test, dan build image Docker.
- Penyimpanan berkas statis (dokumen legal, foto siswa) menggunakan MinIO secara mandiri.
- Observabilitas penuh dengan Prometheus, Grafana, dan OpenTelemetry.

---

## 6. Rencana Rilis (Roadmap)

Untuk mitigasi risiko dan percepatan peluncuran ke pasar (Time-to-Market):

- **Fase 1 (MVP Foundation):** Setup arsitektur, Multi-Tenant DB, Casbin + Zitadel, Modul Core, dan Master Data GTK.
- **Fase 2 (Student & Operation):** Modul Kesiswaan (termasuk mengakomodasi siswa non-formal), Modul Perizinan & Kedisiplinan.
- **Fase 3 (Finance Engine):** Modul Keuangan (Double-entry jurnal), Tagihan Otomatis dengan Asynq.
- **Fase 4 (Mobile & Ecosystem):** Peluncuran Mobile App (React Native/Expo) untuk ekosistem Guru dan Wali Murid.