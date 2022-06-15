package utils

import (
	"archive/zip"
	"io"
	"log"
	"os"

	"github.com/abhishekshree/cdn/db"
	"github.com/spf13/viper"
)

var zipdir string
var cdndir string

func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(cdndir + "/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		return err
	}
	if _, err := io.Copy(wr, file); err != nil {
		return err
	}
	return nil
}

func ZipFiles(files []string, outfile string) (string, error) {
	sep := viper.GetString("SEP")
	zipdir = viper.GetString("FOLDERS.ZIP")
	cdndir = viper.GetString("FOLDERS.CDN")

	uuid, err := GenerateUUID()
	if err != nil {
		return "", err
	}
	x := uuid + sep + outfile
	x = zipdir + "/" + x

	out, err := os.Create(x)

	if err != nil {
		return "", err
	}
	defer out.Close()

	zipw := zip.NewWriter(out)
	defer zipw.Close()

	for _, filename := range files {
		if err := appendFiles(filename, zipw); err != nil {
			os.Remove(x)
			return "", err
		}
	}

	db.CreateZip(files, uuid+sep+outfile)

	return uuid + sep + outfile, nil
}

func DeleteZips() {
	zipdir = viper.GetString("FOLDERS.ZIP")

	// delete everything in the zip folder

	files, err := ListFiles(zipdir)

	if err != nil {
		log.Println(err)
		return
	}

	for _, filename := range files {
		db.DeleteZipRow(filename)
		os.Remove(zipdir + "/" + filename)
	}
}
