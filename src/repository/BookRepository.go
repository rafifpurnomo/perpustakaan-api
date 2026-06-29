package repository

import (
	"library-api-v2/src/database/migrations"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *migrations.Book) error
	GetAll() (*[]migrations.Book, error)
	FindByJudul(judul string) (*migrations.Book, error)
	FindByID(id uint) (*migrations.Book, error)
	Update(book *migrations.Book) error
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) Create(book *migrations.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) GetAll() (*[]migrations.Book, error) {
	var books []migrations.Book
	return &books, r.db.Find(&books).Error
}

func (r *bookRepository) FindByJudul(judul string) (*migrations.Book, error) {
	var book migrations.Book

	return &book, r.db.Where("judul = ?", judul).First(&book).Error
}

func (r *bookRepository) FindByID(id uint) (*migrations.Book, error) {
	var book migrations.Book

	return &book, r.db.First(&book, id).Error
}

func (r *bookRepository) Update(book *migrations.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&migrations.Book{}, id).Error
}
