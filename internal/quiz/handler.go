package quiz

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var svc Service // should be initialized in main or router

func GetQuiz(c *gin.Context) {
	items, err := svc.ListQuizzes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func GetQuizById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	item, err := svc.GetQuiz(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "quiz not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func CreateQuiz(c *gin.Context) {
	var dto CreateQuizDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := svc.CreateQuiz(c.Request.Context(), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateQuiz(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var dto CreateQuizDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := svc.UpdateQuiz(c.Request.Context(), id, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteQuiz(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := svc.DeleteQuiz(c.Request.Context(), id)
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
	err := svc.UpdateQuizStatus(c.Request.Context(), id, payload.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func GetQuizQuestions(c *gin.Context) {
	quizID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	items, err := svc.ListQuestions(c.Request.Context(), quizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func CreateQuizQuestion(c *gin.Context) {
	quizID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var dto CreateQuestionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := svc.AddQuestion(c.Request.Context(), quizID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateQuizQuestion(c *gin.Context) {
	questionID, _ := strconv.ParseUint(c.Param("questionId"), 10, 64)
	var dto CreateQuestionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := svc.UpdateQuestion(c.Request.Context(), questionID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteQuizQuestion(c *gin.Context) {
	questionID, _ := strconv.ParseUint(c.Param("questionId"), 10, 64)
	err := svc.DeleteQuestion(c.Request.Context(), questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func GetQuestionOptions(c *gin.Context) {
	questionID, _ := strconv.ParseUint(c.Param("questionId"), 10, 64)
	items, err := svc.ListOptions(c.Request.Context(), questionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func CreateQuestionOption(c *gin.Context) {
	questionID, _ := strconv.ParseUint(c.Param("questionId"), 10, 64)
	var dto CreateOptionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := svc.AddOption(c.Request.Context(), questionID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateQuestionOption(c *gin.Context) {
	optionID, _ := strconv.ParseUint(c.Param("optionId"), 10, 64)
	var dto CreateOptionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := svc.UpdateOption(c.Request.Context(), optionID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteQuestionOption(c *gin.Context) {
	optionID, _ := strconv.ParseUint(c.Param("optionId"), 10, 64)
	err := svc.DeleteOption(c.Request.Context(), optionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
