package template

import (
	"fmt"
	"html/template"
	"net/http"
)

type admin struct {
	Title, Heading string
	Templates []string
}

func RenderAdmin(w http.ResponseWriter, r *http.Request) {
	a := admin{Title: "Cloudshout Official Blog - Admin", Heading: "Admin Page"}
	a.Templates = []string{"Text Post", "Half Width Text Post"}
	
	t, err := template.ParseFiles("web/admin.html")

	// check err

	t.Execute(w, a)
} // end