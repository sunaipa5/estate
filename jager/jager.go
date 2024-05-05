package jager

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

)

/*
   ___
  |_  |
    | | __ _  __ _  ___ _ __
    | |/ _` |/ _` |/ _ \ '__|
/\__/ / (_| | (_| |  __/ |
\____/ \__,_|\__, |\___|_|
              __/ |
             |___/
  Quick Tools For net/http
	github.com/sunaipa5
*/

func JSON(w http.ResponseWriter, getJson interface{}) error{
	jsonData, err := json.Marshal(getJson)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return nil
}

func Write(w http.ResponseWriter, getJson []byte){
	w.Header().Set("Content-Type", "application/json")
	w.Write(getJson)
}

func Read(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	return body, nil
}

func StringJSON(w http.ResponseWriter, input string){
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(input), &jsonMap)
	if err != nil {
       log.Println("JAGER ERROR:",err)
	}

	jsonData, err := json.Marshal(jsonMap)
	if err != nil {
		log.Println("JAGER ERROR:",err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func IsJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}


func MapJSON(w http.ResponseWriter, data map[string]interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
