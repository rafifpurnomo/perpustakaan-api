# Database Schema

Dokumentasi ini menjelaskan schema database yang digunakan oleh migration GORM di folder `src/database/migrations`.

## Menjalankan Migration dan Seeder

Pastikan `.env` sudah terisi dan database MySQL sudah dibuat, lalu jalankan:

```bash
go run ./src/database
```

Command tersebut akan menjalankan:

- `AutoMigrate` untuk semua model utama
- Seeder admin default

## Tables

GORM juga menambahkan kolom default dari `gorm.Model` pada setiap table:

| Column | Type | Keterangan |
| --- | --- | --- |
| `id` | unsigned integer | Primary key |
| `created_at` | datetime | Waktu data dibuat |
| `updated_at` | datetime | Waktu data diperbarui |
| `deleted_at` | datetime nullable | Soft delete GORM |

## users

Table untuk menyimpan data user.

| Column | Type | Constraint | Keterangan |
| --- | --- | --- | --- |
| `id` | unsigned integer | primary key | ID user |
| `email` | varchar(255) | unique, not null | Email user |
| `password` | varchar(255) | not null | Password yang sudah di-hash |
| `nama_lengkap` | varchar(255) | not null | Nama lengkap user |
| `role` | enum(`admin`, `petugas`, `umum`) | default `umum` | Role user |
| `deleted_status` | boolean | default `false` | Status hapus custom |
| `created_at` | datetime | - | Dibuat oleh `gorm.Model` |
| `updated_at` | datetime | - | Dibuat oleh `gorm.Model` |
| `deleted_at` | datetime nullable | index | Dibuat oleh `gorm.Model` |

## categories

Table untuk menyimpan kategori buku.

| Column | Type | Constraint | Keterangan |
| --- | --- | --- | --- |
| `id` | unsigned integer | primary key | ID kategori |
| `kategori` | varchar(255) | not null | Nama kategori |
| `created_at` | datetime | - | Dibuat oleh `gorm.Model` |
| `updated_at` | datetime | - | Dibuat oleh `gorm.Model` |
| `deleted_at` | datetime nullable | index | Dibuat oleh `gorm.Model` |

## book_statuses

Table untuk menyimpan status buku.

| Column | Type | Constraint | Keterangan |
| --- | --- | --- | --- |
| `id` | unsigned integer | primary key | ID status buku |
| `status` | varchar(255) | not null | Nama status buku |
| `created_at` | datetime | - | Dibuat oleh `gorm.Model` |
| `updated_at` | datetime | - | Dibuat oleh `gorm.Model` |
| `deleted_at` | datetime nullable | index | Dibuat oleh `gorm.Model` |

## books

Table untuk menyimpan data buku.

| Column | Type | Constraint | Keterangan |
| --- | --- | --- | --- |
| `id` | unsigned integer | primary key | ID buku |
| `judul_buku` | varchar(255) | not null | Judul buku |
| `sinopsis` | text | not null | Sinopsis buku |
| `kategori_id` | unsigned integer | not null, foreign key | Relasi ke `categories.id` |
| `status_buku_id` | unsigned integer | not null, foreign key | Relasi ke `book_statuses.id` |
| `deleted_status` | boolean | default `false` | Status hapus custom |
| `created_at` | datetime | - | Dibuat oleh `gorm.Model` |
| `updated_at` | datetime | - | Dibuat oleh `gorm.Model` |
| `deleted_at` | datetime nullable | index | Dibuat oleh `gorm.Model` |

## book_loans

Table untuk menyimpan data peminjaman buku.

| Column | Type | Constraint | Keterangan |
| --- | --- | --- | --- |
| `id` | unsigned integer | primary key | ID peminjaman |
| `user_id` | unsigned integer | not null, foreign key | Relasi ke `users.id` |
| `book_id` | unsigned integer | not null, foreign key | Relasi ke `books.id` |
| `tanggal_peminjaman` | date | not null | Tanggal peminjaman |
| `tanggal_pengembalian` | date | not null | Tanggal pengembalian |
| `deleted_status` | boolean | default `false` | Status hapus custom |
| `created_at` | datetime | - | Dibuat oleh `gorm.Model` |
| `updated_at` | datetime | - | Dibuat oleh `gorm.Model` |
| `deleted_at` | datetime nullable | index | Dibuat oleh `gorm.Model` |

## Relasi

| Relasi | Keterangan |
| --- | --- |
| `categories.id` -> `books.kategori_id` | Satu kategori dapat memiliki banyak buku |
| `book_statuses.id` -> `books.status_buku_id` | Satu status dapat digunakan oleh banyak buku |
| `users.id` -> `book_loans.user_id` | Satu user dapat memiliki banyak peminjaman |
| `books.id` -> `book_loans.book_id` | Satu buku dapat muncul di banyak peminjaman |

## Seeder

Seeder tersedia di:

```txt
src/database/seeders/UserSeeder.go
```

Seeder membuat admin default jika email belum tersedia di database.
