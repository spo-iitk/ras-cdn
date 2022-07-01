package handlers

import (
	"log"
	"net/http"

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
	secret            string
)

func init() {
	MAX_SIZE = viper.GetInt64("MAX_SIZE")
	upload_folder = viper.GetString("FOLDERS.CDN")
	zip_folder = viper.GetString("FOLDERS.ZIP")
	allowed_filetypes = viper.GetString("ALLOWED_FILETYPE")
	secret = viper.GetString("SECRET")
}

func UploadFileHandler(ctx *gin.Context) {
	uid := ctx.GetHeader("token")
	rid := ctx.GetHeader("rid")
	log.Println(MAX_SIZE)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if file.Size > MAX_SIZE {
		log.Printf("%v size is too large \n", file.Filename)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "File size is too large"})
		return
	}

	if file.Header.Get("Content-Type") != allowed_filetypes {
		log.Printf("%v is not allowed \n", file.Header.Get("Content-Type"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "File type is not allowed"})
		return
	}

	uuid, err := utils.GenerateUUID()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newname := utils.GenerateName(uuid, file.Filename, rid)
	ctx.SaveUploadedFile(file, upload_folder+"/"+newname)
	log.Println(file.Filename + " to " + newname)

	db.CreateUpload(uid, newname)

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "uploaded",
		"filename": newname,
	})
}

func ViewAllHandler(ctx *gin.Context) {
	files, err := utils.ListFiles(upload_folder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	viewAllSecret := ctx.GetHeader("secret")
	if viewAllSecret != secret {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

func ViewFileHandler(ctx *gin.Context) {
	filename := ctx.Param("filename")
	ctx.File(upload_folder + "/" + filename)
}

type DeleteRequest struct {
	Filename string `json:"filename"`
	Secret   string `json:"secret"`
}

func DeleteFileHandler(ctx *gin.Context) {
	var req DeleteRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Works for now
	if req.Secret != secret {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	filename := req.Filename

	db.DeleteUpload(filename)

	ok := utils.DeleteFile(upload_folder + "/" + filename)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Could not delete file"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}
