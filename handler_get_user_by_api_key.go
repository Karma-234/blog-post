package main

import (
	"net/http"

	"github.com/karma-234/blog-post/internal/database"
)

func (apiCfg *apiConfig) getUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {

	resp := AppBaseResponse{}
	resp.Code = 200
	resp.Message = "User fetched successfully"
	resp.Data = dataBaseUser(user)
	respondWithJson(w, 200, resp)
}
