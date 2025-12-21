package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "hello backend"})
	})
	http.ListenAndServe(":8000", nil)
}
