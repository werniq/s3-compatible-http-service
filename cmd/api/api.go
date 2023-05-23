package main

import (
	"cubbit_interview/internal/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Application struct {
	DB          *models.DatabaseModel
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	AwsSession  *session.Session
	S3          *s3.S3
	Endpoint    string
}

var (
	endpointUri           = "http://localhost:8080"
	aws_access_secret_key = "HUWoyBo2DHS3B86wC2brVvH3kET7gFVtt6v9cMqF"
	aws_access_key        = "AKIA4JRYNABS6LDXDOKD"
)

// NewApplication creates a new Application instance, which is used to
// configure the routes and handlers, and most important - run application
func NewApplication(s *session.Session, sthr *s3.S3) *Application {
	return &Application{
		ErrorLogger: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLogger:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		AwsSession:  s,
		S3:          sthr,
		Endpoint:    "http://localhost:4001",
	}
}

func main() {
	// in terms of TODO
	// this project will look like this:
	// first, set up development environment. this includes getting appropriate aws keys, and getting go dependencies
	// secondly, set up the server. this includes setting up the routes, and setting up the handlers
	// thirdly, set up the client. this includes setting up the routes, and setting up the handlers
	// looks pretty straightforward.

	// i decided not to use database, because it is not necessary to.
	// it only needed to store the bucket names, and the files in the buckets.
	// however if you try to create bucket or upload file, which already exists, it will throw an error.

	// i provided description of the project in the README.md file.

	creds := credentials.NewStaticCredentials(aws_access_key, aws_access_secret_key, "")

	s, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: creds,
		Endpoint:    aws.String("https://s3.eu-central-1.amazonaws.com"),
	})
	if err != nil {
		log.Printf("Error creating session: %v\n", err)
		return
	}

	svc := s3.New(s)

	app := NewApplication(s, svc)

	r := gin.Default()

	app.ConfigureRoutes(r)

	if err := r.Run(":4001"); err != nil {
		app.ErrorLogger.Printf("error running server on port 4001: %v\n", err)
		return
	}

	app.InfoLogger.Printf("SUCCESS:\t Server is running on port 4001\n")
}
