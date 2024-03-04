package forum

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestErrorHandler(t *testing.T) {

	// Load environment variables from the .env file

	tests := []struct {
		name         string
		path         string
		expectedCode int
	}{
		{
			name:         "Page not Found",
			path:         "/error/404",
			expectedCode: http.StatusSeeOther,
		},
		{
			name:         "Internal Server Error",
			path:         "/error/500",
			expectedCode: http.StatusSeeOther,
		},
		{
			name:         "Bad Request",
			path:         "/error/400",
			expectedCode: http.StatusSeeOther,
		},
		{
			name:         "Method Not Allowed",
			path:         "/error/405",
			expectedCode: http.StatusSeeOther,
		},
		{
			name:         "UnspecifiedError",
			path:         "/error/404",
			expectedCode: http.StatusSeeOther,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Create a new HTTP request with specific path
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			rr := httptest.NewRecorder()

			//Call the ErrorHandler function with request and repose
			ErrorHandler(rr, req)

			//Check if Response status code matches the expected value
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("Expected status code %d, but got %d", tt.expectedCode, rr.Code)
			}
		})
	}
	// Additional Tests
	server := httptest.NewServer(http.HandlerFunc(DislikePostHandler))
	defer server.Close()
	// 1. Method Not Allowed with Different HTTP Methods
	methods := []string{http.MethodPut, http.MethodDelete, http.MethodPatch}
	for _, method := range methods {
		t.Run("MethodNotAllowed_"+method, func(t *testing.T) {
			// Create a request with the specified method
			req, err := http.NewRequest(method, server.URL, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a test HTTP response writer
			w := httptest.NewRecorder()

			// Call the DislikePostHandler function
			DislikePostHandler(w, req)

			// Check the HTTP response status code
			if w.Code != http.StatusSeeOther {
				t.Errorf("Expected HTTP status %d, but got %d", http.StatusSeeOther, w.Code)
			}
		})
	}

	// 2. Edge Cases for post_id
	edgeCases := []struct {
		name         string
		postID       string
		expectedCode int
	}{
		{
			name:         "Large Post ID",
			postID:       "999999999999999",
			expectedCode: http.StatusUnauthorized, // Adjust based on your logic
		},
		{
			name:         "Negative Post ID",
			postID:       "-1",
			expectedCode: http.StatusUnauthorized, // Adjust based on your logic
		},
	}

	for _, tc := range edgeCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request with the POST method and specified post_id
			req, err := http.NewRequest(http.MethodPost, server.URL, strings.NewReader("post_id="+tc.postID))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// Create a test HTTP response writer
			w := httptest.NewRecorder()

			// Call the DislikePostHandler function
			DislikePostHandler(w, req)

			// Check the HTTP response status code
			if w.Code != tc.expectedCode {
				t.Errorf("Test case '%s' failed. Expected HTTP status %d, but got %d", tc.name, tc.expectedCode, w.Code)
			}
		})
	}
}
