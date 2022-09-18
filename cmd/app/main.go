package main

import (
	"kinobot/internal/app"
	"kinobot/internal/config"
	"log"
)

func main() {
	log.Println("config init")
	cfg := config.GetConfig()

	mainApp, err := app.NewApp(cfg)

	if err != nil {
		log.Fatal(err)
	}
	mainApp.Run()
}
