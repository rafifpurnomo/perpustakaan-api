package seeders

import (
	"library-api-v2/src/config"
	"library-api-v2/src/database/migrations"
	"library-api-v2/src/utils"
	"log"
)

func UserSeeder() {

	password, _ := utils.HashPassword("admin123")

	users := []migrations.User{
		{
			NamaLengkap: "admin default",
			Email:       "admin@admin.com",
			Password:    password,
			Role:        "admin",
		},
	}

	for _, user := range users {

		var existing migrations.User

		err := config.DB.
			Where("email = ?", user.Email).
			First(&existing).Error

		if err != nil {
			config.DB.Create(&user)
		}
	}

	log.Println("User Seeder Success")
}
