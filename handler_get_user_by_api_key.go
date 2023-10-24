package main

import (
	"fmt"
	"net/http"

	"github.com/karma-234/blog-post/internal/auth"
)

func (apiCfg *apiConfig) getUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user %v", err))
		return
	}
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user %v", err))
		return
	}
	resp := AppBaseResponse{}
	resp.Code = 200
	resp.Message = "User fetched successfully"
	resp.Data = dataBaseUser(user)
	respondWithJson(w, 200, resp)
}
