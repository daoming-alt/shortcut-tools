package database

import (
	"github.com/kdbrian/shortcut-tool/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() error {

	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return migrate()

}

func migrate() error {
	return DB.AutoMigrate(&models.Shortcut{})
}
