package main

import (
	"cubbit_interview/internal/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"os"
)

type config struct {
	port int
	api  string
	env  string
	db   struct {
		dsn string
	}
}

type Application struct {
	cfg           config
	database      models.DatabaseModel
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	cfg := config{
		port: 4000,
		api:  "http://localhost:8080",
		env:  "development",
		db: struct {
			dsn string
		}{
			dsn: "root:root@tcp(localhost:3306)/cubbit?parseTime=true",
		},
	}

	app := &Application{cfg: cfg}
	app.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.infoLog.Println("Starting server on port 8080")
	r := gin.Default()

	tp := make(map[string]*template.Template)
	app.templateCache = tp

	app.ConfigureClientRoutes(r)

	app.infoLog.Printf("Starting DEVELOPMENT server on port %d", app.cfg.port)
	if err := r.Run(":4000"); err != nil {
		app.errorLog.Println("Error starting server on port 8080")
	}

}
