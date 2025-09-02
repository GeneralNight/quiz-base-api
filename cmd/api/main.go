package main

import (
	"log"

	"github.com/joho/godotenv"

	"quiz-base-api/internal/config"
	"quiz-base-api/internal/db"
	"quiz-base-api/internal/router"
)

func main() {
	_ = godotenv.Load() // opcional

	cfg := config.Load()

	conn, err := db.OpenMySQL(cfg)
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	defer conn.Close()

	r := router.SetupRoutes()

	addr := ":" + cfg.AppPort
	log.Println("listening on", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
