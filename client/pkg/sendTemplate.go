package pkg

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func readPath(path string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	absPath := filepath.Join(cwd, "templates", path)

	return absPath, nil
}

func SendTemplate(w http.ResponseWriter, path string, body interface{}) {
	path, err := readPath(path)
	log.Println(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Println("Error reading template file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
