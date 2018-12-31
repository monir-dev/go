package main

import (
	"Structure/src/system/app"
	"flag"
	"os"
	"github.com/joho/godotenv"
)

var port string

func init() {
	flag.StringVar(&port, "port", "4000", "Assigning the port that the server should listen to.")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if envPort > 0 {
		port = envPort
	}
}

func main() {
	s := app.NewServer()

	s.Init(port)
	s.Start()
}
