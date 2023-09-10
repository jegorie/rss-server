package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT not defined in env")
	}

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Logger)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	r.Mount("/v1", v1Router)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	fmt.Printf("Server ready to use on http://localhost:%v\n", port)
	http.ListenAndServe(":"+port, r)

}
