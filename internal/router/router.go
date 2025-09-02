package router

import (
	"quiz-base-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	// healthcheck
	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })
	quizGroup := r.Group("/quiz")
	{
		quizGroup.GET("", handlers.GetQuiz)
		quizGroup.GET("/:id", handlers.GetQuizById)
		quizGroup.POST("", handlers.CreateQuiz)
		quizGroup.PUT("/:id", handlers.UpdateQuiz)
		quizGroup.DELETE("/:id", handlers.DeleteQuiz)
		quizGroup.PUT("/:id/status", handlers.UpdateQuizStatus)

		quizGroup.GET("/:id/questions", handlers.GetQuizQuestions)
		quizGroup.POST("/:id/questions", handlers.CreateQuizQuestion)
		quizGroup.PUT("/questions/:questionId", handlers.UpdateQuizQuestion)
		quizGroup.DELETE("/questions/:questionId", handlers.DeleteQuizQuestion)

		quizGroup.GET("/questions/:questionId/options", handlers.GetQuestionOptions)
		quizGroup.POST("/questions/:questionId/options", handlers.CreateQuestionOption)
		quizGroup.PUT("/options/:optionId", handlers.UpdateQuestionOption)
		quizGroup.DELETE("/options/:optionId", handlers.DeleteQuestionOption)
	}
	return r
}
