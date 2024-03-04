package forum

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TestRegisterPage(t *testing.T) {
	// Parse the "register.html" template

	//IMPORTANT use the following 1st line to run tests and the 2nd when runnnig the server

	// tmpl, err := template.ParseFiles("../templates/register.html")
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		t.Fatalf("Error parsing template: %v", err)
	}

	// Execute the parsed template and check for errors
	testCases := []struct {
		name             string
		requestPath      string
		expectedStatus   int
		expectedTemplate *template.Template
	}{
		{
			name:             "Valid request",
			requestPath:      "/register",
			expectedStatus:   http.StatusOK,
			expectedTemplate: tmpl,
		},
		// Add more test cases for different scenarios and edge cases
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a test HTTP request
			request := httptest.NewRequest("GET", tc.requestPath, nil)
			// Create a test HTTP response recorder
			response := httptest.NewRecorder()

			RegisterPage(response, request)

			// Check the status code is what we expect.
			if response.Code != tc.expectedStatus {
				t.Errorf("Expected HTTP status %d, but got %d", tc.expectedStatus, response.Code)
			}

			// Render the expected template into a buffer
			var buf bytes.Buffer
			err := tc.expectedTemplate.Execute(&buf, nil)
			if err != nil {
				t.Fatalf("Error rendering expected template: %v", err)
			}
			// DEBUGGING Before comparing the response and template content, log them
			fmt.Printf("Actual Response Body: %s\n", response.Body.String())
			fmt.Printf("Expected Template Content: %s\n", buf.String())

			// Compare the buffer's content with the response body
			if !bytes.Equal(buf.Bytes(), response.Body.Bytes()) {
				t.Errorf("Expected response body to match the template, but they are different.")
			}
		})
	}
}
