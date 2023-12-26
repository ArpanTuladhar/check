package config

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestLoadAppConfig(t *testing.T) {
	_ = godotenv.Load("../../.env")

	config := LoadAppConfig()

	if config.App.Port == "" {
		t.Error("App port is empty")
	}
}

func TestConnectToDatabase(t *testing.T) {
	_ = godotenv.Load("../../.env")

	config := LoadAppConfig()

	db, err := config.Connect()
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SHOW TABLES LIKE 'users'")
	if err != nil {
		t.Fatalf("Error querying the database: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		t.Error("Expected table 'users' to exist, but it does not.")
	}
}
