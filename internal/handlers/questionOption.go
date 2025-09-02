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

func GetQuestionOptions(c *gin.Context) {
	questionID, _ := strconv.ParseUint(c.Param("questionId"), 10, 64)
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuestionOptionRepository(db)
	service := services.NewQuestionOptionService(repo)

	items, err := service.ListOptions(c.Request.Context(), questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func CreateQuestionOption(c *gin.Context) {
	questionID, _ := strconv.ParseUint(c.Param("questionId"), 10, 64)
	var dto models.CreateOptionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuestionOptionRepository(db)
	service := services.NewQuestionOptionService(repo)

	id, err := service.AddOption(c.Request.Context(), questionID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateQuestionOption(c *gin.Context) {
	optionID, _ := strconv.ParseUint(c.Param("optionId"), 10, 64)
	var dto models.CreateOptionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuestionOptionRepository(db)
	service := services.NewQuestionOptionService(repo)

	err = service.UpdateOption(c.Request.Context(), optionID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteQuestionOption(c *gin.Context) {
	optionID, _ := strconv.ParseUint(c.Param("optionId"), 10, 64)
	db, err := db.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repo := repos.NewQuestionOptionRepository(db)
	service := services.NewQuestionOptionService(repo)

	err = service.DeleteOption(c.Request.Context(), optionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
