package main

import (
	"fmt"
	"net/http"

	"github.com/karma-234/blog-post/internal/database"
)

func (apiCfg *apiConfig) handlerGetAllFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetAllFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}
	resp := AppBaseResponse{}
	resp.Code = 200
	resp.Message = "Success"
	resp.Data = dataBaseFeedFollowsToFollow(feedFollows)
	respondWithJson(w, 200, resp)
}
