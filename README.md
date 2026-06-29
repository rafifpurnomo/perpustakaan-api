# Library API V2

REST API untuk sistem perpustakaan menggunakan Go, Gin, GORM, MySQL, JWT, dan bcrypt.

## Fitur Utama

- Autentikasi user menggunakan JWT
- Register dan login user
- Manajemen profil user
- Struktur service, repository, controller, dan route yang terpisah
- Migration database menggunakan GORM
- Seeder admin default
- Persiapan dokumentasi API menggunakan Swagger

## Tech Stack

- Go 1.26.4
- Gin
- GORM
- MySQL
- JWT
- bcrypt
- godotenv
- Air untuk hot reload

## Struktur Project

```txt
.
|-- app/                         # Setup router dan dependency injection
|-- documentation/               # ERD dan dokumentasi tambahan
|-- src/
|   |-- config/                  # Konfigurasi env dan database
|   |-- controllers/             # Handler HTTP
|   |-- database/                # Migration dan seeder
|   |-- middleware/              # Middleware JWT
|   |-- repository/              # Query database
|   |-- routes/                  # Definisi route
|   |-- service/                 # Business logic
|   |-- shared/                  # Request DTO
|   `-- utils/                   # Helper JWT dan bcrypt
|-- main.go                      # Entry point server
`-- go.mod
```

## Prasyarat

- Go sesuai versi project
- MySQL
- Git

## Konfigurasi Environment

Buat file `.env` di root project:

```env
APP_PORT=8080

DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=library_api_v2

JWT_SECRET=your-secret-key
```

Pastikan database sudah dibuat di MySQL:

```sql
CREATE DATABASE library_api_v2;
```

## Instalasi

Install dependency:

```bash
go mod download
```

Jalankan migration dan seeder:

```bash
go run ./src/database
```

Jalankan server:

```bash
go run .
```

Server berjalan di:

```txt
http://localhost:8080
```

Port mengikuti nilai `APP_PORT` di `.env`.

## Development

Menjalankan compile check:

```bash
go test ./...
```

Menjalankan server dengan hot reload:

```bash
go tool air
```

## Dokumentasi

- Schema database: [`src/database/README.md`](src/database/README.md)
- ERD: folder [`documentation`](documentation)
- Dokumentasi endpoint akan menggunakan Swagger.

## Catatan

- Password user disimpan dalam bentuk hash menggunakan bcrypt.
- JWT berlaku selama 24 jam.
- File `.env` wajib tersedia karena konfigurasi database dan JWT dibaca dari environment.
