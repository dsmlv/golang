package migrations

import (
	"final/config"
	"final/models"
)

func RunMigrations() {
	db := config.DB
	models.Migrate(db)
}
