package handlers

import(
	"net/http"
	"estate/options"
)

func PagePublic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, options.FrontendDist+"index.html")
}

func PageAdmin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, options.FrontendDist+"admin.html")
}

func PageLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, options.FrontendDist+"login.html")
}