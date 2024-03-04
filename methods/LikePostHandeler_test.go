package forum

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestLikePostHandler(t *testing.T) {
	// Create a request with the required form values
	form := url.Values{}
	form.Add("post_id", "test") // Change the post_id as needed

	req := httptest.NewRequest("POST", "/", nil)
	req.Form = form

	//
	// Simulate an authenticated user by setting the user ID in the request context
	req = req.WithContext(context.WithValue(req.Context(), req, "test"))
	// Replace "123" with the actual user ID(req, "test")
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler with the test request
	LikePostHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusSeeOther {
		t.Errorf("Expected status code %d, but got %d", http.StatusSeeOther, rr.Code)
	}

	// You can add more assertions to test the behavior further

	// Example: Check if a redirect header is set
	redirectLocation, headerFound := rr.HeaderMap["Location"]
	if !headerFound {
		t.Error("Expected 'Location' header in the response")
	} else {
		expectedLocation := "/login"
		if redirectLocation[0] != expectedLocation {
			t.Errorf("Expected redirect location %s, but got %s", expectedLocation, redirectLocation[0])
		}
	}
}
