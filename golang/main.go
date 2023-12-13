package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Menentukan port yang akan digunakan
	port := 9091

	// Menetapkan endpoint "/healthcheck" dengan handler yang sesuai
	http.HandleFunc("/healthcheck", healthCheckHandler)

	// Memulai server dan menangani kesalahan jika ada
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

// Handler untuk endpoint "/healthcheck"
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Menentukan bahwa response adalah JSON
	w.Header().Set("Content-Type", "application/json")

	// Mengirim response JSON sederhana
	fmt.Fprintf(w, `{"status": "ok"}`)
}
