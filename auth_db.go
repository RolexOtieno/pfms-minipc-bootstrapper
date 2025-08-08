package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

// InitDB initializes the SQLite database
func InitDB() {
	var err error
	db, err = sql.Open("sqlite", "./clients.db")
	if err != nil {
		log.Fatal("❌ Failed to open DB:", err)
	}

	// Create the clients table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		fingerprint TEXT NOT NULL,
        authorized BOOLEAN DEFAULT 0
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("❌ Failed to create table:", err)
	}

	log.Println("✅ SQLite Database initialized ")
}

// SaveClient inserts or updates a client in the DB
func SaveClient(name, publicKey string) (string, error) {
	hash := hashKey(publicKey)

	_, err := db.Exec(`
	INSERT OR IGNORE INTO clients (name, fingerprint) VALUES (?, ?)
	`, name, hash)
	return hash, err
}

// IsClientAuthorized checks if a public key is marked as authenticated
func IsClientAuthorized(publicKey string) bool {
	hash := hashKey(publicKey)

	var auth int
	err := db.QueryRow(`
	SELECT authenticated FROM clients WHERE fingerprint = ?
	`, hash).Scan(&auth)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("❌ DB error:", err)
		}
		return false
	}
	return auth == 1
}

// Utility: Hash the public key
func hashKey(key string) string {
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}
