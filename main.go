package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RouteResponse struct {
	Message string `json:"message"`
}

func main() {

	routerMux := http.NewServeMux()

	routerMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(RouteResponse{"Hello World!"})
	})

	log.Fatal(http.ListenAndServe(":8080", routerMux))

}
