package controllers

import (
	"github.com/gin-gonic/gin"
	"go-interview/src/models"
	"gorm.io/gorm"
	"net/http"
)

type CompanyController struct {
	db *gorm.DB
}

func (uc *CompanyController) GetAllCompany(c *gin.Context) {
	var companies []models.Company
	rs := uc.db.Find(&companies)
	if rs.Error != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	c.JSON(http.StatusOK, companies)
}

func NewCompanyController(db *gorm.DB) CompanyController {
	return CompanyController{
		db: db,
	}
}
