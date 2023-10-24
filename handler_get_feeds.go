package main

import (
	"fmt"
	"net/http"
)

func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}
	resp := AppBaseResponse{}
	resp.Code = 200
	resp.Message = "Success"
	resp.Data = dataBaseFeedsToFeeds(feeds)
	respondWithJson(w, 200, resp)
}
