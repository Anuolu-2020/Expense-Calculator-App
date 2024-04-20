package pkg

import (
	"bytes"
	"encoding/gob"
	"strings"
)

type SessionData struct {
	Username string
	Photo    string
}

func EncodeSessionData(
	Username string,
	Photo string,
) (bytes.Buffer, error) {
	gob.Register(&SessionData{})

	data := SessionData{Username, Photo}

	// Initialize a buffer to hold the gob-encoded data.
	var buf bytes.Buffer

	// Gob-encode the user data, storing the encoded output in the buffer.
	err := gob.NewEncoder(&buf).Encode(&data)
	if err != nil {
		return bytes.Buffer{}, err
	}

	return buf, nil
}

func DecodeSessionData(value string) (SessionData, error) {
	var userSession SessionData

	// Create an strings.Reader containing the gob-encoded value.
	reader := strings.NewReader(value)

	// Decode it into the userSession type. Notice that we need to pass a *pointer* to
	// the Decode() target here?
	if err := gob.NewDecoder(reader).Decode(&userSession); err != nil {
		return SessionData{}, err
	}

	return userSession, nil
}
