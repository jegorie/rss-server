package main

import (
	"fmt"
	"net/http"

	"github.com/jegorie/rss-server/internal/auth"
	"github.com/jegorie/rss-server/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		}

		handler(w, r, user)
	}
}
