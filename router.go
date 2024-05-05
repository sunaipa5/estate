package main

import (
	"crypto/sha512"
	"encoding/json"
	"estate/jager"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
	"strconv"
	"html/template"
)



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
		MaxAge:           300,
	}))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		sendErrorPage(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	})

	servdir := getFrontendDist()
	fs := http.FileServer(http.Dir(servdir))
	r.Handle("/assets/*", fs)
	r.Handle("/img/*", http.FileServer(http.Dir("./")))

	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"index.html")
		})
		r.Get("/ilanlar", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"index.html")
		})
	})

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

			token := refresh(username, r)
			hashToken := sha512.Sum512([]byte(token))
			hashTokenString := fmt.Sprintf("%x", hashToken)
			addToken(hashTokenString)

			if check {
				cookie := &http.Cookie{
					Name:     "jwt",
					Value:    token,
					HttpOnly: true,
				}
				http.SetCookie(w, cookie)
				http.Redirect(w, r, "/admin/ilanlar", 304)
			}

		})
		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			jwtToken, err := r.Cookie("jwt")
			if err != nil {
				sendErrorPage(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			} else {
				log.Println("LOGOUT// TOKEN:", jwtToken)
				hashToken := sha512.Sum512([]byte(jwtToken.Value))
				hashTokenString := fmt.Sprintf("%x", hashToken)
				check, err := logout(hashTokenString)
				if check {
					http.Redirect(w, r, "/login", 301)
				} else {
					w.Write([]byte("Token can't delete!"))
					log.Println("Token can't delete err:", err)
				}
			}
		})
	    r.Get("/ilan/{page}",func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"index.html")
		})

		r.Get("/api/getnotices/{page}", func(w http.ResponseWriter, r *http.Request) {
			pageNumber, err := strconv.Atoi(chi.URLParam(r, "page"))
			if err != nil {
				jager.StringJSON(w, `{"error":"Page number type must be an integer !"}`)
				return
			}
			jsonData := getNotices(pageNumber)
			jager.JSON(w, jsonData)
		})

		r.Post("/getnotice", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("get notice running.....")
			jsonData, err := jager.Read(w, r)
			if err != nil {
				jager.StringJSON(w, `{"message":"Error"}`)
				return
			}

			type Page struct {
				PageName string `json:"pageName"`
			}

			var data Page
			err = json.Unmarshal([]byte(jsonData), &data)
			if err != nil {
				jager.MapJSON(w, map[string]interface{}{
					"Error": err,
				})
				return
			}
			jager.Write(w, getNoticePage(data.PageName))
		})
	})

	//Refresh token required routes
	r.Group(func(r chi.Router) {
		r.Use(refreshTokenMiddleware(refresh_token))
		//r.Use(jwtauth.Authenticator(refresh_token))

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"admin.html")
		})

		r.Get("/admin/ilanlar", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"admin.html")
		})

		r.Get("/admin/getaccesstoken", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			if claims["type"] == "refresh" {
				token := access(claims["username"].(string), r)
				jager.StringJSON(w, `{"token":"`+token+`"}`)

			}

		})

		r.Get("/admin/ilan/{title}", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"admin.html")
		})


		r.Get("/admin/ilanlar/{page}", func(w http.ResponseWriter, r *http.Request) {
			pageNumber, err := strconv.Atoi(chi.URLParam(r, "page"))
			if err != nil {
				jager.StringJSON(w, `{"error":"Page number type must be an integer !"}`)
				return
			}
			jsonData := getNotices(pageNumber)
			jager.JSON(w, jsonData)
		})

		r.Get("/admin/ilan-ekle", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"admin.html")
		})

		r.Get("/admin/send", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, servdir+"admin.html")
		})

	})

	//API - Access token required routes
	r.Group(func(r chi.Router) {
		r.Use(refreshTokenMiddleware(refresh_token))
		r.Use(accessTokenMiddleware(access_token))

		r.Post("/addnotice", func(w http.ResponseWriter, r *http.Request) {
			saveImages(r)
		})

	})

	http.ListenAndServe(getListenAdress(), r)
}

func sendErrorPage(w http.ResponseWriter, code int, name string) {
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

