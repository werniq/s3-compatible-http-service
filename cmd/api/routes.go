package main

import "github.com/gin-gonic/gin"

func (app *Application) ConfigureRoutes(r *gin.Engine) {
	r.PUT("/buckets", app.CreateNewBucket)
	r.POST("/buckets/:bucket-name/files", app.UploadFiles)
	r.GET("/buckets/:bucket-name/files", app.ListFiles)
	r.POST("/buckets/:bucket-name/files/:filename", app.DownloadFile)
	r.POST("/check-bucket-name", app.CheckBucketName)
}
