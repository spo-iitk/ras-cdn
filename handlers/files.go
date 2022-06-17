package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "github.com/spo-iitk/ras-cdn/config"
	"github.com/spo-iitk/ras-cdn/db"
	"github.com/spo-iitk/ras-cdn/utils"
)

var (
	MAX_SIZE          int64
	upload_folder     string
	zip_folder        string
	allowed_filetypes string
	sep               string
)

func init() {
	MAX_SIZE = viper.GetInt64("MAX_SIZE")
	upload_folder = viper.GetString("FOLDERS.CDN")
	zip_folder = viper.GetString("FOLDERS.ZIP")
	allowed_filetypes = viper.GetString("ALLOWED_FILETYPE")
	sep = viper.GetString("SEP")
}

func UploadFileHandler(ctx *gin.Context) {
	uid := ctx.GetHeader("token")
	log.Println(MAX_SIZE)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if file.Size > MAX_SIZE {
		log.Printf("%v size is too large \n", file.Filename)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "File size is too large"})
		return
	}

	if file.Header.Get("Content-Type") != allowed_filetypes {
		log.Printf("%v is not allowed \n", file.Header.Get("Content-Type"))
		ctx.AbortWithStatusJSON(400, gin.H{"error": "File type is not allowed"})
		return
	}

	uuid, err := utils.GenerateUUID()
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	newname := uuid + sep + file.Filename
	ctx.SaveUploadedFile(file, upload_folder+"/"+newname)
	log.Println(file.Filename + " to " + newname)

	db.CreateUpload(uid, newname)

	ctx.JSON(200, gin.H{
		"message":  "uploaded",
		"filename": newname,
	})
}

func ViewAllHandler(ctx *gin.Context) {
	files, err := utils.ListFiles(upload_folder)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"files": files,
	})
}

func ViewFileHandler(ctx *gin.Context) {
	filename := ctx.Param("filename")
	ctx.File(upload_folder + "/" + filename)
}

type DeleteRequest struct {
	Filename string `json:"filename"`
}

func DeleteFileHandler(ctx *gin.Context) {
	var req DeleteRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	//TODO: can add some secret code before deletion later for security
	filename := req.Filename

	db.DeleteUpload(filename)

	ok := utils.DeleteFile(upload_folder + "/" + filename)
	if !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Could not delete file"})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "deleted",
	})
}
