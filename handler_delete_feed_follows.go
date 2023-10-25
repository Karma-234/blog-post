package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/karma-234/blog-post/internal/database"
)

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	followID := chi.URLParam(r, "feedFollowId")
	feedUUID, err := uuid.Parse(followID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing path: %v", err))
		return
	}
	deleteError := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedUUID,
		UserID: user.ID,
	})
	if deleteError != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	resp := AppBaseResponse{}
	resp.Code = 200
	resp.Message = "Success"
	respondWithJson(w, 200, resp)
}
