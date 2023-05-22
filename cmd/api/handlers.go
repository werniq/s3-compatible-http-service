package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type BucketCreateRequest struct {
	BucketName string `json:"bucketName"`
}

type BucketUploadRequest struct {
	Filename string `json:"file-name"`
}

// CreateNewBucket function creates a new bucket in AWS S3
// bucketName is the name of the bucket to be created
// returns an error if the bucket could not be created
func (app *Application) CreateNewBucket(c *gin.Context) {
	bodyBytes, err := c.GetRawData()
	if err != nil {
		app.ErrorLogger.Printf("error binding bucket name: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error binding bucket name: %v", err.Error()),
		})
		return
	}

	// TODO: validate bucket name
	// TODO: check if bucket name already exists in database

	var bucketName struct {
		BucketName string `json:"bucketName"`
	}

	err = json.Unmarshal(bodyBytes, &bucketName)
	if err != nil {
		app.ErrorLogger.Printf("error unmarshalling bucket name: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error unmarshalling bucket name: %v", err.Error()),
		})
		return
	}

	err = app.validateBucketName(bucketName.BucketName)
	if err != nil {
		app.ErrorLogger.Printf("error validating bucket name: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error validating bucket name: %v", err.Error()),
		})
		return
	}

	err = app.DB.VerifyBucketName(bucketName.BucketName)
	if err != nil {
		app.ErrorLogger.Printf("error verifying bucket name: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error verifying bucket name: %v", err.Error()),
		})
		return
	}

	_, err = app.S3.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName.BucketName),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String("eu-central-1"),
		},
	})

	if err != nil {
		app.ErrorLogger.Printf("error creating bucket: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error creating bucket: %v", err.Error()),
		})
		return
	}

	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("Accept", "application/json")

	err = app.DB.StoreBucket(bucketName.BucketName)
	if err != nil {
		app.ErrorLogger.Printf("error storing bucket name: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error storing bucket name: %v", err.Error()),
		})
		return
	}

	app.InfoLogger.Printf("successfully created bucket: %s\n", bucketName.BucketName)
	c.JSON(200, gin.H{
		"error":   false,
		"message": fmt.Sprintf("successfully created bucket: %s", bucketName.BucketName),
	})
}

func (app *Application) validateBucketName(bucketName string) error {
	// Bucket name length constraints
	if len(bucketName) < 3 || len(bucketName) > 63 {
		return fmt.Errorf("bucket name must be between 3 and 63 characters long")
	}

	// Bucket name format constraints
	match, _ := regexp.MatchString("^[a-zA-Z0-9.-]+$", bucketName)
	if !match {
		return fmt.Errorf("bucket name can only contain alphanumeric characters, periods, and hyphens")
	}

	// Bucket name label format constraints
	labels := regexp.MustCompile("[^.]+").FindAllString(bucketName, -1)
	for _, label := range labels {
		if label[0] == '.' || label[len(label)-1] == '.' {
			return fmt.Errorf("bucket name labels cannot start or end with a period")
		}
		if label == "-" || label == ".." {
			return fmt.Errorf("bucket name labels cannot be just a hyphen or double periods")
		}
	}

	return nil
}

// UploadFiles function uploads a file to a bucket in AWS S3
// bucketName is the name of the bucket to which the file will be uploaded
// fileURI is the path to the file to be uploaded
// returns an error if the file could not be uploaded
func (app *Application) UploadFiles(c *gin.Context) {
	bucketName := c.Param("bucket-name")

	var bucketUploadRequest struct {
		BucketName string `json:"bucketName"`
		File       []byte `json:"file"`
		Filename   string `json:"filename"`
	}
	err := c.ShouldBindJSON(&bucketUploadRequest)
	if err != nil {
		app.ErrorLogger.Printf("error binding file name: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error binding file name: %v", err.Error()),
		})
		return
	}

	// if you want to upload a file by providing the file path
	// Firslty, in struct bucketUploadRequest, remove file field
	// then uncomment the following code, and comment all another unnecessary code
	// file, err := os.Open(bucketUploadRequest.Filename)
	// if err != nil {
	// 	app.ErrorLogger.Printf("error opening file: %v\n", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error":   true,
	// 		"message": fmt.Sprintf("error opening file: %v", err.Error()),
	// 	})
	// 	return
	// }
	//
	// _, err = app.S3.PutObject(&s3.PutObjectInput{
	// 	Bucket: aws.String(bucketName),
	// 	Key:    aws.String(filepath.Base(bucketUploadRequest.Filename)),
	// 	Body:   file,
	// })

	err = app.DB.CheckIfFileExists(bucketName, bucketUploadRequest.Filename)
	if err != nil {
		app.ErrorLogger.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	_, err = app.S3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filepath.Base(bucketUploadRequest.Filename)),
		Body:   bytes.NewReader(bucketUploadRequest.File),
	})

	if err != nil {
		app.ErrorLogger.Printf("error putting object into bucket: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error putting object into bucket %v\n", err),
		})
		return
	}

	// till here

	err = app.DB.StoreFiles(bucketName, bucketUploadRequest.Filename)
	if err != nil {
		app.ErrorLogger.Printf("error storing file name: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error storing file name: %v", err.Error()),
		})
		return
	}

	app.InfoLogger.Printf("successfully uploaded file %s to bucket %s\n", filepath.Base(bucketUploadRequest.Filename), bucketName)
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "successfully uploaded file into bucket!",
	})
}

func (app *Application) ListFiles(c *gin.Context) {
	var payload struct {
		BucketName string `json:"bucketName"`
	}
	payload.BucketName = c.Param("bucket-name")

	out, err := app.S3.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(payload.BucketName),
	})

	if err != nil {
		app.ErrorLogger.Printf("error listing objects in bucket: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error listing objects in bucket: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully listed objects in bucket!",
		"objects": out,
	})
}

func (app *Application) DownloadFile(c *gin.Context) {
	var payload struct {
		BucketName string `json:"bucketName"`
		ObjectKey  string `json:"objectKey"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		app.ErrorLogger.Printf("error binding bucket name and object key: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error binding bucket name and object key: %v", err.Error()),
		})
		return
	}

	params := &s3.GetObjectInput{
		Bucket: aws.String(payload.BucketName),
		Key:    aws.String(payload.ObjectKey),
	}

	resp, err := app.S3.GetObject(params)
	if err != nil {
		app.ErrorLogger.Printf("error getting object from bucket: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error getting object from bucket: %v", err.Error()),
		})
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		app.ErrorLogger.Printf("error getting current working directory: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error getting current working directory: %v", err.Error()),
		})
		return
	}

	filepath := filepath.Join(wd, payload.ObjectKey)

	defer resp.Body.Close()
	file, err := os.Create(filepath)
	if err != nil {
		app.ErrorLogger.Printf("error creating file: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error creating file: %v", err.Error()),
		})
		return
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		app.ErrorLogger.Printf("error copying file: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error copying file: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "successfully downloaded file!",
	})
}

func (app *Application) CheckBucketName(c *gin.Context) {
	var payload struct {
		BucketName string `json:"bucketName"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		app.ErrorLogger.Printf("error binding bucket name: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error binding bucket name: %v", err.Error()),
		})
		return
	}

	err = app.validateBucketName(payload.BucketName)
	if err != nil {
		app.ErrorLogger.Printf("error validating bucket name: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": fmt.Sprintf("error validating bucket name: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "bucket name is valid!",
	})
}
