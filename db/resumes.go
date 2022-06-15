package db

import (
	"log"

	"github.com/abhishekshree/cdn/models"
	"github.com/spf13/viper"
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
