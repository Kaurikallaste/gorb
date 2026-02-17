package utils

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Env struct {
	PORT string
}

func GetEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	return &Env{
		PORT: os.Getenv("PORT"),
	}
}
