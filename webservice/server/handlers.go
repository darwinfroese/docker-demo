package server

import (
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
	health := healthCheckFull()

	err := json.NewEncoder(w).Encode(&health)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, "We've fully fallen apart")
		return
	}

	return
}

func apiShopHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API Shop Handler")
}

func apiLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API Login Handler")

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
