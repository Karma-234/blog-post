package main

import "net/http"

type readyResponse struct {
	Code    int
	Message string
	Data    map[string][]string
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	var resp readyResponse
	resp.Code = 200
	resp.Message = "Successful request"
	resp.Data = map[string][]string{"reports": {}, "generalSummary": {}}
	respondWithJson(w, 200, resp)
}
