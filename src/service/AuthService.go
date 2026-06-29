package service

import (
	"errors"
	"library-api-v2/src/database/migrations"
	"library-api-v2/src/repository"
	"library-api-v2/src/utils"

	"strings"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(email string, password string) (string, error)
	Me(userID uint) (*migrations.User, error)
	RegisterPetugas(user *migrations.User, currentRole string) error
	RegisterUmum(user *migrations.User) error
	RegisterAdmin(user *migrations.User, currentRole string) error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Login(email string, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("akun tidak terdaftar")
		}

		return "", err
	}

	if err := utils.CheckPassword(password, user.Password); err != nil {
		return "", errors.New("password salah")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Me(userID uint) (*migrations.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) RegisterPetugas(user *migrations.User, currentRole string) error {
	if currentRole != "admin" && currentRole != "petugas" {
		return errors.New("hanya admin dan petugas yang dapat menambahkan petugas")
	}

	if err := validateRegisterUser(user); err != nil {
		return err
	}

	userFound, userFoundErr := s.userRepo.FindByEmail(user.Email)
	if userFoundErr == nil && userFound != nil {
		return errors.New("email sudah digunakan")
	}

	user.Role = "petugas"
	hashedPassword, hashedPasswordErr := utils.HashPassword(user.Password)

	if hashedPasswordErr != nil {
		return hashedPasswordErr
	}

	user.Password = hashedPassword
	return s.userRepo.Create(user)
}

func (s *authService) RegisterUmum(user *migrations.User) error {
	if err := validateRegisterUser(user); err != nil {
		return err
	}

	user.Role = "umum"
	userFound, userFoundErr := s.userRepo.FindByEmail(user.Email)
	if userFoundErr == nil && userFound != nil {
		return errors.New("email sudah digunakan")
	}

	hashedPassword, hashedPasswordErr := utils.HashPassword(user.Password)
	if hashedPasswordErr != nil {
		return hashedPasswordErr
	}

	user.Password = hashedPassword
	return s.userRepo.Create(user)
}

func (s *authService) RegisterAdmin(user *migrations.User, currentRole string) error {
	if currentRole != "admin" {
		return errors.New("hanya admin yang dapat menambahkan admin")
	}

	if err := validateRegisterUser(user); err != nil {
		return err
	}

	userFound, userFoundErr := s.userRepo.FindByEmail(user.Email)
	if userFoundErr == nil && userFound != nil {
		return errors.New("email sudah digunakan")
	}

	user.Role = "admin"
	hashedPassword, hashedPasswordErr := utils.HashPassword(user.Password)
	if hashedPasswordErr != nil {
		return hashedPasswordErr
	}

	user.Password = hashedPassword
	return s.userRepo.Create(user)
}

func validateRegisterUser(user *migrations.User) error {
	if user == nil {
		return errors.New("data user tidak boleh kosong")
	}

	user.NamaLengkap = strings.TrimSpace(user.NamaLengkap)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if user.NamaLengkap == "" {
		return errors.New("nama lengkap tidak boleh kosong")
	}

	if user.Email == "" {
		return errors.New("email tidak boleh kosong")
	}

	if user.Password == "" {
		return errors.New("password tidak boleh kosong")
	}

	return nil
}
