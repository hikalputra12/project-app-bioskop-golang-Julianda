# 🎬 Aplikasi Pemesanan Tiket Bioskop (Backend)

![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791?logo=postgresql&logoColor=white)
![Clean Architecture](https://img.shields.io/badge/Architecture-Clean%20Architecture-success)
![Status](https://img.shields.io/badge/Status-Development-yellow)

Backend service untuk aplikasi pemesanan tiket bioskop yang dibangun menggunakan **Golang** dengan penerapan **Clean Architecture**.  
Aplikasi ini menangani proses bisnis end-to-end mulai dari **registrasi user**, **verifikasi OTP email**, **manajemen bioskop**, hingga **transaksi pemesanan tiket**.

---

## 📋 Daftar Isi

- [Tentang Proyek](#-tentang-proyek)
- [Teknologi](#-teknologi)
- [Struktur Proyek](#-struktur-proyek)
- [Fitur Utama](#-fitur-utama)
- [Skema Database](#-skema-database)
- [Instalasi & Menjalankan](#️-instalasi--menjalankan)
- [Dokumentasi API](#-dokumentasi-api)

---

## 📖 Tentang Proyek

Sistem ini dirancang untuk menangani **trafik tinggi** dengan arsitektur yang **modular, scalable, dan maintainable**.

- Menggunakan **PostgreSQL** sebagai database utama  
- **Redis** disiapkan untuk caching *(future implementation)*  
- Autentikasi menggunakan **Session-based Authentication**
  - Session disimpan di database
  - Dikirim ke client melalui **HTTP Cookie**
- Verifikasi email menggunakan **OTP (One Time Password)**

---

## 🛠 Teknologi

- **Bahasa**: Go (Golang) v1.25+
- **Database**: PostgreSQL
- **Driver DB**: `pgx/v5`
- **Router**: `go-chi`
- **Configuration**: `Viper` & `godotenv`
- **Logging**: `Zap`
- **Email Service**: `Gomail` (SMTP)
- **Dependency Injection**: `Google Wire`
- **Migration / Backup**: SQL File (`backup_ticket.sql`)

---

## 📂 Struktur Proyek

```text
.
├── cmd/
│   └── server.go              # Entry point aplikasi
├── internal/
│   ├── adaptor/               # HTTP Handler / Controller
│   ├── data/
│   │   ├── entity/            # Struct / Model Database
│   │   └── repository/        # Query & akses database
│   ├── dto/                   # Request & Response DTO
│   ├── middleware/            # Auth, Logging, Session middleware
│   ├── usecase/               # Business Logic
│   └── wire/                  # Dependency Injection (Google Wire)
├── pkg/
│   ├── database/              # Setup koneksi PostgreSQL
│   └── utils/
│       ├── config.go          # Load environment variables
│       ├── email_worker.go    # Background worker email
│       ├── send_otp.go        # SMTP OTP logic
│       ├── hashed_password.go # Hash password (bcrypt)
│       ├── logger.go          # Zap logger
│       └── response.go        # Standard JSON response
├── backup_ticket.sql
├── ticket_bioskop.postman_collection.json
├── go.mod
└── .env
