package test

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func Test() string {
	err := godotenv.Load(fmt.Sprintf(".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("LINE_BOT_CHANNEL_TOKEN")
	return env
}