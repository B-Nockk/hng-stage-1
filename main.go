package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	jsonMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "API is running"})
	})

	mux.HandleFunc("/ci-bot", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "CI - CD worked fine"})
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "healthy"})
	})

	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"name":   "Benson OgheneOchuko",
			"email":  "ob.ogheneochuko@gmail.com",
			"github": "https://github.com/B-Nockk",
		})
	})

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", jsonMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
