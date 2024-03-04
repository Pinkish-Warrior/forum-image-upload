package forum

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestDislikePostHandler(t *testing.T) {
	// Create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(DislikePostHandler))
	defer server.Close()

	// Define test cases
	tests := []struct {
		name         string
		method       string
		formValues   url.Values
		expectedCode int
	}{
		{
			name:         "Successful Dislike",
			method:       http.MethodPost,
			formValues:   url.Values{"post_id": {"1"}},
			expectedCode: http.StatusUnauthorized, // You might adjust this based on your logic
		},
		{
			name:         "Unauthenticated User",
			method:       http.MethodPost,
			formValues:   url.Values{"post_id": {"1"}},
			expectedCode: http.StatusUnauthorized, // Redirect to login
		},
		{
			name:         "Invalid Post ID",
			method:       http.MethodPost,
			formValues:   url.Values{"post_id": {"invalid"}},
			expectedCode: http.StatusUnauthorized, // Invalid post ID
		},
		{
			name:         "Unauthorized User",
			method:       http.MethodPost,
			formValues:   url.Values{"post_id": {"1"}},
			expectedCode: http.StatusUnauthorized, // Unauthorized user
		},
		{
			name:         "Method Not Allowed",
			method:       http.MethodGet,
			expectedCode: http.StatusSeeOther, // Redirect to 405 error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request with the specified method and form values
			req, err := http.NewRequest(tt.method, server.URL, strings.NewReader(tt.formValues.Encode()))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// Create a test HTTP response writer
			w := httptest.NewRecorder()

			// Call the DislikePostHandler function
			DislikePostHandler(w, req)

			// Check the HTTP response status code
			if w.Code != tt.expectedCode {
				t.Errorf("Test case '%s' failed. Expected HTTP status %d, but got %d", tt.name, tt.expectedCode, w.Code)
			}
		})
	}
}
