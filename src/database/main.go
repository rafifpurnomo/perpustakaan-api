package main

import (
	"library-api-v2/src/config"
	"library-api-v2/src/database/migrations"
	"library-api-v2/src/database/seeders"
	"log"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&migrations.User{},
		&migrations.Book{},
		&migrations.Category{},
		&migrations.BookStatus{},
		&migrations.BookLoan{},
	)

	if err != nil {
		log.Fatal("Migration Failed:", err)
	}

	log.Println("Migration Success")
}

func main() {
	config.ConnectDatabase()
	Migrate()
	seeders.UserSeeder()
	log.Println("Seeding Success")
}
