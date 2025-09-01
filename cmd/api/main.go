package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/GeneralNight/quiz-base-api/internal/config"
	"github.com/GeneralNight/quiz-base-api/internal/db"
	"github.com/GeneralNight/quiz-base-api/internal/quiz"
	"github.com/GeneralNight/quiz-base-api/internal/router"
)

func main() {
	_ = godotenv.Load() // opcional

	cfg := config.Load()

	conn, err := db.OpenMySQL(cfg)
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	defer conn.Close()

	quizRepo := quiz.NewRepository(conn)
	quizSvc := quiz.NewService(quizRepo)
	quizHandler := quiz.NewHandler(quizSvc)

	r := router.New()
	quiz.RegisterRoutes(r, quizHandler)

	addr := ":" + cfg.AppPort
	log.Println("listening on", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
