package app

import (
	"library-api-v2/src/config"
	controller "library-api-v2/src/controllers"
	"library-api-v2/src/repository"
	"library-api-v2/src/routes"
	"library-api-v2/src/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	// auth
	userRepo := repository.NewUserRepository(config.DB)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	// user
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Routes
	api := router.Group("/api")
	routes.AuthRoutes(api, authController)
	routes.UserRoutes(api, userController)

	return router
}
