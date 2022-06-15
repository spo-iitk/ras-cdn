package config

import (
	"log"
	"os"
)

func logConfig() {
	f, err := os.OpenFile("cdn.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(f)
}
