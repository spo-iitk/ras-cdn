package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "github.com/spo-iitk/ras-cdn/config"
	"github.com/spo-iitk/ras-cdn/handlers"
	"github.com/spo-iitk/ras-cdn/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORS())

	fs := r.Group("/cdn/")
	{
		fs.POST("/upload", handlers.UploadFileHandler)
		fs.DELETE("/delete", handlers.DeleteFileHandler)
		fs.GET("/view/:filename", handlers.ViewFileHandler)
		fs.GET("/view/all", handlers.ViewAllHandler)
		fs.GET("/zip/:filename", handlers.DownloadZipHandler)
		fs.POST("/zip", handlers.ZipFilesHandler)
		fs.DELETE("/zip", handlers.DeleteOneZipHandler)
		fs.DELETE("/zip/all", handlers.DeleteZipsHandler)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	port := viper.GetString("PORT")

	if err := r.Run(port); err != nil {
		panic(err)
	}
}
