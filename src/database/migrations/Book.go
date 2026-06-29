package migrations

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	JudulBuku string `gorm:"type:varchar(255);not null" json:"judul_buku"`
	Sinopsis  string `gorm:"type:text;not null" json:"sinopsis"`

	KategoriID uint     `gorm:"not null" json:"kategori_id"`
	Kategori   Category `gorm:"foreignKey:KategoriID"`

	StatusBukuID uint       `gorm:"not null" json:"status_buku_id"`
	StatusBuku   BookStatus `gorm:"foreignKey:StatusBukuID"`

	DeletedStatus bool `gorm:"default:false" json:"-"`
}
