package template

import (
	"embed"
	"html/template"
)

//go:embed static/*.gohtml
var templates2 embed.FS

var MyTemplates = template.Must(template.ParseFS(templates2, "static/*.gohtml"))
