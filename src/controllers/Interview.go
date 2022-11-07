package controllers

import (
	"github.com/gin-gonic/gin"
	"go-interview/src/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"time"
)

type InterviewController struct {
	db *gorm.DB
}

func (ic *InterviewController) GetMyInterviews(c *gin.Context) {
	uId, _ := c.Get("userId")
	userId, _ := strconv.Atoi(uId.(string))
	var interviews []models.Interviews
	ic.db.Where("user_id = ?", userId).Preload(clause.Associations).Find(&interviews)

	if interviews != nil {
		c.JSON(http.StatusOK, interviews)
		return
	}
	c.JSON(http.StatusOK, []int{})
}

func (ic InterviewController) CreateInterview(c *gin.Context) {
	uId, _ := c.Get("userId")
	userId, _ := strconv.Atoi(uId.(string))
	interviewDto := struct {
		CompanyId     int       `json:"company_id"`
		InterviewDate time.Time `json:"interview_date"`
	}{}

	err := c.ShouldBindJSON(&interviewDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	interview := models.Interviews{
		UserId:        userId,
		CompanyId:     interviewDto.CompanyId,
		InterviewDate: interviewDto.InterviewDate,
		Status:        models.Ongoing,
	}

	result := ic.db.Create(&interview)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "internal server error")
	}

	c.JSON(http.StatusCreated, interview)

}

func NewInterviewController(db *gorm.DB) InterviewController {
	return InterviewController{
		db: db,
	}
}
