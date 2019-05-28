package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	item struct {
		Name, Description string
		Price             float32
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
}

func apiShopHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("web api - shop handler hit")

	resp, err := http.Get("http://shop.docker.demo/api/v1/items")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("error communicating with shop.docker.demo"))
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("error reading repsonse from shop.docker.demo"))
		return
	}

	log.Println("Response: ", string(body))

	var items []item
	err = json.Unmarshal(body, &items)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("error parsing response from shop.docker.demo"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&items)
}

func apiAddItemHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("web api - add item handler hit")

	var body item
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "That request is no good")
		return
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&body)

	res, err := http.Post("http://shop.docker.demo/api/v1/items", "application/json; charset=utf-8", buf)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if res.StatusCode != http.StatusCreated {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("That request is no good"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Item added"))
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
