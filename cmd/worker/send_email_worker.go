package main

import (
	"isaevfeed/notifier/internal/notifier"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	noti := notifier.New(true)

	if err := noti.SendEmail(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
