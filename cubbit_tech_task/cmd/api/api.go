package main

import (
	"cubbit_interview/internal/driver"
	"cubbit_interview/internal/models"
	"database/sql"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
func NewApplication(db *sql.DB, s *session.Session, sthr *s3.S3) *Application {
	return &Application{
		DB:          &models.DatabaseModel{DB: db},
		ErrorLogger: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLogger:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		AwsSession:  s,
		S3:          sthr,
		Endpoint:    "http://localhost:8080",
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		return
	}

	creds := credentials.NewStaticCredentials(aws_access_key, aws_access_secret_key, "")

	s, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: creds,
		Endpoint:    aws.String("https://s3.eu-central-1.amazonaws.com"),
	})

	svc := s3.New(s)

	db, err := driver.OpenDb()
	if err != nil {
		log.Printf("Error opening database: %v\n", err)
		return
	}

	app := NewApplication(db, s, svc)

	r := gin.Default()

	app.ConfigureRoutes(r)

	if err := r.Run(":4001"); err != nil {
		app.ErrorLogger.Printf("error running server on port 4001: %v\n", err)
		return
	}

	app.InfoLogger.Printf("SUCCESS:\t Server is running on port 4001\n")
}
