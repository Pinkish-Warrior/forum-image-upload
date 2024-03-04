package forum

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestLikePost(t *testing.T) {
	// Create a mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %s", err)
	}
	defer mockDB.Close()

	// Assuming db is a global variable or passed as an argument.
	db = mockDB

	// Define the expected queries and their results
	mock.ExpectQuery(`SELECT like_status FROM likes_dislikes`).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"like_status"}).AddRow(true))

	mock.ExpectExec(`DELETE FROM likes_dislikes`).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Call the function
	err = LikePost(1, 1)
	if err != nil {
		t.Errorf("LikePost() error = %v, expected nil", err)
	}

	// Make sure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
