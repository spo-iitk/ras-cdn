package db

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/abhishekshree/cdn/models"
	"github.com/spf13/viper"
)

func CreateZip(files []string, outfile string) {
	tx := DB.Create(&models.Zips{
		Files:      strings.Join(files, ","),
		OutFile:    outfile,
		AccessedAt: time.Now(),
	})

	if tx.Error != nil {
		log.Println(tx.Error)
	}
}

func CheckFilesZipExists(files []string) string {
	var zip models.Zips
	tx := DB.Where("files = ?", strings.Join(files, ",")).First(&zip)

	if tx.Error != nil {
		log.Println(tx.Error)
	}

	if zip.ID != 0 {
		return zip.OutFile
	}

	return ""
}

func DeleteZipRow(filename string) {
	tx := DB.Where("out_file = ?", filename).Unscoped().Delete(&models.Zips{})

	if tx.Error != nil {
		log.Println(tx.Error)
	}
}

func UpdateAccessedAt(filename string) {
	tx := DB.Model(&models.Zips{}).Where("out_file = ?", filename).Update("accessed_at", time.Now())

	if tx.Error != nil {
		log.Println(tx.Error)
	}
}

func CleanupZips() {
	zipdir := viper.GetString("FOLDERS.ZIP")

	for {
		time.Sleep(time.Hour * 24)
		var zips []models.Zips
		tx := DB.Unscoped().Where("accessed_at > ?", time.Now().Add(-24*time.Hour)).Find(&zips)

		if tx.Error != nil {
			log.Println(tx.Error)
		}

		for _, zip := range zips {
			DeleteZipRow(zip.OutFile)
			os.Remove(zipdir + "/" + zip.OutFile)
		}
	}
}
