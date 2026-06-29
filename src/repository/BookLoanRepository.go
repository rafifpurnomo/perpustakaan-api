package repository

import (
	"library-api-v2/src/database/migrations"

	"gorm.io/gorm"
)

type BookLoanRepository interface {
	Create(bookLoan *migrations.BookLoan) error
	GetAll(bookLoan *[]migrations.BookLoan) error
	FindByID(id uint) (*migrations.BookLoan, error)
	Update(bookLoan *migrations.BookLoan) error
	Delete(id uint) error
}

type bookLoanRepository struct {
	db *gorm.DB
}

func NewBookLoanRepository(db *gorm.DB) BookLoanRepository {
	return &bookLoanRepository{
		db: db,
	}
}

func (r *bookLoanRepository) Create(bookLoan *migrations.BookLoan) error {
	return r.db.Create(bookLoan).Error
}

func (r *bookLoanRepository) GetAll(bookLoan *[]migrations.BookLoan) error {
	return r.db.Find(bookLoan).Error
}

func (r *bookLoanRepository) FindByID(id uint) (*migrations.BookLoan, error) {
	var bookLoan migrations.BookLoan
	return &bookLoan, r.db.First(&bookLoan, id).Error
}

func (r *bookLoanRepository) Update(bookLoan *migrations.BookLoan) error {
	return r.db.Save(bookLoan).Error
}

func (r *bookLoanRepository) Delete(id uint) error {
	return r.db.Delete(&migrations.BookLoan{}, id).Error
}
