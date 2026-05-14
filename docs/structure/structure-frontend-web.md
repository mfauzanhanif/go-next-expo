# WEB STRUCTURE

Struktur direktori aktual dari proyek web Next.js (`frontend/web`).

/web
├── /app                            # Next.js App Router (Pusat routing aplikasi)
│   ├── Layout.tsx                  # Root layout aplikasi (perlu diubah ke huruf kecil: layout.tsx)
│   ├── global.css                  # Entry point CSS global
│   └── page.tsx                    # Halaman utama (Home Page)
├── /components                     # Direktori komponen
├── /hooks                          # Custom React Hooks
├── /lib                            # Fungsi utilitas
├── /public                         # Aset publik statis
│   └── .gitkeep                    # Placeholder agar Git mendeteksi folder ini
├── /store                          # Manajemen state lokal
├── next-env.d.ts                   # Deklarasi tipe otomatis dari Next.js
├── next.config.mjs                 # Konfigurasi Next.js (termasuk output standalone)
├── package.json                    # Daftar dependensi khusus web (Next.js, React)
└── tsconfig.json                   # Konfigurasi TypeScript untuk web