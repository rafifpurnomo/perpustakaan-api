package repository

import (
	"library-api-v2/src/database/migrations"

	"gorm.io/gorm"
)

type BookStatusRepository interface {
	Create(bookStatus *migrations.BookStatus) error
	GetAll(bookStatus *[]migrations.BookStatus) error
	FindByStatus(status string) (*migrations.BookStatus, error)
	FindByID(id uint) (*migrations.BookStatus, error)
	Update(bookStatus *migrations.BookStatus) error
	Delete(id uint) error
}

type bookStatusRepository struct {
	db *gorm.DB
}

func NewBookStatusRepository(db *gorm.DB) BookStatusRepository {
	return &bookStatusRepository{
		db: db,
	}
}

func (r *bookStatusRepository) Create(bookStatus *migrations.BookStatus) error {
	return r.db.Create(bookStatus).Error
}

func (r *bookStatusRepository) GetAll(bookStatus *[]migrations.BookStatus) error {
	return r.db.Find(bookStatus).Error
}

func (r *bookStatusRepository) FindByStatus(status string) (*migrations.BookStatus, error) {
	var bookStatus migrations.BookStatus
	return &bookStatus, r.db.Where("status = ?", status).First(&bookStatus).Error
}

func (r *bookStatusRepository) FindByID(id uint) (*migrations.BookStatus, error) {
	var bookStatus migrations.BookStatus
	return &bookStatus, r.db.First(&bookStatus, id).Error
}

func (r *bookStatusRepository) Update(bookStatus *migrations.BookStatus) error {
	return r.db.Save(bookStatus).Error
}

func (r *bookStatusRepository) Delete(id uint) error {
	return r.db.Delete(&migrations.BookStatus{}, id).Error
}
