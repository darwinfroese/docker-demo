package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func InitializeRouter() *http.Server {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1/").Subrouter()
	router.PathPrefix("/static/").Handler(http.FileServer(http.Dir("www")))
	router.PathPrefix("/").HandlerFunc(indexHandler)

	apiRouter = registerAPIRoutes(apiRouter)

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
	r.HandleFunc("/healthfull", apiHealthHandler)
	r.HandleFunc("/newitem", apiAddItemHandler)
	r.HandleFunc("/shop", apiShopHandler)
	r.HandleFunc("/login", apiLoginHandler)

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/index.html")
}
