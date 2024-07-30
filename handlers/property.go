package handlers

import(
	"estate/jager"
	"estate/database"
	"net/http"
	"encoding/json"
)

func UpdateProperty(w http.ResponseWriter, r *http.Request) {
	jsonData, err := jager.Read(w, r)
	if err != nil {
		w.WriteHeader(400)
	} else {
		err = database.UpdateProperty(jsonData)
		if err != nil {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(200)
		}
	}
}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {
	jsonData, err := jager.Read(w, r)
	if err != nil {
		w.WriteHeader(400)
	} else {
		err = database.DeleteProperty(jsonData)
		if err != nil {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(200)
		}
	}
}


func GetProperty(w http.ResponseWriter, r *http.Request) {
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
	jager.Write(w, database.GetPropertyPage(data.PageName))
}