package main

import (
	"estate/jager"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"time"
)

var access_token *jwtauth.JWTAuth
var refresh_token *jwtauth.JWTAuth

func init() {
	access_token = jwtauth.New("HS256", []byte("secret"), nil)
	refresh_token = jwtauth.New("HS256", []byte("secret2"), nil)
}

func access(username string) string {
	now := time.Now()
	expireTime := now.Add(30 * time.Second)
	expireUnix := expireTime.Unix()

	_, access_string, _ := access_token.Encode(map[string]interface{}{
		"username":      username,
		"type":          "access",
		"authorization": "admin",
		"exp":           expireUnix,
	})

	log.Println("Access Token Created.")
	return access_string

}

func refresh(username string) string {
	_, refresh_string, _ := refresh_token.Encode(map[string]interface{}{
		"username":      username,
		"type":          "refresh",
		"authorization": "admin",
	})

	log.Println("Refresh Token Created.")
	return refresh_string
}

func router() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	servdir := getFrontendDist()
	fs := http.FileServer(http.Dir(servdir))
	r.Handle("/assets/*", fs)

	//LOGIN - Not token required routes
	r.Group(func(r chi.Router) {
		r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"login.html")
		})

		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			log.Println("POST REQUEST:", r.RemoteAddr)
			jsonData, _ := jager.Read(w, r)
			status, check, username := checkUser(jsonData)
			if status == 404 {
				w.WriteHeader(404)
			}

			if check {
				cookie := &http.Cookie{
					Name:     "jwt",
					Value:    refresh(username),
					HttpOnly: true,
				}
				http.SetCookie(w, cookie)
				http.Redirect(w, r, "/ilanlar", 304)
			}

		})
	})

	//Refresh token required routes
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware(refresh_token))
		r.Use(jwtauth.Authenticator(refresh_token))

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"index.html")
		})

		r.Get("/ilanlar", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"index.html")
		})

		r.Get("/getaccesstoken", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			if claims["type"] == "refresh" {
				token := access(claims["username"].(string))
				jsonD := jager.StringJSON(`{"token":"` + token + `"}`)
				log.Println("Access toke created:", jsonD)
				jager.Write(w, jsonD)
			}

		})

	})

	//API - Access token required routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(access_token))
		r.Use(jwtauth.Authenticator(access_token))

		r.Get("/getnotice", func(w http.ResponseWriter, r *http.Request) {
			jager.JSON(w, getNotices())
		})

		r.Get("/checkget", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("checkget running"))
		})

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims)))
		})
	})

	http.ListenAndServe(getListenAdress(), r)
}

func authMiddleware(ja *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Authmiddleware running....")
			ctx := r.Context()
			token, err := jwtauth.VerifyRequest(ja, r, jwtauth.TokenFromHeader, jwtauth.TokenFromCookie)
			if err != nil {
				log.Println("err:", err)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return // Buradan çık, diğer kodları çalıştırma
			}
			log.Println("token:", token, "err:", err)
			ctx = jwtauth.NewContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(hfn)
	}
}
