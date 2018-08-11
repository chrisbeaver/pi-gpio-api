package server

import (
	// "os"
	// "encoding/json"
	// "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func Listen(port string) {
	r := mux.NewRouter()
    r.HandleFunc("/", IndexHandler).Methods("GET")
    // r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	// cmd := fmt.Sprintf(":%d", port)
	// fmt.Println(cmd)
    log.Fatal(http.ListenAndServe(":" + port, r))
}
