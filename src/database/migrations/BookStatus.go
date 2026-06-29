package migrations

import "gorm.io/gorm"

type BookStatus struct {
	gorm.Model

	Status string `gorm:"type:varchar(255);not null" json:"status"`

	Buku []Book `gorm:"foreignKey:StatusBukuID"`
}
