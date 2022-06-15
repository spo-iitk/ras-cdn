package main

import (
	_ "github.com/abhishekshree/cdn/config"
	"github.com/abhishekshree/cdn/handlers"
	"github.com/abhishekshree/cdn/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORS())

	fs := r.Group("/api/")
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

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
