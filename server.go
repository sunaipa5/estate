package main

import (
	"estate/jager"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func runServer() {
	app := chi.NewRouter()
	app.Use(middleware.Logger)

	servdir := getFrontendDist()
	fs := http.FileServer(http.Dir(servdir))
	app.Handle("/assets/*", fs)

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, servdir+"index.html")
	})

	app.Get("/ilanlar", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, servdir+"index.html")
	})

	app.Get("/getnotice", func(w http.ResponseWriter, r *http.Request) {
		jager.JSON(w, getNotices())
	})

	app.Post("/addnotice", func(w http.ResponseWriter, r *http.Request) {
		
		jsonData, err := jager.Read(w, r)
		if err != nil {
			jager.StringJSON(w,`{"message":"Error"}`)
		}
		jager.StringJSON(w,`{"message":"Success"}`)

		addnotice(jsonData)
	})

	http.ListenAndServe(getListenAdress(), app)
}
