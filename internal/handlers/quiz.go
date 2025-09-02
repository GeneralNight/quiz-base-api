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

func GetQuiz(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuizRepository(db)
	service := services.NewQuizService(repo)

	items, err := service.ListQuizzes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func GetQuizById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuizRepository(db)
	service := services.NewQuizService(repo)

	item, err := service.GetQuiz(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "quiz not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func CreateQuiz(c *gin.Context) {
	var dto models.CreateQuizDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuizRepository(db)
	service := services.NewQuizService(repo)
	id, err := service.CreateQuiz(c.Request.Context(), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateQuiz(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var dto models.CreateQuizDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuizRepository(db)
	service := services.NewQuizService(repo)

	err = service.UpdateQuiz(c.Request.Context(), id, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteQuiz(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuizRepository(db)
	service := services.NewQuizService(repo)
	err = service.DeleteQuiz(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func UpdateQuizStatus(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var payload struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuizRepository(db)
	service := services.NewQuizService(repo)
	err = service.UpdateQuizStatus(c.Request.Context(), id, payload.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
