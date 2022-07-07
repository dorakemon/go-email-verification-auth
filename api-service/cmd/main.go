package main

import (
	"api-service/cmd/server/router"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &Template{
		templates: template.Must(template.ParseGlob("./server/templates/*.go.html")),
	}
	e.Renderer = renderer

	e = router.InitRouter(e)

	e.Logger.Fatal(e.Start(":9999"))
}
