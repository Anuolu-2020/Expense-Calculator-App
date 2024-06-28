package handlers

import (
	"io"
	"net/http"
)

func getImage(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h Handler) GetProfilePic(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	if url == "" {
		resp, err := getImage("https://ui-avatars.com/api/?background=random")
		if err != nil {
			http.Error(w, "Failed to copy image data", http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, "Failed to copy image data", http.StatusInternalServerError)
			return
		}
	}

	resp, err := getImage(url)
	if err != nil {
		http.Error(w, "Failed to fetch image", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to copy image data", http.StatusInternalServerError)
		return
	}
}
