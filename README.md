# Library API V2

REST API untuk sistem perpustakaan menggunakan Go, Gin, GORM, MySQL, JWT, dan bcrypt.

## Fitur

- Register dan login user
- Autentikasi menggunakan JWT Bearer Token
- Endpoint profil user yang sedang login
- CRUD dasar user yang sudah tersedia sebagian
- Migration database untuk user, buku, kategori, status buku, dan peminjaman buku
- Seeder admin default

## Tech Stack

- Go 1.26.4
- Gin
- GORM
- MySQL
- JWT
- bcrypt
- godotenv

## Struktur Project

```txt
.
|-- app/                         # Setup router dan dependency injection
|-- documentation/               # ERD dan dokumentasi database
|-- src/
|   |-- config/                  # Konfigurasi env dan database
|   |-- controllers/             # Handler HTTP
|   |-- database/
|   |   |-- migrations/          # Model GORM
|   |   `-- seeders/             # Seeder awal
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

Seeder akan membuat admin default:

```txt
email: admin@admin.com
password: admin123
role: admin
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

## Autentikasi

Endpoint yang dilindungi membutuhkan header:

```http
Authorization: Bearer <token>
```

Token didapat dari endpoint login.

## Endpoint

Base URL:

```txt
/api
```

### Auth

| Method | Endpoint | Auth | Deskripsi |
| --- | --- | --- | --- |
| POST | `/auth/registerUmum` | Tidak | Register user umum |
| POST | `/auth/login` | Tidak | Login dan mendapatkan token |
| GET | `/auth/me` | Ya | Mengambil data user yang sedang login |

### User

| Method | Endpoint | Auth | Deskripsi |
| --- | --- | --- | --- |
| GET | `/user/` | Tidak | Mengambil semua user |
| GET | `/user/:id` | Tidak | Mengambil user berdasarkan ID |
| PUT | `/user/profile` | Ya | Update profil user yang sedang login |
| DELETE | `/user/:id` | Ya | Hapus user berdasarkan ID |

Catatan: beberapa aksi user melakukan pengecekan role di service. Role `admin` dan `petugas` memiliki akses lebih tinggi untuk operasi tertentu.

## Contoh Request

### Register User Umum

```http
POST /api/auth/registerUmum
Content-Type: application/json
```

```json
{
  "name": "User Umum",
  "email": "user@example.com",
  "password": "password123"
}
```

### Login

```http
POST /api/auth/login
Content-Type: application/json
```

```json
{
  "email": "admin@admin.com",
  "password": "admin123"
}
```

Contoh response:

```json
{
  "message": "login berhasil",
  "success": true,
  "token": "<jwt-token>"
}
```

### Me

```http
GET /api/auth/me
Authorization: Bearer <jwt-token>
```

### Update Profile

```http
PUT /api/user/profile
Authorization: Bearer <jwt-token>
Content-Type: application/json
```

```json
{
  "name": "Nama Baru",
  "email": "nama.baru@example.com",
  "password": "passwordbaru"
}
```

### Get All Users

```http
GET /api/user/
```

### Get User By ID

```http
GET /api/user/1
```

### Delete User

```http
DELETE /api/user/1
Authorization: Bearer <jwt-token>
```

## Model Database

Migration yang tersedia:

- `User`
- `Book`
- `Category`
- `BookStatus`
- `BookLoan`

Relasi utama:

- `Book` memiliki `Category`
- `Book` memiliki `BookStatus`
- `BookLoan` memiliki `User`
- `BookLoan` memiliki `Book`

## Development

Menjalankan test/compile check:

```bash
go test ./...
```

Jika menggunakan Air untuk hot reload:

```bash
go tool air
```

## Catatan

- Password user disimpan dalam bentuk hash menggunakan bcrypt.
- JWT berlaku selama 24 jam.
- File `.env` wajib tersedia karena konfigurasi database dan JWT dibaca dari environment.
- Dokumentasi ERD tersedia di folder `documentation/`.
