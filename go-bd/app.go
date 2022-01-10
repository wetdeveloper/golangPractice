package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "birthdateForm.html", map[string]interface{}{
			"name": "Ehsan!",
		})
	}).Name = "studentName"

	e.POST("/birthdate", func(c echo.Context) error {
		bd := c.FormValue("bd")
		return c.String(http.StatusOK, bd)
	}).Name = "birthdate"
	e.Start(":")
}
