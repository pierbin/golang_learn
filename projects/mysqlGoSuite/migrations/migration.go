package migrations

import (
	"learnGo/projects/mysqlGoSuite/config"
	"learnGo/projects/mysqlGoSuite/entity"
)

func getModels() []interface{} {
	return []interface{}{&entity.User{}}
}

func MigrateTable() {
	db := config.GetDB()
	db.AutoMigrate(getModels()...)
}
