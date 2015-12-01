package views

import (
	"html/template"
	"net/http"

	"goldecto/util"
)

func HelloPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	tmpl, err := template.ParseFiles(util.GetResourcePath("templates", "helloworld.html"))
	if err != nil { panic(err) }

	err = tmpl.Execute(w, nil)
	if err != nil { panic(err) }
}
