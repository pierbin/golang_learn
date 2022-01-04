package migrations

import (
	"learnGo/projects/mysqlSuiteTests/config"
	"learnGo/projects/mysqlSuiteTests/entity"
)

func getModels() []interface{} {
	return []interface{}{&entity.User{}}
}

func MigrateTable() {
	db := config.GetDB()
	db.AutoMigrate(getModels()...)
}
