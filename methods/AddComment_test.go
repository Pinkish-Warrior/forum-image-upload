package forum

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddComment(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Faield to create mock database conection: %v", err)
	}
	defer mockDB.Close()
	// Replace the global 'db' variable with the mockDB
	db = mockDB

	tests := []struct {
		name    string
		postID  int
		userID  int
		content string
		wantErr bool
	}{
		{
			name:    "ValidComment",
			postID:  1,
			userID:  1,
			content: "This is a valid comment.",
			wantErr: false,
		},
		{
			name:    "InvalidCommentEmptyContent",
			postID:  1,
			userID:  1,
			content: "",
			wantErr: true,
		},
		{
			name:    "InvalidPostID",
			postID:  -1,
			userID:  1,
			content: "This is an invalid post ID.",
			wantErr: true,
		},
		// Add more test cases for other scenarios here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				mock.ExpectExec("INSERT INTO comments").WillReturnError(sql.ErrNoRows)
			} else {
				mock.ExpectExec("INSERT INTO comments").WillReturnResult(sqlmock.NewResult(1, 1))
			}
			err := AddComment(tt.postID, tt.userID, tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddComment() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
