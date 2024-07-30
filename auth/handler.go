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
		handlers.SendErrorPage(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))(w, r)
	} else {
		hashToken := sha512.Sum512([]byte(jwtToken.Value))
		hashTokenString := fmt.Sprintf("%x", hashToken)
		check, err := database.DeleteToken(hashTokenString)
		if check {
			fmt.Println("logout runnning....")
			RedirectLogin(w,r)

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

func RedirectLogin(w http.ResponseWriter, r *http.Request) {
	htmlContent := "<script>window.location.replace('/login');</script>"

	w.WriteHeader(401)
	w.Write([]byte(htmlContent))

}
