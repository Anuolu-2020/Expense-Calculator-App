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
		log.Fatal(err.Error())
		return
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Printf("Error reading template file: %v", err)
		return
	}

	err = tmpl.Execute(w, body)
	if err != nil {
		log.Printf("Error Executing Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SendExecutedTemplate(w http.ResponseWriter, path string, body interface{}) {
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Printf("Error reading template file: %v", err)
		return
	}

	err = tmpl.ExecuteTemplate(w, path, body)
	if err != nil {
		log.Printf("Error Executing Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
