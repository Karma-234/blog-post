package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(r http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marsha payload: %v", payload)
		r.WriteHeader(500)
	}
	r.Header().Add("Content-Type", "application/json")
	r.WriteHeader(code)
	r.Write(data)

}
func respondWithError(r http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Fatal("Responsing with 500, Server error")
	}
	type ErrorResponse struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
		Data  interface{}
	}
	respondWithJson(r, code, ErrorResponse{
		Error: message,
		Code:  code,
		Data:  struct{}{},
	})
}
