package server

import (
	// "os"
	// "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func Listen() {
	r := mux.NewRouter()
    r.HandleFunc("/", IndexHandler).Methods("GET")
    // r.HandleFunc("/products", ProductsHandler)
    // r.HandleFunc("/articles", ArticlesHandler)
    log.Fatal(http.ListenAndServe(":8080", r))
}
