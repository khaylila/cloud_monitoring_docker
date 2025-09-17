package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Definisikan struktur untuk respons JSON
type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Siapkan data respons
	response := Response{Message: "Hello World"}

	// Set header content-type ke application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode struct ke JSON dan kirim sebagai respons
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Daftarkan handler untuk endpoint "/"
	http.HandleFunc("/", helloHandler)

	// Mulai server di port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
