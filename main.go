package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Iniciar el servidor en el puerto especificado
	fmt.Println("Server started on port:", port)
	http.ListenAndServe(":"+port, nil)
}