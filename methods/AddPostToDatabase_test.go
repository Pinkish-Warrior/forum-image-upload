package forum

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddPostToDatabase(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection: %v", err)
	}
	defer mockDB.Close()

	// Replace the global 'db' variable with the mockDB
	db = mockDB

	type args struct {
		category   string
		title      string
		content    string
		userID     int
		image_path string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ValidPost",
			args: args{
				category: "example",
				title:    "ExampleTitle",
				content:  "ExampleContent",
				userID:   1,
			},
			wantErr: false,
		},
		{
			name: "InvalidPostEmptyCategory",
			args: args{
				category: "",
				title:    "ExmpleTitle",
				content:  "ExampleContent",
				userID:   1,
			},
			wantErr: true,
		},
		{
			name: "InvalidCPostEmptyTitle",
			args: args{
				category: "example",
				title:    "",
				content:  "ExampleContent",
				userID:   1,
			},
			wantErr: true,
		},
		{
			name: "InvalidPostID",
			args: args{
				category: "example",
				title:    "ExampleTitle",
				content:  "ExampleContent",
				userID:   -1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//set up the mock expectations
			if tt.wantErr {
				// If we expect an error, return an error result from the mock
				mock.ExpectExec("INSERT INTO posts").WillReturnError(errors.New("expected error"))
			} else {
				// If we don't expect an error, return a successful result from the mock
				mock.ExpectExec("INSERT INTO posts").WillReturnResult(sqlmock.NewResult(1, 1))
				// Call the function and check for errors
				err := AddPostToDatabase(tt.args.category, tt.args.title, tt.args.content, tt.args.userID, tt.args.image_path)

				// Verify the error condition based on 'wantErr'
				if (err != nil) != tt.wantErr {
					t.Errorf("Test case '%s': AddPostToDatabase() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				}
				// Ensure that all expected SQL queries were executed
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("Test case '%s': there were unfulfilled expectations:  %s", tt.name, err)
				}
			}
		})
	}
}
