package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func InitializeRouter() *http.Server {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1/").Subrouter()
	webRouter := router.PathPrefix("/").Subrouter()

	apiRouter = registerAPIRoutes(apiRouter)
	webRouter = registerWebRoutes(webRouter)

	address := "0.0.0.0:80"

	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return srv
}

func registerAPIRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/health", apiHealthHandler)

	return r
}

func registerWebRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/health", webHealthHandler)

	return r
}
