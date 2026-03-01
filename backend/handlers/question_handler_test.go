package handlers

import (
	"bytes"
	"encoding/json"
	"exam-api/database"
	"exam-api/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup test DB
func TestMain(m *testing.M) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect test DB:", err)
		os.Exit(1)
	}
	db.AutoMigrate(&models.Question{})
	database.DB = db

	os.Exit(m.Run())
}

func TestCreateQuestion(t *testing.T) {
	gin.SetMode(gin.TestMode)

	database.DB.Exec("DELETE FROM questions")

	q := models.Question{
		Question: "ข้อสอบ 1",
		Choice1:  "A",
		Choice2:  "B",
		Choice3:  "C",
		Choice4:  "D",
	}

	jsonQ, _ := json.Marshal(q)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/questions", bytes.NewBuffer(jsonQ))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateQuestion(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var created models.Question
	json.Unmarshal(w.Body.Bytes(), &created)

	if created.Number != 1 {
		t.Fatalf("Expected question number 1, got %d", created.Number)
	}
	if created.Question != "ข้อสอบ 1" {
		t.Fatalf("Expected question text 'ข้อสอบ 1', got '%s'", created.Question)
	}

	t.Cleanup(func() {
		database.DB.Exec("DELETE FROM questions")
	})
}

func TestGetQuestions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	database.DB.Exec("DELETE FROM questions")
	q := models.Question{
		Question: "ข้อสอบ 2",
		Choice1:  "A",
		Choice2:  "B",
		Choice3:  "C",
		Choice4:  "D",
	}
	database.DB.Create(&q)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/questions", nil)

	GetQuestions(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var questions []models.Question
	json.Unmarshal(w.Body.Bytes(), &questions)
	if len(questions) == 0 {
		t.Fatalf("Expected at least 1 question, got 0")
	}

	t.Cleanup(func() {
		database.DB.Exec("DELETE FROM questions")
	})
}

func TestDeleteQuestion(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := database.DB
	db.Exec("DELETE FROM questions")

	question := models.Question{
		Question: "ข้อสอบ ลบ",
		Choice1:  "A",
		Choice2:  "B",
		Choice3:  "C",
		Choice4:  "D",
	}
	db.Create(&question)
	questionID := question.ID

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: fmt.Sprint(questionID)}}
	c.Request, _ = http.NewRequest("DELETE", "/questions/"+fmt.Sprint(questionID), nil)

	DeleteQuestion(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var count int64
	db.Model(&models.Question{}).Count(&count)
	if count != 0 {
		t.Fatalf("Expected 0 questions after delete, got %d", count)
	}

	t.Cleanup(func() {
		db.Exec("DELETE FROM questions")
	})
}
