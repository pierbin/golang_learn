package mysqlGoSuite

import (
	"learnGo/projects/mysqlGoSuite/config"
	"learnGo/projects/mysqlGoSuite/controllers"
	"learnGo/projects/mysqlGoSuite/migrations"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var defaultPort = "8080"

func init() {
	godotenv.Load()
}

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
