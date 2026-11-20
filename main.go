package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "pong",
			"arch":    runtime.GOARCH,
		})
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
