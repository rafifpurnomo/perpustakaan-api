package migrations

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	Kategori string `gorm:"type:varchar(255);not null" json:"kategori"`

	Buku []Book `gorm:"foreignKey:KategoriID"`
}
