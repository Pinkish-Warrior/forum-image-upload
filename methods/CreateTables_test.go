package forum

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateTables(t *testing.T) {
	// Set up an in-memory SQLite database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Assign the in-memory database to the global 'db' variable
	// This simulates the behavior of the 'CreateTables' function
	Init()

	// Run the 'CreateTables' function to create tables in the in-memory database
	CreateTables()

	// to be added further checks here to ensure that tables were created successfully
}
