package routes

import (
	controller "library-api-v2/src/controllers"
	middleware "library-api-v2/src/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, authController *controller.AuthController) {

	auth := router.Group("/auth")

	auth.POST("/login", authController.Login)
	auth.GET("/me", middleware.VerifyJWT(), authController.Me)
	auth.POST("/registerUmum", authController.RegisterUmum)
}
