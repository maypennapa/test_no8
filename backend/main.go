package main

import (
	"exam-api/database"
	"exam-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDB()

	r := gin.Default()

	r.GET("/questions", handlers.GetQuestions)
	r.POST("/questions", handlers.CreateQuestion)
	r.DELETE("/questions/:id", handlers.DeleteQuestion)

	r.Run(":8080")
}
