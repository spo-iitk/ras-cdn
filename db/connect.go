package db

import (
	"log"

	"github.com/abhishekshree/cdn/models"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	name := viper.GetString("DATABASE.NAME")

	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	err = db.AutoMigrate(&models.Uploads{}, &models.Zips{})
	if err != nil {
		log.Fatalf("Error migrating database: %v\n", err)
	}

	DB = db

	log.Println("Connected to database")
}

func init() {
	Connect()
}
