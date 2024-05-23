package auth

import (
	"crypto/sha512"
	"estate/database"
	"estate/handlers"
	"estate/jager"
	"fmt"
	"log"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	jwtToken, err := r.Cookie("jwt")
	if err != nil {
		handlers.SendErrorPage(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	} else {
		log.Println("LOGOUT// TOKEN:", jwtToken)
		hashToken := sha512.Sum512([]byte(jwtToken.Value))
		hashTokenString := fmt.Sprintf("%x", hashToken)
		check, err := database.DeleteToken(hashTokenString)
		if check {
			fmt.Println("logout runnning....")
			//http.Redirect(w, r, "/login", 301)
		} else {
			w.Write([]byte("Token can't delete!"))
			log.Println("Token can't delete err:", err)
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	jsonData, _ := jager.Read(w, r)
	status, check, username := database.CheckUser(jsonData)
	if status == 404 {
		w.WriteHeader(404)
	}

	token := CreateRefreshToken(username, r)
	hashToken := sha512.Sum512([]byte(token))
	hashTokenString := fmt.Sprintf("%x", hashToken)
	database.AddToken(hashTokenString)

	if check {
		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    token,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/admin", 304)
	}
}
