package handlers

import (
	"net/http"
	"strconv"

	"quiz-base-api/internal/db"
	"quiz-base-api/internal/models"
	"quiz-base-api/internal/repos"
	"quiz-base-api/internal/services"

	"github.com/gin-gonic/gin"
)

func GetQuizQuestions(c *gin.Context) {
	quizID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuizQuestionRepository(db)
	service := services.NewQuizQuestionService(repo)

	items, err := service.ListQuestions(c.Request.Context(), quizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func CreateQuizQuestion(c *gin.Context) {
	quizID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var dto models.CreateQuestionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuizQuestionRepository(db)
	service := services.NewQuizQuestionService(repo)

	id, err := service.AddQuestion(c.Request.Context(), quizID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateQuizQuestion(c *gin.Context) {
	questionID, _ := strconv.ParseUint(c.Param("questionId"), 10, 64)
	var dto models.CreateQuestionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuizQuestionRepository(db)
	service := services.NewQuizQuestionService(repo)

	err = service.UpdateQuestion(c.Request.Context(), questionID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteQuizQuestion(c *gin.Context) {
	questionID, _ := strconv.ParseUint(c.Param("questionId"), 10, 64)

	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuizQuestionRepository(db)
	service := services.NewQuizQuestionService(repo)

	err = service.DeleteQuestion(c.Request.Context(), questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
