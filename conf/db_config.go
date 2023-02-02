package conf

import (
	"btc-billionaire/models"

	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Record{})

	DB = database

	// err = os.Remove("test.db")
	//
	//	if err != nil {
	//			log.Fatalf("failed to remove database file: %v", err)
	//	}
}
