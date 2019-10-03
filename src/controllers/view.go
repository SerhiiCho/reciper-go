package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

func view(w http.ResponseWriter, viewName string, data map[string]interface{}) {
	tpl := template.Must(template.ParseGlob("templates/*.gohtml"))
	err := tpl.ExecuteTemplate(w, viewName+".gohtml", data)

	if err != nil {
		// TODO: return error view
		fmt.Print(err)
	}
}
