package web

import (
	"encoding/json"
	"errors"
	"net/http"
)

func WriteAsJson(w http.ResponseWriter, data interface{}) error {

	payload, err := json.Marshal(data)
	if err != nil {
		return errors.New("unable to marshal: " + err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)

	return nil
}
