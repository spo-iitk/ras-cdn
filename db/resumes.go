package db

import (
	"log"

	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-cdn/models"
)

func CreateUpload(uid string, name string) {
	viewurl := viper.GetString("VIEW_URL")

	tx := DB.Create(&models.Uploads{
		UserID: uid,
		Name:   name,
		URL:    viewurl + name,
	})

	if tx.Error != nil {
		log.Println(tx.Error)
	}
}

func DeleteUpload(filename string) {
	tx := DB.Where("name = ?", filename).Delete(&models.Uploads{})

	if tx.Error != nil {
		log.Println(tx.Error)
	}
}
