package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError bool
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey 12345abcde"},
			},
			expectedKey:   "12345abcde",
			expectedError: false,
		},
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: true,
		},
		{
			name: "Malformed Authorization Header",
			headers: http.Header{
				"Authorization": []string{"Bearer some-token"},
			},
			expectedKey:   "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			// Check if we expected an error and if we got one
			if (err != nil) != tt.expectedError {
				t.Errorf("GetAPIKey() error = %v, expectedError %v", err, tt.expectedError)
				return
			}

			// Check if the returned key matches our expectation
			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() got = %v, want %v", key, tt.expectedKey)
			}
		})
	}
}
