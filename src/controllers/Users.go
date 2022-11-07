package controllers

import (
	"github.com/gin-gonic/gin"
	"go-interview/src/models"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	db *gorm.DB
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	var users []models.User
	rs := uc.db.Find(&users)
	if rs.Error != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user.EncryptPassword()

	result := uc.db.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error occurred")
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) LoginUser(c *gin.Context) {
	loginDto := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{}
	result := uc.db.Where("email = ?", loginDto.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, "user not found")
		return
	}

	if !user.CheckPassword(loginDto.Password) {
		c.JSON(http.StatusBadRequest, "wrong password")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"user": user, "token": user.CreateClaim()})

}

func NewUserController(db *gorm.DB) UserController {
	return UserController{
		db: db,
	}
}
