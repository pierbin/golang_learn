package main

import (
	"fmt"
	"net/http"
	"os"

	"learnGo/projects/mysqlSuiteTests/config"
	"learnGo/projects/mysqlSuiteTests/controllers"
	"learnGo/projects/mysqlSuiteTests/migrations"

	"github.com/Valgard/godotenv"
	"github.com/gin-gonic/gin"
)

var defaultPort = "8080"
var path = ".env"

func init() {
	dotenv := godotenv.New()
	fmt.Println(dotenv)
	if err := dotenv.Load(path); err != nil {
		panic(err)
	}
}

/**
	In mysqlSuiteTests folder, run.

	% go test ./tests
	ok      learnGo/projects/mysqlSuiteTests/tests  0.122s

	% go test -v ./tests
	...
	--- PASS: TestSuite (0.11s)
    --- PASS: TestSuite/TestCreateUser (0.01s)
	PASS
	ok      learnGo/projects/mysqlSuiteTests/tests  0.131s

	% go run server.go
	...
	[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
	 - using env:   export GIN_MODE=release
	 - using code:  gin.SetMode(gin.ReleaseMode)

	[GIN-debug] GET    /                         --> main.main.func1 (2 handlers)
	[GIN-debug] GET    /users                    --> learnGo/projects/mysqlSuiteTests/controllers.userControllerInterface.GetAll-fm (2 handlers)
	[GIN-debug] POST   /users                    --> learnGo/projects/mysqlSuiteTests/controllers.userControllerInterface.Create-fm (2 handlers)
	[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
	[GIN-debug] Listening and serving HTTP on :8080
*/
func main() {
	config.ConnectGorm()
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	migrations.MigrateTable()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// Testing Port
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/users", controllers.UserController.GetAll)
	router.POST("/users", controllers.UserController.Create)

	router.Run(":" + port)
}
