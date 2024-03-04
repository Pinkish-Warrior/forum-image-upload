package forum

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFilterPosts(t *testing.T) {
	// Create a mock HTTP response writer
	w := httptest.NewRecorder()

	// Set a mock session cookie for authentication
	sessionCookie := &http.Cookie{
		Name:  "session-name",
		Value: "1", // Set this to a valid user ID for your test
	}
	request := httptest.NewRequest("POST", "/filter", nil)
	request.AddCookie(sessionCookie)

	// Before calling FilterPosts
	fmt.Println("Before calling FilterPosts: HTTP status code", w.Result().StatusCode)

	// Call the FilterPosts function
	FilterPosts(w, request)

	// After calling FilterPosts
	fmt.Println("After calling FilterPosts: HTTP status code", w.Result().StatusCode)

	// Get the recorded HTTP response
	resp := w.Result()

	// Check the HTTP response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	// You can further check the response content or other expectations here
}

// ALERT This is a alternative for a database mocking

// type MockSessionManager struct{}

// func (m *MockSessionManager) GetAuthenticatedUserID(r *http.Request) (int, bool) {
// 	return 1, true
// }

// type MockDatabase struct{}

// func TestFilterPosts(t *testing.T) {
// 	// Create a mock HTTP response writer
// 	w := httptest.NewRecorder()

// 	// Set up mock dependencies
// 	sm := &MockSessionManager{}
// 	db := &MockDatabase{}

// 	// Create a request
// 	request := httptest.NewRequest("POST", "/filter", nil)

// 	// Call the FilterPosts function with mock dependencies
// 	FilterPosts(w, request, sm, db)

// 	// Get the recorded HTTP response
// 	resp := w.Result()

// 	// Check the HTTP response status code
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, resp.StatusCode)
// 	}

// 	// You can further check the response content or other expectations here
// }
