package main

import (
	"github.com/gin-gonic/gin"
	"go-interview/src/models"
	"go-interview/src/routes"
)

func main() {
	db, err := models.CreateSqliteDb("test.db")
	if err != nil {
		panic(err)
	}

	//err = models.MigrateToSqliteDb(db, []interface{}{models.User{}, models.Company{}, models.Interviews{}})
	//if err != nil {
	//	panic(err)
	//}

	r := gin.Default()
	routes.CreateRoutes(r, db)
	r.Run()
}
