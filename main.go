package main

import (
	"log"
	"os"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"weatherLineBot/weather"
	"github.com/joho/godotenv"
	//"weatherLineBot/test"
)

func main() {
	err := godotenv.Load(fmt.Sprintf(".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	result, err := weather.GetWeather()

	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(result)

	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}