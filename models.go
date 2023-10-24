package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/karma-234/blog-post/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
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
