package pkg

import (
	"encoding/json"
	"io"
)

type Json struct {
}

func (j *Json) Encode(responseBody interface{}) ([]byte, error) {
	encodedBody, err := json.Marshal(responseBody)

	if err != nil {
		return nil, err
	}

	return encodedBody, nil
}

func (j *Json) Decode(body io.ReadCloser, schema interface{}) (interface{}, error) {

	err := json.NewDecoder(body).Decode(schema)

	if err != nil {
		return nil, err
	}

	return schema, nil
}
