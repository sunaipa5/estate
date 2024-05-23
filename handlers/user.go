package handlers

import (
	"estate/database"
	"estate/jager"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	jsonData, err := jager.Read(w, r)
	if err != nil {
		jager.StringJSON(w, `{"Error":"User can't found!"}`)
	} else {
		err = database.AddUser(jsonData)
		if err != nil {
			jager.StringJSON(w, `{"Error":"User can't add!"}`)
		} else {
			jager.StringJSON(w, `{"Success":"User successfully added."}`)
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	jsonData, err := jager.Read(w, r)
	if err != nil {
		jager.StringJSON(w, `{"Error":"User can't found!"}`)
	} else {
		err = database.DeleteUser(jsonData)
		if err != nil {
			jager.StringJSON(w, `{"Error":"User can't delete!"}`)
		} else {
			jager.StringJSON(w, `{"Success":"User successfully deleted."}`)
		}
	}
}
