/*
This is the entry point of my application

	starting the server

regirters route handlers

	/init- For miniPC requests
	/files/-to serve installer files from disk
*/
package main

import (
	"log"
	"net/http"
)

func main() {
	InitDB()

	// Handle /init for download validation
	http.HandleFunc("/init", InitHandler)

	// Serve static files from the ./files directory at /files/ URL path
	fs := http.FileServer(http.Dir("./files"))
	http.Handle("/files/", http.StripPrefix("/files/", fs))

	log.Println("✅ Server running on http://localhost:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
