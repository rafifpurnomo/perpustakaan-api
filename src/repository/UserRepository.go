package repository

import (
	"library-api-v2/src/database/migrations"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *migrations.User) error
	GetAll() (*[]migrations.User, error)
	FindByEmail(email string) (*migrations.User, error)
	FindByID(id uint) (*migrations.User, error)
	Update(id uint, user *migrations.User) error
	UpdatePassword(id uint, hashedPassword string) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *migrations.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAll() (*[]migrations.User, error) {
	var users []migrations.User
	return &users, r.db.Find(&users).Error
}

func (r *userRepository) FindByEmail(email string) (*migrations.User, error) {
	var user migrations.User

	err := r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}

func (r *userRepository) FindByID(id uint) (*migrations.User, error) {
	var user migrations.User

	return &user, r.db.First(&user, id).Error
}

func (r *userRepository) Update(id uint, user *migrations.User) error {
	return r.db.Model(&migrations.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepository) UpdatePassword(id uint, hashedPassword string) error {
	return r.db.Model(&migrations.User{}).Where("id = ?", id).Update("password", hashedPassword).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&migrations.User{}, id).Error
}
