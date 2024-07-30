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
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	http.ServeFile(w, r, options.FrontendDist+"login.html")
}