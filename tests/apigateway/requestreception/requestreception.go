package requestreception

import (
	"net/http"
	"testing"
)

func TestParser_ParseRequest(t *testing.T) {
	// Create a new Parser instance
	parser := NewParser(1024) // Assuming maximum request size of 1024 bytes

	// Create a sample HTTP request for testing
	req, err := http.NewRequest("GET", "http://example.com/path?param=value", nil)
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}
	// Add some headers if needed
	req.Header.Set("Content-Type", "application/json")

	// Call the ParseRequest method
	data, err := parser.ParseRequest(req)

	// Check if the parsing was successful
	if err != nil {
		t.Errorf("ParseRequest failed: %v", err)
	}

	// Perform additional assertions as needed
	// Example:
	// if data.SomeField != expectedValue {
	//     t.Errorf("Unexpected value for SomeField. Expected: %v, Got: %v", expectedValue, data.SomeField)
	// }
}

// Add more test cases as needed for different scenarios
