package forum

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestUserExist(t *testing.T) {

	// Replace the db parameter with the mockDB
	tests := []struct {
		name     string
		email    string
		username string
		want     bool
		wantErr  int
	}{
		{
			name:     "UserExistsByEmail",
			email:    "existing@example.com",
			username: "nonexistent",
			want:     true,
			wantErr:  0,
		},
		{
			name:     "UserExistsByUsername",
			email:    "nonexistent@example.com",
			username: "existing",
			want:     true,
			wantErr:  0,
		},
		{
			name:     "UserDoesNotExist",
			email:    "nonexistent@example.com",
			username: "nonexistent",
			want:     false,
			wantErr:  1,
		},
		{
			name:     "ErrorOccurred",
			email:    "error@example.com",
			username: "error",
			want:     false,
			wantErr:  1, // Simulating no rows found
		},
		// Add more test cases for different scenarios here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new mock database connection
			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			if tt.wantErr == 0 {
				// Simulate rows found for non-error cases
				mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users WHERE email=\? OR username=\?`).
					WithArgs(tt.email, tt.username).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1)) // Return 1 count for existence
			} else {
				// Simulate no rows found for error case
				mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users WHERE email=\? OR username=\?`).
					WithArgs(tt.email, tt.username).
					WillReturnRows(sqlmock.NewRows([]string{"count"}))
			}

			got, err := UserExist(tt.email, tt.username, mockDB)
			if (err != nil) != (tt.wantErr != 0) {
				t.Errorf("UserExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)

			// Ensure that all expected SQL queries were executed
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
