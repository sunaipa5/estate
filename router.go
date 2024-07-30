package main

import (
	"estate/auth"
	"estate/database"
	"estate/handlers"
	"estate/jager"
	"estate/options"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func router() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(options.CORS)

	r.NotFound(handlers.SendErrorPage(http.StatusNotFound, http.StatusText(http.StatusNotFound)))

	r.Get("/assets/*", handlers.AssetHandler(options.FrontendDist,"assets"))
	r.Get("/img/*", handlers.AssetHandler("./","img"))

	//Not token required routes
	r.Group(func(r chi.Router) {
		r.Get("/logout", auth.Logout)
		r.Get("/login", handlers.PageLogin)
		r.Post("/login", auth.Login)

		r.Get("/", handlers.PagePublic)
		r.Get("/ilan/{page}", handlers.PagePublic)
		r.Get("/ilanlar", handlers.PagePublic)
		r.Get("/ilanlar/{status}", handlers.PagePublic)
		r.Get("/ilanlar/{type}/{status}", handlers.PagePublic)
	})

	//Admin
	r.Route("/admin", func(r chi.Router) {
		r.Use(auth.RefreshTokenMiddleware(auth.Refresh_token))

		r.Get("/", handlers.PageAdmin)

		r.Get("/getaccesstoken", auth.SetAccessToken)
		r.Get("/kullanicilar", handlers.PageAdmin)

		r.Get("/ilanlar", handlers.PageAdmin)
		r.Get("/ilanlar/{page}", func(w http.ResponseWriter, r *http.Request) {
			pageNumber, err := strconv.Atoi(chi.URLParam(r, "page"))
			if err != nil {
				jager.StringJSON(w, `{"error":"Page number type must be an integer !"}`)
				return
			}
			jsonData := database.GetProperties(pageNumber, "", "")
			jager.JSON(w, jsonData)
		})
		r.Get("/ilan/{title}", handlers.PageAdmin)
		r.Get("/ilan-ekle", handlers.PageAdmin)
	})

	//API
	r.Route("/api", func(r chi.Router) {
		//Not Token Required Routes
		r.Get("/getProperties/{page}/{status}/{type}", func(w http.ResponseWriter, r *http.Request) {
			pageNumber, err := strconv.Atoi(chi.URLParam(r, "page"))
			if err != nil {
				jager.StringJSON(w, `{"error":"Page number type must be an integer !"}`)
				return
			}
			jsonData := database.GetProperties(pageNumber, chi.URLParam(r, "status"), chi.URLParam(r, "type"))
			jager.JSON(w, jsonData)
		})

		r.Post("/getProperty", handlers.GetProperty)

		//Access and Refresh Token Required Routes
		r.Group(func(r chi.Router) {
			r.Use(auth.RefreshTokenMiddleware(auth.Refresh_token))
			r.Use(auth.AccessTokenMiddleware(auth.Access_token))

			r.Get("/getUsers", func(w http.ResponseWriter, r *http.Request) {
				jager.JSON(w, database.GetUsers())
			})
			r.Post("/addUser", handlers.AddUser)
			r.Post("/deleteUser", handlers.DeleteUser)

			r.Post("/addProperty", handlers.SaveImages)
			r.Post("/updateProperty", handlers.UpdateProperty)
			r.Post("/deleteProperty", handlers.DeleteProperty)

		})
	})

	http.ListenAndServe(options.Port, r)
}
