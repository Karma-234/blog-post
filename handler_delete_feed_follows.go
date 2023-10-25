package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/karma-234/blog-post/internal/database"
)

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feedID"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}
	deleteError := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		FeedID: params.FeedID,
		UserID: user.ID,
	})
	if deleteError != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed: %v", err))
		return
	}
	resp := AppBaseResponse{}
	resp.Code = 200
	resp.Message = "Success"
	respondWithJson(w, 200, resp)
}
