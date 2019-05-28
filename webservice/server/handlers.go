package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/satori/go.uuid"
)

type (
	session struct {
		ID      uuid.UUID
		Valid   bool
		Expires time.Time
	}

	loginBody struct {
		Username, Password string
	}

	healthInfo struct {
		WebStatus   string
		LoginStatus string
		ShopStatus  string
	}
)

var (
	sessions    = map[uuid.UUID]session{}
	sessionLock sync.Mutex
)

func apiHealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("API Health Handler Hit")

	health := healthCheckFull()

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(&health)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		fmt.Fprintln(w, "We've fully fallen apart")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func apiShopHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API Shop Handler")
}

func apiLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("API Login Handler hit")

	var body loginBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "That request is no good")
		return
	}

	log.Printf("API body: %+v\n", body)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&body)

	res, err := http.Post("http://login.docker.demo/api/v1/login", "application/json; charset=utf-8", buf)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if res.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "That request is no good")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Login accepted")
}

/* Web Structure
/
	/health	- healthcheck
	/shop		- shop page
	/login 	- login page
*/

func webHealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WEB Health Handler")
}

func webShopHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WEB Shop Handler")
}

func webLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WEB Login Handler")
}
