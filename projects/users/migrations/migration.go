package migrations

import (
	"learnGo/projects/users/config"
	"learnGo/projects/users/entity"
)

func getModels() []interface{} {
	return []interface{}{&entity.User{}}
}

func MigrateTable() {
	db := config.GetDB()
	db.AutoMigrate(getModels()...)
}
