package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	fileContent := string(fileBytes)
	
	convertedText, err := service.ConvertText(fileContent)
	if err != nil {
		http.Error(w, "Unable to convert text", http.StatusInternalServerError)
		return
	}

	fileName := time.Now().UTC().String() + filepath.Ext(fileHeader.Filename)
	err = os.WriteFile(fileName, []byte(convertedText), 0644)
	if err != nil {
		log.Printf("Unable to write file: %v", err)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, convertedText)
}
