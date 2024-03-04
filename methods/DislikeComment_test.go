package forum

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDislikeComment(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection: %v", err)
	}
	defer mockDB.Close()

	// Replace the global 'db' variable with the mockDB
	db = mockDB

	// Test cases
	tests := []struct {
		name              string
		userID, commentID int
		currentLikeStatus bool
		expectedSQL       string
		expectedArgs      []interface{}
		expectedError     error
	}{
		{
			name:              "UserDislikesComment",
			userID:            1,
			commentID:         1,
			currentLikeStatus: false,
			expectedSQL:       "SELECT like_status FROM comments_likes_dislikes WHERE user_id = ? AND comment_id = ?",
			expectedArgs:      []interface{}{1, 1},
			expectedError:     nil,
		},
		{
			name:              "UserUndislikesComment",
			userID:            2,
			commentID:         2,
			currentLikeStatus: true,
			expectedSQL:       "SELECT like_status FROM comments_likes_dislikes WHERE user_id = ? AND comment_id = ?",
			expectedArgs:      []interface{}{2, 2},
			expectedError:     nil,
		},
		// Add more test cases for different scenarios here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Expect the SELECT query with arguments

			// Ensure that all expected SQL queries were executed
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
