package controller

import (
	"library-api-v2/src/database/migrations"
	"library-api-v2/src/service"
	"library-api-v2/src/shared"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data User Berhasil Diambil",
		"data":    users,
	})
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	idParam, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id parameter is required",
		})
		return
	}

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id parameter must be a valid number",
		})
		return
	}

	user, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data User Berhasil Diambil",
		"data":    user,
	})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	idParam, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id parameter is required",
		})
		return
	}

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id parameter must be a valid number",
		})
		return
	}

	var req shared.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	newDataUser := &migrations.User{
		NamaLengkap: req.Name,
		Email:       req.Email,
		Role:        req.Role,
	}

	err = uc.userService.UpdateUser(uint(id), newDataUser, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data User Berhasil Diperbarui",
	})

}

func (uc *UserController) UpdateProfile(c *gin.Context) {
	id, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "userID not found in context",
		})
		return
	}

	var req shared.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	newDataUser := &migrations.User{
		NamaLengkap: req.Name,
		Email:       req.Email,
	}

	userID, ok := id.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid userID in context",
		})
		return
	}

	err := uc.userService.UpdateProfile(userID, newDataUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data Profile Berhasil Diperbarui",
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id parameter is required",
		})
		return
	}

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id parameter must be a valid number",
		})
		return
	}

	err = uc.userService.DeleteUser(uint(id), c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data User Berhasil Dihapus",
	})
}
