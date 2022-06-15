package utils

import (
	"log"
	"os"
)

func DeleteFile(path string) bool {
	err := os.Remove(path)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
