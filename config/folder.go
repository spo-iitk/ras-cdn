package config

import "os"

func initFolder() {
	x := "cdn"
	if _, err := os.Stat(x); os.IsNotExist(err) {
		os.Mkdir(x, 0755)
	}

	x = "zip"
	if _, err := os.Stat(x); os.IsNotExist(err) {
		os.Mkdir(x, 0755)
	}
}
