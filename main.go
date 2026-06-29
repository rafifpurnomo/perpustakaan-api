package main

import (
	"library-api-v2/app"
	"library-api-v2/src/config"
)

func main() {
	config.ConnectDatabase()

	router := app.SetupRouter()
	router.Run(":" + config.GetEnv("APP_PORT"))
}
