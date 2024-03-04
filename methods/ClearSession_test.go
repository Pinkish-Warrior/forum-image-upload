package forum

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClearSession(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ClearSession",
			args: args{
				w: httptest.NewRecorder(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClearSession(tt.args.w)
			// Retrieve the response cookies
			cookies := tt.args.w.(*httptest.ResponseRecorder).Result().Cookies()
			// Check if the "session-name" cookie is cleared (MaxAge set to -1)
			found := false
			for _, cookie := range cookies {
				if cookie.Name == "session-name" {
					if cookie.MaxAge != -1 {
						t.Errorf("Expected MaxAge to be -1 for the 'session-name' cookie, got %d", cookie.MaxAge)
					}
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected 'session-name' cookie to be cleared, but it was not found")
			}
		})
	}
}
