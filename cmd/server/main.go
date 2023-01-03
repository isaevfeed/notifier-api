package main

import (
	"fmt"
	"isaevfeed/notifier/internal/server"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	host, _ := os.LookupEnv("SERVER_HOST")
	port, _ := os.LookupEnv("SERVER_PORT")
	srv := server.New(fmt.Sprintf("%s:%s", host, port))
	srv.Start()
}
