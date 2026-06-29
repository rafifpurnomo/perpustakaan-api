package migrations

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email       string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password    string `gorm:"type:varchar(255);not null" json:"-"`
	NamaLengkap string `gorm:"type:varchar(255);not null" json:"nama_lengkap"`
	Role        string `gorm:"type:enum('admin','petugas','umum');default:'umum'" json:"role"`

	DeletedStatus bool `gorm:"default:false" json:"-"`
}
