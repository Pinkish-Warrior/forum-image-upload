package forum

import (
	"net/http/httptest"
	"testing"
)

func TestCreateSession(t *testing.T) {
	// type args struct {
	// 	w      http.ResponseWriter
	// 	userID int
	// }
	tests := []struct {
		name   string
		userID int
	}{
		{
			name:   "ValidSession",
			userID: 1,
		},
	}
	// TODO: Add test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP response writer
			w := httptest.NewRecorder()
			// Call the CreateSession function
			CreateSession(w, tt.userID)
			// Get the recorded HTTP response
			resp := w.Result()
			// Check if the cookie was set correctly
			cookie := resp.Cookies()[0]
			if cookie.Name != "session-name" {
				t.Errorf("Test case '%s': Expected cookie name 'session-name', got '%s'", tt.name, cookie.Name)
			}

			if cookie.Value == "" {
				t.Errorf("Test case '%s': Expected a non-empty cookie value", tt.name)
			}
			// Check if the session was stored in the map correctly
			storedUserID, ok := sessions[cookie.Value]
			if !ok {
				t.Errorf("Test case '%s': Session not found in sessions map", tt.name)
			}
			if storedUserID != tt.userID {
				t.Errorf("Test case '%s': Expected userID %d, got %d", tt.name, tt.userID, storedUserID)
			}
		})
	}
}
