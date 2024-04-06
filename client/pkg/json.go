package pkg

import (
	"encoding/json"
	"io"
	"net/http"
)

type Json struct {
}

func (j *Json) Encode(w http.ResponseWriter, responseBody interface{}) []byte {
	encodedBody, err := json.Marshal(responseBody)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}

	return encodedBody
}

func (j *Json) Decode(w http.ResponseWriter, body io.ReadCloser, schema interface{}) interface{} {

	err := json.NewDecoder(body).Decode(schema)

	if err != nil {
		http.Error(w, "Error occurred while decoding json", http.StatusBadRequest)
		return nil
	}

	return schema
}
