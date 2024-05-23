package auth

import (
	"crypto/sha512"
	"estate/database"
	"estate/jager"
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
	"time"
)

var Access_token *jwtauth.JWTAuth
var Refresh_token *jwtauth.JWTAuth

func init() {
	Access_token = jwtauth.New("HS256", []byte("secret"), nil)
	Refresh_token = jwtauth.New("HS256", []byte("secret2"), nil)
}

func CreateAccessToken(username string, r *http.Request) string {
	now := time.Now()
	expireTime := now.Add(15 * time.Minute)
	expireUnix := expireTime.Unix()

	_, access_string, _ := Access_token.Encode(map[string]interface{}{
		"username":      username,
		"type":          "access",
		"authorization": "admin",
		"exp":           expireUnix,
		"userAgent":     r.UserAgent(),
	})

	log.Println("Access Token Created.")
	return access_string

}

func CreateRefreshToken(username string, r *http.Request) string {
	now := time.Now()
	expireTime := now.Add(3 * 30 * 24 * time.Hour)
	expireUnix := expireTime.Unix()

	_, refresh_string, _ := Refresh_token.Encode(map[string]interface{}{
		"username":      username,
		"type":          "refresh",
		"authorization": "admin",
		"exp":           expireUnix,
		"userAgent":     r.UserAgent(),
	})

	log.Println("Refresh Token Created.")
	return refresh_string
}

func RefreshTokenMiddleware(ja *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			log.Println("Refresh Authmiddleware running...")
			ctx := r.Context()
			jwtToken, err := r.Cookie("jwt")
			if err != nil {
				fmt.Println("TOKEN COOKIE NOT FOUND")
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			token, err := jwtauth.VerifyToken(ja, jwtToken.Value)
			log.Println("err:", err, "token:", token)
			if err != nil {
				log.Println("err:", err)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			} else {
				log.Println("here run")
				claims := token.PrivateClaims()
				if claims["userAgent"] != r.UserAgent() {
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
				hashToken := sha512.Sum512([]byte(jwtToken.Value))
				hashTokenString := fmt.Sprintf("%x", hashToken)
				check, err := database.CheckToken(hashTokenString)
				if !check {
					fmt.Println("Token not found:", err)
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
			}
			ctx = jwtauth.NewContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(hfn)
	}
}

func AccessTokenMiddleware(ja *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			log.Println("Access Authmiddleware running...")
			ctx := r.Context()

			token, err := jwtauth.VerifyRequest(ja, r, jwtauth.TokenFromHeader)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			ctx = jwtauth.NewContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(hfn)
	}
}

func SetAccessToken(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	if claims["type"] == "refresh" {
		token := CreateAccessToken(claims["username"].(string), r)
		jager.StringJSON(w, `{"token":"`+token+`"}`)

	}

}
