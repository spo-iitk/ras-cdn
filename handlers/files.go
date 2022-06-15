package handlers

import (
	"log"

	"github.com/abhishekshree/cdn/db"
	"github.com/abhishekshree/cdn/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func UploadFileHandler(ctx *gin.Context) {
	MAX_SIZE := viper.GetInt64("MAX_SIZE")
	upload_folder := viper.GetString("FOLDERS.CDN")
	allowed_filetypes := viper.GetString("ALLOWED_FILETYPES")
	sep := viper.GetString("SEP")

	uid := ctx.GetHeader("token")

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
	upload_folder := viper.GetString("FOLDERS.CDN")

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
	upload_folder := viper.GetString("FOLDERS.CDN")

	filename := ctx.Param("filename")
	ctx.File(upload_folder + "/" + filename)
}

type DeleteRequest struct {
	Filename string `json:"filename"`
}

func DeleteFileHandler(ctx *gin.Context) {
	upload_folder := viper.GetString("FOLDERS.CDN")

	var req DeleteRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	//TODO: can add some secret code before deletion later for security
	filename := req.Filename

	ok := utils.DeleteFile(upload_folder + "/" + filename)
	if !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Could not delete file"})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "deleted",
	})
}
