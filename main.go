package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No Port found")
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
