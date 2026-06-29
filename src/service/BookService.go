package service

import (
	"errors"
	"library-api-v2/src/database/migrations"
	"library-api-v2/src/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookService interface {
	getAllBuku() (*[]migrations.Book, error)
	tambahBuku(buku *migrations.Book, c *gin.Context) error
	editBuku(buku *migrations.Book, c *gin.Context) error
	hapusBuku(bukuID uint, c *gin.Context) error
	cariBukuByID(bukuID uint) (*migrations.Book, error)
	cariBukuByJudul(judul string) (*migrations.Book, error)
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (s *bookService) getAllBuku() (*[]migrations.Book, error) {
	books, err := s.bookRepo.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tidak ada buku yang ditemukan")
		}
		return nil, err
	}

	return books, nil
}

func (s *bookService) tambahBuku(buku *migrations.Book, c *gin.Context) error {

	role := c.GetString("role")
	if role != "petugas" {
		return errors.New("hanya petugas yang dapat menambahkan buku")
	}

	if buku == nil {
		return errors.New("buku tidak boleh kosong")
	}

	foundBook, err := s.bookRepo.FindByJudul(buku.JudulBuku)

	if err == nil && foundBook != nil {
		return errors.New("buku dengan judul yang sama sudah ada")
	}

	return s.bookRepo.Create(buku)
}

func (s *bookService) editBuku(buku *migrations.Book, c *gin.Context) error {

	role := c.GetString("role")
	if role != "petugas" {
		return errors.New("hanya petugas yang dapat mengedit buku")
	}

	if buku == nil {
		return errors.New("buku tidak boleh kosong")
	}

	foundBook, err := s.bookRepo.FindByJudul(buku.JudulBuku)

	if err == nil && foundBook != nil {
		return errors.New("buku dengan judul yang sama sudah ada")
	}

	return s.bookRepo.Update(buku)
}

func (s *bookService) hapusBuku(bukuID uint, c *gin.Context) error {
	role := c.GetString("role")

	if role != "petugas" {
		return errors.New("hanya petugas yang dapat menghapus buku")
	}

	foundBook, _ := s.bookRepo.FindByID(bukuID)

	if foundBook == nil {
		return errors.New("buku tidak ditemukan")
	}

	return s.bookRepo.Delete(bukuID)
}

func (s *bookService) cariBukuByID(bukuID uint) (*migrations.Book, error) {
	buku, err := s.bookRepo.FindByID(bukuID)
	if err != nil {
		return nil, err
	}

	return buku, nil
}

func (s *bookService) cariBukuByJudul(judul string) (*migrations.Book, error) {
	buku, err := s.bookRepo.FindByJudul(judul)
	if err != nil {
		return nil, err
	}

	return buku, nil
}
