package forum

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddCommentHandler(t *testing.T) {
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()
	db = mockDB

	tests := []struct {
		name               string
		request            *http.Request
		expectedStatusCode int
	}{
		{
			name: "ValidPost",
			request: func() *http.Request {
				// Create a valid POST request with required form values
				form := "post_id=1&content=This is a valid comment."
				request := httptest.NewRequest("POST", "/addcomment", strings.NewReader(form))
				request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				// Assume the user is authenticated
				request = AddAuthenticatedUserIDToRequest(request, 1)
				return request
			}(),
			expectedStatusCode: http.StatusSeeOther,
		},
		{
			name: "InvalidPostID",
			request: func() *http.Request {
				// Create a POST request with an invalid post ID (non-integer)
				form := "post_id=invalid&content=This is an invalid post ID."
				request := httptest.NewRequest("POST", "/addcomment", strings.NewReader(form))
				request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				// Assume the user is authenticated
				request = AddAuthenticatedUserIDToRequest(request, 1)
				return request
			}(),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "UnauthorizedUser",
			request: func() *http.Request {
				// Create a valid POST request but with an unauthorized user
				form := "post_id=1&content=This is a comment from an unauthorized user."
				request := httptest.NewRequest("POST", "/addcomment", strings.NewReader(form))
				request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				// Assume the user is not authenticated
				request = AddAuthenticatedUserIDToRequest(request, 0)
				return request
			}(),
			expectedStatusCode: http.StatusUnauthorized,
		},
		// Add more test cases for other scenarios here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			AddCommentHandler(w, tt.request)
		})
	}
}

// Helper function to add the user ID to the request for testing
func AddAuthenticatedUserIDToRequest(request *http.Request, userID int) *http.Request {
	// Encode the user ID into a session-like format (you can adjust this as needed)
	sessionValue := EncodeSessionValue(userID)
	// Add the encoded session value to the request
	request.AddCookie(&http.Cookie{Name: "session", Value: sessionValue})
	return request
}

// Mock implementation of EncodeSessionValue (you can replace this with your actual encoding logic)
func EncodeSessionValue(userID int) string {
	// Encode the user ID as a string (dummy encoding)
	return strconv.Itoa(userID)
}
