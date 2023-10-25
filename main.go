package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/karma-234/blog-post/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No Port found")
	}
	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("No Database found")
	}
	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Error connecting to db", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}
	if port == "" {
		log.Fatal("No Database url found")
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https//*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1router := chi.NewRouter()
	v1router.Get("/ready", handlerReadiness)
	v1router.Get("/error", handlerErr)
	v1router.Post("/user", apiCfg.handlerCreateUser)
	v1router.Get("/user", apiCfg.middleWareAuth(apiCfg.getUserByAPIKey))
	v1router.Post("/feed", apiCfg.middleWareAuth(apiCfg.handlerCreateFeed))
	v1router.Get("/feed", apiCfg.handlerGetAllFeeds)
	v1router.Post("/feed-follow", apiCfg.middleWareAuth(apiCfg.handlerCreateFeedFollow))
	v1router.Delete("/feed-follow/{feedFollowId}", apiCfg.middleWareAuth(apiCfg.handlerDeleteFeedFollow))
	v1router.Get("/feed-follow", apiCfg.middleWareAuth(apiCfg.handlerGetAllFeedFollows))

	router.Mount("/v1", v1router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("Started server on port %v\n", port)
	startError := srv.ListenAndServe()
	if startError != nil {
		log.Fatal(startError)
	}

}
