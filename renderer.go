package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Renderer struct {
	templates *template.Template
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
