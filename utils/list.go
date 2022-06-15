package utils

import "os"

func ListFiles(dir string) ([]string, error) {
	d, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}

func ListFileURLS(files []string) []string {
	x := "http://localhost:8080/api/view/"
	var fileURLS []string
	for _, file := range files {
		fileURLS = append(fileURLS, x+file)
	}
	return fileURLS
}
