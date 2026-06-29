package migrations

import (
	"time"

	"gorm.io/gorm"
)

type BookLoan struct {
	gorm.Model

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"foreignKey:UserID"`

	BookID uint `gorm:"not null" json:"book_id"`
	Book   Book `gorm:"foreignKey:BookID"`

	TanggalPeminjaman   time.Time `gorm:"type:date;not null" json:"tanggal_peminjaman"`
	TanggalPengembalian time.Time `gorm:"type:date;not null" json:"tanggal_pengembalian"`

	DeletedStatus bool `gorm:"default:false" json:"-"`
}
