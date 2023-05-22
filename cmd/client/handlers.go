package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Response struct {
	Objects Objects `json:"Objects"`
}

type Objects struct {
	CommonPrefixes interface{} `json:"CommonPrefixes"`
	Contents       []Contents  `json:"Contents"`
	Delimiter      interface{} `json:"Delimiter"`
	EncodingType   interface{} `json:"EncodingType"`
	IsTruncated    bool        `json:"IsTruncated"`
	Marker         string      `json:"Marker"`
	MaxKeys        int         `json:"MaxKeys"`
	Name           string      `json:"Name"`
	NextMarker     interface{} `json:"NextMarker"`
	Prefix         string      `json:"Prefix"`
}

type Contents struct {
	ChecksumAlgorithm interface{} `json:"ChecksumAlgorithm"`
	ETag              string      `json:"ETag"`
	Key               string      `json:"Key"`
	LastModified      string      `json:"LastModified"`
	Owner             Owner       `json:"Owner"`
	Size              int         `json:"Size"`
	StorageClass      string      `json:"StorageClass"`
}

type Owner struct {
	DisplayName interface{} `json:"DisplayName"`
	ID          string      `json:"ID"`
}

func (app *Application) HomeHandler(c *gin.Context) {
	if err := app.renderTemplate(c, "home", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *Application) UploadFilesHandler(c *gin.Context) {
	if err := app.renderTemplate(c, "upload-files", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *Application) ListFilesHandler(c *gin.Context) {
	if err := app.renderTemplate(c, "list-files", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *Application) DownloadLastUploadedFileHandler(c *gin.Context) {
	if err := app.renderTemplate(c, "download-files", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *Application) CreateBucketHandler(c *gin.Context) {
	if err := app.renderTemplate(c, "create-bucket", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *Application) CreateBucketPostHandler(c *gin.Context) {
	var payload struct {
		BucketName string `form:"bucketName"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	app.infoLog.Printf("Bucket name: %s", payload.BucketName)

	out, err := json.Marshal(payload)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	req, _ := http.NewRequest("PUT", "http://localhost:4001/buckets", bytes.NewBuffer(out))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	if response.Error {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": response.Message,
	})
}

func (app *Application) CheckBucketname(c *gin.Context) {
	var payload struct {
		BucketName string `form:"bucketName"`
	}

	if err := c.ShouldBind(&payload); err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	out, err := json.Marshal(payload)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	req, _ := http.NewRequest("POST", "http://localhost:4001/check-bucket-name", bytes.NewBuffer(out))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Internal Server Error",
		})
		return
	}

	if response.Error {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   response.Error,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": response.Message,
	})
}

func (app *Application) UploadFilesPostHandler(c *gin.Context) {
	var payload struct {
		BucketName string `json:"bucketName"`
	}
	payload.BucketName = "my-bebra"

	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	bucketName := c.Request.FormValue("bucketName")

	fileData, err := io.ReadAll(file)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	payload.BucketName = bucketName

	requestBody := struct {
		BucketName string `json:"bucketName"`
		Files      []byte `json:"file"`
		Filename   string `json:"filename"`
	}{
		BucketName: payload.BucketName,
		Files:      fileData,
		Filename:   handler.Filename,
	}

	out, err := json.Marshal(requestBody)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:4001/buckets/"+payload.BucketName+"/files", bytes.NewBuffer(out))
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "failed to decode response",
		})
		return
	}

	if response.Error {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   response.Error,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": response.Message,
	})
}

func (app *Application) ListFilesPostHandler(c *gin.Context) {
	var payload struct {
		BucketName string `json:"bucketName"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	req, _ := http.NewRequest("GET", "http://localhost:4001/buckets/"+payload.BucketName+"/files", nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "failed to decode response",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": response,
	})
}

func (app *Application) DownloadLastUploadedFilePostHandler(c *gin.Context) {
	var payload struct {
		BucketName string `json:"bucketName"`
		ObjectKey  string `json:"objectKey"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	out, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "http://localhost:4001/buckets/"+payload.BucketName+"/files/"+payload.ObjectKey, bytes.NewBuffer(out))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// here stopped
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		app.errorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "failed to decode response",
		})
		return
	}

	if response.Error {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   response.Error,
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": response.Message,
	})
}
