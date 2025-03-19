package database

import (
	"github.com/kdbrian/shortcut-tool/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func Connect() error {
	db, err := gorm.Open(sqlite.Open("shortcut.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return migrate()

}

func migrate() error {
	return DB.AutoMigrate(&models.Shortcut{})
}
