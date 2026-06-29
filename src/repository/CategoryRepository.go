package repository

import (
	"library-api-v2/src/database/migrations"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *migrations.Category) error
	GetAll(category *[]migrations.Category) error
	FindByNama(nama string) (*migrations.Category, error)
	FindByID(id uint) (*migrations.Category, error)
	Update(category *migrations.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(category *migrations.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetAll(category *[]migrations.Category) error {
	return r.db.Find(category).Error
}

func (r *categoryRepository) FindByNama(nama string) (*migrations.Category, error) {
	var category migrations.Category
	return &category, r.db.Where("nama = ?", nama).First(&category).Error
}

func (r *categoryRepository) FindByID(id uint) (*migrations.Category, error) {
	var category migrations.Category
	return &category, r.db.First(&category, id).Error
}

func (r *categoryRepository) Update(category *migrations.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&migrations.Category{}, id).Error
}
