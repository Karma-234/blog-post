package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/karma-234/blog-post/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserId    uuid.UUID `json:"userId"`
}
type FeedFollow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

func dataBaseUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type AppBaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func dataBaseFeed(dbUser database.Feed) Feed {
	return Feed{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		URL:       dbUser.Url,
		UserId:    dbUser.UserID,
	}
}
func dataBaseFeedFollow(dbUser database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		FeedID:    dbUser.FeedID,
		UserID:    dbUser.UserID,
	}
}

func dataBaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, v := range dbFeeds {
		feeds = append(feeds, dataBaseFeed(v))
	}
	return feeds
}
