package main

import (
	"isaevfeed/notifier/internal/server"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	srv := server.New()
	srv.Start()
}
