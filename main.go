package main

import (
	"apigateway"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", apigateway.HandleRequest)
	http.ListenAndServe(":8080", router)
}
