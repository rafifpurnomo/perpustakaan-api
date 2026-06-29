CREATE TABLE `user`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `nama_lengkap` VARCHAR(255) NOT NULL,
    `role` ENUM('admin', 'petugas', 'umum') NOT NULL,
    `created_at` DATETIME NOT NULL,
    `update_at` DATETIME NOT NULL,
    `deleted_at` DATETIME NOT NULL,
    `deleted_status` TINYINT NOT NULL DEFAULT 0
);
CREATE TABLE `buku`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `id_kategori` INT NOT NULL,
    `id_status_buku` INT NOT NULL,
    `judul_buku` VARCHAR(255) NOT NULL,
    `sinopsis` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME NOT NULL,
    `deleted_status` TINYINT NOT NULL DEFAULT 0
);
CREATE TABLE `kategori`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `kategori` VARCHAR(255) NOT NULL
);
CREATE TABLE `peminjaman`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `id_user` INT NOT NULL,
    `id_buku` BIGINT NOT NULL,
    `tanggal_peminjaman` DATE NOT NULL,
    `tanggal_pengembalian` DATE NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME NOT NULL,
    `deleted_status` TINYINT NOT NULL DEFAULT 0
);
CREATE TABLE `status_buku`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `status` VARCHAR(255) NOT NULL
);
ALTER TABLE
    `peminjaman` ADD CONSTRAINT `peminjaman_id_buku_foreign` FOREIGN KEY(`id_buku`) REFERENCES `buku`(`id`);
ALTER TABLE
    `buku` ADD CONSTRAINT `buku_id_status_buku_foreign` FOREIGN KEY(`id_status_buku`) REFERENCES `status_buku`(`id`);
ALTER TABLE
    `peminjaman` ADD CONSTRAINT `peminjaman_id_user_foreign` FOREIGN KEY(`id_user`) REFERENCES `user`(`id`);
ALTER TABLE
    `buku` ADD CONSTRAINT `buku_id_kategori_foreign` FOREIGN KEY(`id_kategori`) REFERENCES `kategori`(`id`);