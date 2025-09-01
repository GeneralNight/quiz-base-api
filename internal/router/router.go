package router

import (
	"github.com/GeneralNight/quiz-base-api/internal/quiz"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	// healthcheck
	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })

	quizGroup := r.Group("/quiz")
	{
		quizGroup.GET("", quiz.GetQuiz)
		quizGroup.GET("/:id", quiz.GetQuizById)
		quizGroup.POST("", quiz.CreateQuiz)
		quizGroup.PUT("/:id", quiz.UpdateQuiz)
		quizGroup.DELETE("/:id", quiz.DeleteQuiz)
		quizGroup.PUT("/:id/status", quiz.UpdateQuizStatus)

		quizGroup.GET("/:id/questions", quiz.GetQuizQuestions)
		quizGroup.POST("/:id/questions", quiz.CreateQuizQuestion)
		quizGroup.PUT("/questions/:questionId", quiz.UpdateQuizQuestion)
		quizGroup.DELETE("/questions/:questionId", quiz.DeleteQuizQuestion)

		quizGroup.GET("/questions/:questionId/options", quiz.GetQuestionOptions)
		quizGroup.POST("/questions/:questionId/options", quiz.CreateQuestionOption)
		quizGroup.PUT("/options/:optionId", quiz.UpdateQuestionOption)
		quizGroup.DELETE("/options/:optionId", quiz.DeleteQuestionOption)
	}
	return r
}
