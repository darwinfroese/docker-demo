package server

import (
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
	log.Println("auth health check hit")
	fmt.Fprintln(w, "Auth healthy")

	return
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
