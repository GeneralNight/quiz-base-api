package config

import "os"

type Config struct {
	AppPort string
	DSN     string
}

func Load() Config {
	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("DB_PORT", "3306")
	user := getenv("DB_USER", "root")
	pass := getenv("DB_PASS", "secret")
	name := getenv("DB_NAME", "quizdb")
	params := getenv("DB_PARAMS", "parseTime=true&charset=utf8mb4&loc=Local")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?" + params

	return Config{
		AppPort: getenv("APP_PORT", "8080"),
		DSN:     dsn,
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
