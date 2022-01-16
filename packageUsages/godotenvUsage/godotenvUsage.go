package main

import (
	"github.com/Valgard/godotenv"
)

/*
	The path should from the current run code path.
	If you run the path is at the learnGo path, the path is "./packageUsages/godotenvUsage/.env"
	If you run the path is at the godotenvUsage path, the path is ".env"
*/
var path = "./packageUsages/godotenvUsage/.env"
var extraPath = "./packageUsages/godotenvUsage/.env.dev"

func main() {
	dotenv := godotenv.New()
	if err := dotenv.Load(path); err != nil {
		panic(err)
	}

	// You can also load several files
	// if err := dotenv.Load(extraPath); err != nil {
	//	panic(err)
	// }
}
