package main

import "github.com/gin-gonic/gin"

func (app *Application) ConfigureClientRoutes(r *gin.Engine) {
	r.GET("/", app.HomeHandler)
	r.GET("/create-bucket", app.CreateBucketHandler)
	r.POST("/create-bucket", app.CreateBucketPostHandler)

	r.GET("/upload-files", app.UploadFilesHandler)
	r.POST("/upload-files", app.UploadFilesPostHandler)

	r.GET("/list-files", app.ListFilesHandler)
	r.POST("/list-files", app.ListFilesPostHandler)

	r.GET("/download-file", app.DownloadLastUploadedFileHandler)
	r.POST("/download-file", app.DownloadLastUploadedFilePostHandler)

	r.POST("/check-bucket-name", app.CheckBucketname)
}

// DATABASE_URL="host=localhost port=5432 dbname=s3 user=postgres password=Matwyenko1_ sslmode=disable"
// AWS_ACCESS_KEY_ID=AKIA4JRYNABS6LDXDOKD
// AWS_SECRET_ACCESS_KEY=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
