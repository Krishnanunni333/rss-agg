package main

import (
	"net/http"

	"github.com/Krishnanunni333/rss-agg-server/internal/auth"
	"github.com/Krishnanunni333/rss-agg-server/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		handler(w, r, user)
	}
}
