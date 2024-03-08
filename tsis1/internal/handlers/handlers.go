package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

type Drama struct {
    Name   string `json:"name"`
    Actors string `json:"actors"`
    Year   int    `json:"year"`
}

var dramas = []Drama{
    {Name: "Descendants of the Sun", Actors: "Song Joong-ki, Song Hye-kyo", Year: 2016},
    {Name: "Goblin", Actors: "Gong Yoo, Kim Go-eun", Year: 2016},
    {Name: "Crash Landing on You", Actors: "Hyun Bin, Son Ye-jin", Year: 2019},
}

// GetDramas returns a list of all Korean dramas
func GetDramas(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(dramas)
}

// GetDrama returns details of a specific Korean drama by name
func GetDrama(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range dramas {
        if item.Name == params["name"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Drama not found", http.StatusNotFound)
}

// HealthCheck provides a simple health check response
func HealthCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Korean Drama API is healthy. Developed by Your Name")
}
