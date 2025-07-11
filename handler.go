package main

import (
	"encoding/json"
	"net/http"
)

// MiniPC request payload
type InitRequest struct {
	DeviceID string `json:"deviceId"`
	OS       string `json:"os"`
	Version  string `json:"version"`
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
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	if !IsAuthorizedDevice(req.DeviceID) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(InitResponse{
			Status:  "unauthorized",
			Message: "Device not recognized",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(InitResponse{
		Status:      "authorized",
		DownloadURL: "http://localhost:8080/files/linux-installer-v1.0.0.sh",
	})
}
