package handlers

import (
	"exam-api/database"
	"exam-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {
	var q models.Question
	if err := c.BindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	var count int64
	if err := database.DB.Model(&models.Question{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot count questions", "details": err.Error()})
		return
	}
	q.Number = int(count) + 1

	if err := database.DB.Create(&q).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create question", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, q)
}

func GetQuestions(c *gin.Context) {
	var questions []models.Question
	if err := database.DB.Order("number asc").Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get questions", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, questions)
}

func DeleteQuestion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := database.DB.Delete(&models.Question{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete question", "details": err.Error()})
		return
	}

	var questions []models.Question
	if err := database.DB.Order("number asc").Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch questions", "details": err.Error()})
		return
	}

	for i, q := range questions {
		newNumber := i + 1
		if q.Number != newNumber {
			database.DB.Model(&q).Update("number", newNumber)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}
