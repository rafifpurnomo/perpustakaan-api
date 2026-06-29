package controller

import (
	"library-api-v2/src/database/migrations"
	"library-api-v2/src/service"
	"library-api-v2/src/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {

	var req shared.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})

		return

	}

	token, err := ac.authService.Login(req.Email, req.Password)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "login berhasil",
		"token":   token,
	})

}

func (ac *AuthController) Me(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "userID not found in context",
		})
		return
	}

	userID, ok := userID.(uint)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid user id",
		})
		return
	}

	user, err := ac.authService.Me(userID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func (ac *AuthController) RegisterUmum(c *gin.Context) {

	var req shared.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	user := &migrations.User{
		NamaLengkap: req.Name,
		Email:       req.Email,
		Password:    req.Password,
	}

	err := ac.authService.RegisterUmum(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "user created successfully",
		"data":    user,
	})

}
