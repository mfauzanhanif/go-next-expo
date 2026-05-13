# MOBILE STRUCTURE

/mobile
├── /app                            # Expo Router (Pusat routing aplikasi mobile)
│   ├── /(auth)                     # Route Group: Alur autentikasi (Login/Lupa Password)
│   │   ├── login.tsx
│   │   └── _layout.tsx
│   ├── /(tabs)                     # Route Group: Navigasi utama menggunakan Bottom Tabs
│   │   ├── /teacher-portal         # Fitur khusus untuk guru (absensi, nilai)
│   │   ├── /parent-portal          # Fitur khusus untuk wali murid (tabungan siswa, tagihan)
│   │   └── _layout.tsx             # Konfigurasi Bottom Tabs navigation
│   ├── +not-found.tsx              # Halaman 404 (Fallback jika rute tidak ditemukan)
│   └── _layout.tsx                 # Root layout (Provider React Query, Zustand, Zitadel Auth)
├── /assets                         # Aset statis lokal yang dibundle bersama aplikasi
│   ├── /fonts                      # Custom fonts
│   └── /images                     # Logo institusi, placeholder, splash screen
├── /components                     # Komponen UI spesifik mobile
│   ├── /native                     # Komponen yang memanfaatkan API Native (Kamera, Maps)
│   └── /ui                         # Komponen UI spesifik mobile (Bottom Sheet, dll.)
│                                   # Catatan: Skema validasi & logika bisnis sebisa mungkin
│                                   # diimpor dari @workspace/shared
├── /hooks                          # Custom Hooks spesifik platform mobile
│   ├── use-push-notification.ts    # Manajemen push notification (Expo Notifications)
│   └── use-device-info.ts          # Utilitas info perangkat (Platform.OS, versi)
├── /lib                            # Konfigurasi dan utilitas
│   ├── /api-client                 # Konfigurasi Axios/Fetch ke backend (mengatasi isu localhost/IP)
│   └── secure-store.ts             # Pembungkus expo-secure-store untuk token sesi yang aman
├── /store                          # Manajemen state lokal (Zustand)
│   └── use-auth-store.ts           # Sinkronisasi status login di perangkat mobile
├── app.json                        # Konfigurasi utama Expo (Nama app, bundle ID, splash screen, icon)
├── babel.config.js                 # Konfigurasi Babel (Penting untuk NativeWind & Expo Router)
├── metro.config.js                 # Konfigurasi Metro Bundler (Sangat krusial agar Monorepo berfungsi)
├── tailwind.config.ts              # Konfigurasi NativeWind untuk styling
├── tsconfig.json                   # Aturan TypeScript mobile
└── package.json                    # Daftar dependensi mobile (Expo, React Native)
