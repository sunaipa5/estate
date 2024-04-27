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

func Write(w http.ResponseWriter, getJson []byte) error{
	w.Header().Set("Content-Type", "application/json")
	w.Write(getJson)
	return nil
}

func Read(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	return body, nil
}

func StringJSON(jsonString string) []byte {
    var jsonData interface{}

    err := json.Unmarshal([]byte(jsonString), &jsonData)
    if err != nil {
        log.Println("Unmarshal // Failed to convert JSON")
        return nil
    }

    encodedJSON, err := json.Marshal(jsonData)
    if err != nil {
        log.Println("Marshal // Failed to convert JSON")
        return nil
    }

    return encodedJSON
}


