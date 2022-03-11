package main

import (
	"embed"
	"html/template"
)

//go:embed view/*.gohtml
var templates2 embed.FS

var MyTemplates = template.Must(template.ParseFS(templates2, "view/*.gohtml"))
