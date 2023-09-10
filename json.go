package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("5xx response:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	responseWithJson(w, code, errResponse{Error: msg})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
