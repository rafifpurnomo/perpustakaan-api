package routes

import (
	controller "library-api-v2/src/controllers"
	"library-api-v2/src/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userController *controller.UserController) {
	user := router.Group("/user")

	user.GET("/", userController.GetAllUsers)
	user.GET("/:id", userController.GetUserByID)
	user.PUT("/profile", middleware.VerifyJWT(), userController.UpdateProfile)
	user.DELETE("/:id", middleware.VerifyJWT(), userController.DeleteUser)
}
