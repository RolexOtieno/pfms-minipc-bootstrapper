/*
Handles logic when a MiniPC sends a download request to /init
Parses the JSON request (deviceId, os, version).
Checks if the MiniPC is authorized.
Responds with either:

	    A download URL (if authorized).
		An error message (if unauthorized).
*/
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// MiniPC request payload
type InitRequest struct {
	Name      string `json:"name"`
	publicKey string `json:"publicKey"`
}

// Server response
type InitResponse struct {
	Status      string `json:"status"`
	DownloadURL string `json:"downloadUrl,omitempty"`
	Message     string `json:"message,omitempty"`
}

// Handler function
func InitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req InitRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("❌ Error decoding JSON:", err)
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	log.Println(" Incoming request: Name=%s, Publickey length=%d", req.Name, len(req.publicKey))

	//Save if new
	_, err = SaveClient(req.Name, req.publicKey)
	if err != nil {
		log.Println("❌Error saving client:", err)
		http.Error(w, "Server errror", http.StatusInternalServerError)
		return
	}

	if !IsClientAuthorized(req.publicKey) {
		log.Println("🔒 Unauthorized access attept:", req.Name)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(InitResponse{
			Status:  "❌unauthorized",
			Message: "Device not recognized or not authenticated yet",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(InitResponse{
		Status:      "✅authorized",
		DownloadURL: "http://localhost:8080/files/linux-installer-v1.0.0.sh",
	})
}

//  /files/
/*  Directory that stores actual install scripts or software packages.
When the MiniPC is authorized, the backend sends a download link that points here.
Static file serving happens from this folder.*/
