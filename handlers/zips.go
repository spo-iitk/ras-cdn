package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-cdn/db"
	"github.com/spo-iitk/ras-cdn/utils"
)

type ZipRequest struct {
	Files   []string `json:"files"`
	OutFile string   `json:"outfile"`
}

func ZipFilesHandler(ctx *gin.Context) {
	var req ZipRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if f := db.CheckFilesZipExists(req.Files); f != "" {
		ctx.JSON(200, gin.H{
			"message":  "zipped",
			"filename": f,
		})
		return
	}

	uuid, err := utils.GenerateUUID()
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	x := uuid + sep + req.OutFile

	go func() {
		utils.ZipFiles(req.Files, x)
	}()

	ctx.JSON(200, gin.H{
		"message":  "zipped",
		"filename": uuid + sep + req.OutFile,
	})
}

func DownloadZipHandler(ctx *gin.Context) {
	filename := ctx.Param("filename")
	db.UpdateAccessedAt(filename)

	ctx.File(upload_folder + "/" + filename)
}

type DeleteZipRequest struct {
	Filename string `json:"filename"`
}

func DeleteOneZipHandler(ctx *gin.Context) {
	var req DeleteZipRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	filename := req.Filename

	db.DeleteZipRow(filename)

	ok := utils.DeleteFile(upload_folder + "/" + filename)
	if !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "could not delete file"})
		return
	}
}

func DeleteZipsHandler(ctx *gin.Context) {
	utils.DeleteZips()
	ctx.JSON(200, gin.H{
		"message": "zips cleared",
	})
}
