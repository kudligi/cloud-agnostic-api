package main

import (
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode("pong")
    })


    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode("Hello world from go api")
    })

    log.Println("API is running!")
    http.ListenAndServe(":4000", router)
}
