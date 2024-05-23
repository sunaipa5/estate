package options

import (
	"github.com/go-chi/cors"
)

var (
	//server
	Port         = ":3000"
	FrontendDist = "./frontend/dist/"
	Debug        = true

	//Database
	DBhost = "127.0.0.1:3306"
	DBname = "estate"
	DBuser = "dbadmin"
	DBpass = "dbadmin"
)

var CORS = cors.Handler(cors.Options{
	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	AllowedOrigins: []string{"https://*", "http://*"},
	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: false,
	MaxAge:           300,
})