package routes

import (
	"github.com/gin-gonic/gin"
	"go-interview/src/controllers"
	"go-interview/src/middlewares"
	"gorm.io/gorm"
)

func CreateRoutes(r *gin.Engine, db *gorm.DB) {
	userControllers := controllers.NewUserController(db)
	r.GET("users", userControllers.GetAllUsers)
	r.POST("register", userControllers.RegisterUser)
	r.POST("login", userControllers.LoginUser)

	companyControllers := controllers.NewCompanyController(db)
	r.GET("companies", companyControllers.GetAllCompany)

	interviewControllers := controllers.NewInterviewController(db)
	r.GET("interviews", middlewares.AuthRequired(interviewControllers.GetMyInterviews))
	r.POST("interviews", middlewares.AuthRequired(interviewControllers.CreateInterview))
}
