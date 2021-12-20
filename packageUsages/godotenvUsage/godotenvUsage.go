package main

import (
	"github.com/Valgard/godotenv"
)

func main() {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}

	// You can also load several files
	if err := dotenv.Load(".env", ".env.dev"); err != nil {
		panic(err)
	}
}
