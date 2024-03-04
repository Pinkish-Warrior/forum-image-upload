package forum

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddPost(t *testing.T) {

	tests := []struct {
		name               string
		request            *http.Request
		expectedStatusCode int // Expected HTTP status code in the response
	}{
		{
			name: "ValidPost",
			request: func() *http.Request {
				// Create a request with form values for a valid post
				form := "category=example&title=ExampleTitle&content=ExampleContent"
				request := httptest.NewRequest("POST", "/addpost", strings.NewReader(form))
				request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return request
			}(),
			expectedStatusCode: http.StatusSeeOther, // Expected redirect status
		},
		{
			name: "InvalidPostMissingCategory",
			request: func() *http.Request {
				// Create a request with missing category field
				form := "title=ExampleTitle&content=ExampleContent"
				request := httptest.NewRequest("POST", "/addpost", strings.NewReader(form))
				request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return request
			}(),
			expectedStatusCode: http.StatusSeeOther,
		},
		{
			name: "InvalidPostMissingTitle",
			request: func() *http.Request {
				// Create a request with missing title field
				form := "category=example&content=ExampleContent"
				request := httptest.NewRequest("POST", "/addpost", strings.NewReader(form))
				request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return request
			}(),
			expectedStatusCode: http.StatusSeeOther,
		},
		{
			name: "InvalidPostMissingTitle",
			request: func() *http.Request {
				// Create a request with missing title field
				form := "category=example&content=ExampleContent"
				request := httptest.NewRequest("POST", "/addpost", strings.NewReader(form))
				request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return request
			}(),
			expectedStatusCode: http.StatusSeeOther,
		},
		{
			name: "InvalidPostMissingContent",
			request: func() *http.Request {
				// Create a request with missing content field
				form := "category=example&title=ExampleTitle"
				request := httptest.NewRequest("POST", "/addpost", strings.NewReader(form))
				request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return request
			}(),
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "UnauthenticatedUser",
			request: func() *http.Request {
				form := "category=example&title=ExampleUser"
				request := httptest.NewRequest("POST", "/addpost", strings.NewReader(form))
				request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return request
			}(),
			expectedStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		tt := tt // create a new variable to avoid variable capture in closures
		t.Run(tt.name, func(t *testing.T) {

			//Crete response recorder
			w := httptest.NewRecorder()

			// Call AddPost with the response recorder, request, and mockGetAuthenticatedUserID
			AddPost(w, tt.request)

			// Check the HTTP response status code
			if w.Code != tt.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatusCode, w.Code)
			}
		})
	}
}
