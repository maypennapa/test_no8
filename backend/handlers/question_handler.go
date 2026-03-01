package handlers

import (
	"exam-api/database"
	"exam-api/models"

	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {
	var q models.Question
	c.BindJSON(&q)

	var count int64
	database.DB.Model(&models.Question{}).Count(&count)
	q.Number = int(count) + 1

	database.DB.Create(&q)
	c.JSON(200, q)
}

func GetQuestions(c *gin.Context) {
	var q []models.Question
	database.DB.Order("number asc").Find(&q)
	c.JSON(200, q)
}

func DeleteQuestion(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Question{}, id).Error; err != nil {
		c.JSON(400, gin.H{"error": "ไม่สามารถลบคำถามได้", "details": err.Error()})
		return
	}

	var questions []models.Question
	if err := database.DB.Order("number asc").Find(&questions).Error; err != nil {
		c.JSON(500, gin.H{"error": "ไม่สามารถดึงคำถามได้", "details": err.Error()})
		return
	}

	for i, q := range questions {
		newNumber := i + 1
		if q.Number != newNumber {
			database.DB.Model(&q).Update("number", newNumber)
		}
	}

	c.JSON(200, gin.H{"message": "ลบคำถามเรียบร้อย"})
}
