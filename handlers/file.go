package handlers

import (
	"encoding/json"
	"estate/database"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func AssetHandler(baseDir, folder string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath := strings.TrimPrefix(r.URL.Path, folder)
		filePath = filepath.Join(baseDir, filePath)

		fileInfo, err := os.Stat(filePath)
		if os.IsNotExist(err) || fileInfo.IsDir() {
			SendErrorPage(404, http.StatusText(404))(w, r)
			return
		}

		http.ServeFile(w, r, filePath)
	}
}

type Property struct {
	Title           string `json:"Title"`
	Address         string `json:"Address"`
	Room_number     string `json:"Room_number"`
	Centare         int    `json:"Centare"`
	Floor           int    `json:"Floor"`
	Floor_number    int    `json:"Floor_number"`
	Description     string `json:"Description"`
	Page_name       string `json:"Page_name"`
	File_names      string `json:"File_names"`
	Property_type   string `json:"Property_type"`
	Property_status string `json:"Property_status"`
}

func SaveImages(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat("img")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("img", 0755)
		if errDir != nil {
			log.Println(errDir)
			return
		}
	}

	jsonData := r.FormValue("jsonData")
	var name Property
	err = json.Unmarshal([]byte(jsonData), &name)
	if err != nil {
		log.Println(err)
		return
	}

	r.ParseMultipartForm(10 << 20)
	files := r.MultipartForm.File["files[]"]

	for i, file := range files {
		src, err := file.Open()
		if err != nil {
			log.Println(err)
			return
		}
		defer src.Close()

		newFilename := fmt.Sprintf("%s-%d%s", name.Page_name, i+1, filepath.Ext(file.Filename))
		name.File_names += newFilename + " "
		dstPath := filepath.Join("img", newFilename)
		dst, err := os.Create(dstPath)
		if err != nil {
			log.Println(err)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			log.Println(err)
			return
		}

	}

	jsonOutput, _ := json.Marshal(name)
	database.AddProperty(jsonOutput)
}
