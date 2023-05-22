package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type templateData struct {
	StringMap            map[string]string
	IntMap               map[string]int
	FloatMap             map[string]float32
	Data                 map[string]interface{}
	IsAuthenticated      int
	MainnetAddress       string
	ErrorData            []string
	StripeSecretKey      string
	StripePublishableKey string
	CSRFToken            string
	Flash                string
	Warning              string
	Error                string
	API                  string
	CSSVersion           string
}

//go:embed templates
var templateFS embed.FS

func (app *Application) addDefaultData(td *templateData, r *http.Request) *templateData {
	td.API = app.cfg.api
	return td
}

func (app *Application) renderTemplate(c *gin.Context, page string, td *templateData) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	_, templateInMap := app.templateCache[templateToRender]

	if app.cfg.env == "production" && templateInMap {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}

	td = app.addDefaultData(td, c.Request)

	err = t.Execute(c.Writer, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *Application) parseTemplate(page, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).ParseFS(templateFS, "templates/base.layout.gohtml", templateToRender)

	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.templateCache[templateToRender] = t
	return t, nil
}
