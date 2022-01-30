package main

import (
	"github.com/Valgard/godotenv"
)

/*
	The path should from the current run code path.
	If you run the path is at the learnGo path, the path is "./pkgusages/godotenv/.env"
	If you run the path is at the godotenv path, the path is ".env"
*/
var path = "./pkgusages/godotenv/.env"
var extraPath = "./pkgusages/godotenv/.env.dev"

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
