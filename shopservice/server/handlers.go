package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/darwinfroese/docker-demo/shopservice/repository"
)

func apiHealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("shop api - health hit")
	fmt.Fprintln(w, "Shop Healthy")

	return
}

func apiGetItemsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("shop api - get items hit")

	items, err := repository.GetAllItems()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Item retrieval is broken"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&items)
}

func apiCreateItemHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("shop api - create item hit")

	var i repository.Item
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad item sent"))
		return
	}

	err = repository.CreateItem(i)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't create the item"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
