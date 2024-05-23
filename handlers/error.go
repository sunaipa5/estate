package handlers

import (
	"net/http"
	"html/template"
)

func SendErrorPage(w http.ResponseWriter, code int, name string) {
	w.WriteHeader(code)
	tmpl, err := template.ParseFiles("error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Code": code,
		"Name": name,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
