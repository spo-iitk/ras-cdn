package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-cdn/db"
	"github.com/spo-iitk/ras-cdn/utils"
)

type ZipRequest struct {
	Files   []string `json:"files"`
	Rid     string   `json:"rid"`
	OutFile string   `json:"outfile"`
}

func ZipFilesHandler(ctx *gin.Context) {
	var req ZipRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if f := db.CheckFilesZipExists(req.Files); f != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "zipped",
			"filename": f,
		})
		return
	}

	uuid, err := utils.GenerateUUID()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	x := utils.GenerateName(uuid, req.OutFile, req.Rid)

	go func() {
		utils.ZipFiles(req.Files, x)
	}()

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "zipped",
		"filename": x,
	})
}

func DownloadZipHandler(ctx *gin.Context) {
	filename := ctx.Param("filename")
	db.UpdateAccessedAt(filename)

	ctx.File(zip_folder + "/" + filename)
}

type DeleteZipRequest struct {
	Filename string `json:"filename"`
	Secret   string `json:"secret"`
}

func DeleteOneZipHandler(ctx *gin.Context) {
	var req DeleteZipRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	filename := req.Filename

	if req.Secret != secret {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	db.DeleteZipRow(filename)

	ok := utils.DeleteFile(zip_folder + "/" + filename)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "could not delete file"})
		return
	}
}

func DeleteZipsHandler(ctx *gin.Context) {
	delete_secret := ctx.GetHeader("secret")
	if delete_secret != secret {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	utils.DeleteZips()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "zips cleared",
	})
}
