package service

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("Incoming update cannot be decoded! %s", err.Error())
		return nil, err
	}
	return &update, nil
}
