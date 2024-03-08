package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/itsaril/tsis1/internal/handlers"
)

func main() {
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/dramas", handlers.GetDramas).Methods("GET")
    r.HandleFunc("/dramas/{name}", handlers.GetDrama).Methods("GET")
    r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

    // Start server
    log.Fatal(http.ListenAndServe(":8080", r))
}
