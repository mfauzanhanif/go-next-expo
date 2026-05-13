# WEB STRUCTURE

/web
├── /app                            # Next.js App Router (Pusat routing aplikasi)
│   ├── /(auth)                     # Route Group: Autentikasi (tanpa layout dashboard)
│   │   ├── /login/page.tsx
│   │   └── layout.tsx
│   ├── /(dashboard)                # Route Group: Area utama setelah login
│   │   ├── /administrator          # Modul dan kontrol akses khusus peran Administrator
│   │   ├── /staff-tu               # Modul operasional sehari-hari untuk peran Staff TU
│   │   ├── /student-savings        # Modul pencatatan dan pengelolaan tabungan siswa
│   │   └── layout.tsx              # Layout dashboard (Sidebar, Header, persistensi state)
│   ├── /api                        # (Opsional) Route Handlers/BFF untuk proxy ke backend Go
│   ├── globals.css                 # Entry point untuk Tailwind CSS web
│   └── layout.tsx                  # Root layout (Provider React Query, integrasi Zitadel)
├── /components                     # Komponen spesifik yang hanya dipakai di Web Next.js
│   ├── /layouts                    # Komponen tata letak (Navbar, Sidebar, Page Wrapper)
│   └── /features                   # Komponen spesifik domain (contoh: formulir pendaftaran)
│                                   # Catatan: Komponen dasar Shadcn UI sebaiknya diambil dari @workspace/ui
├── /hooks                          # Custom React Hooks spesifik alur kerja web
│   ├── use-auth.ts                 # Logika sesi pengguna
│   └── use-keyboard-shortcut.ts    # Aksesibilitas khusus desktop/web
├── /lib                            # Fungsi utilitas murni dan konfigurasi library pihak ketiga
│   ├── /api-client                 # Konfigurasi instance fetch/Axios menuju backend Echo
│   ├── /query-keys                 # Standarisasi key untuk React Query / SWR
│   └── utils.ts                    # Helper umum spesifik web
├── /store                          # Manajemen state lokal sisi klien (Zustand)
│   └── use-ui-store.ts             # State UI global web (misal: toggle sidebar, dark mode)
├── .eslintrc.json                  # Linter khusus proyek Next.js
├── next.config.mjs                 # Konfigurasi Next.js (optimasi gambar, env, transpilePackages)
├── tailwind.config.ts              # Konfigurasi Tailwind (biasanya meng-extend dari @workspace/config)
├── tsconfig.json                   # Aturan TypeScript web
└── package.json                    # Daftar dependensi web (mengambil referensi dari workspace)