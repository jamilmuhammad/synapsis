package main

import (
	"api-gateway/internal/server"
	"log"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug: true,
	})

	r.Use(corsMiddleware.Handler)

	s := server.NewServer(r)

	log.Fatal(s.Run())
}
