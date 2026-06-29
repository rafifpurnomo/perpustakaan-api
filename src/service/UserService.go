package service

import (
	"errors"
	"library-api-v2/src/database/migrations"
	"library-api-v2/src/repository"
	"library-api-v2/src/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUsers() (*[]migrations.User, error)
	GetUserByID(id uint) (*migrations.User, error)
	UpdateUser(id uint, user *migrations.User, c *gin.Context) error
	UpdateProfile(id uint, user *migrations.User) error
	UpdatePassword(id uint, newPassword string) error
	DeleteUser(id uint, c *gin.Context) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetAllUsers() (*[]migrations.User, error) {
	users, err := s.userRepo.GetAll()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Data User Kosong")
		}
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUserByID(id uint) (*migrations.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User yang dicari tidak ada")
		}
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(id uint, user *migrations.User, c *gin.Context) error {
	role := c.GetString("role")

	if role != "petugas" && role != "admin" {
		return errors.New("hanya petugas dan admin yang dapat mengubah data user")
	}

	existingUser, err := s.userRepo.FindByID(id)
	existingUserEmail, errEmail := s.userRepo.FindByEmail(user.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("User yang dicari tidak ada")
		}
		return err
	}

	if errEmail == nil && existingUserEmail.ID != existingUser.ID {
		return errors.New("Email sudah digunakan oleh user lain")
	}

	existingUser.NamaLengkap = user.NamaLengkap
	existingUser.Email = user.Email
	existingUser.Role = user.Role

	return s.userRepo.Update(existingUser.ID, existingUser)
}

func (s *userService) UpdateProfile(id uint, user *migrations.User) error {
	existingUser, err := s.userRepo.FindByID(id)
	existingUserEmail, errEmail := s.userRepo.FindByEmail(user.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("User yang dicari tidak ada")
		}
		return err
	}

	if errEmail == nil && existingUserEmail.ID != existingUser.ID {
		return errors.New("Email sudah digunakan oleh user lain")
	}

	existingUser.NamaLengkap = user.NamaLengkap
	existingUser.Email = user.Email

	return s.userRepo.Update(existingUser.ID, existingUser)
}

func (s *userService) UpdatePassword(id uint, newPassword string) error {
	existingUser, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("User yang dicari tidak ada")
		}
		return err
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	return s.userRepo.UpdatePassword(existingUser.ID, hashedPassword)
}

func (s *userService) DeleteUser(id uint, c *gin.Context) error {
	role := c.GetString("role")

	if role != "petugas" && role != "admin" {
		return errors.New("hanya petugas dan admin yang dapat menghapus data user")
	}

	err := s.userRepo.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("User yang dicari tidak ada")
		}

		return err
	}

	return nil
}
