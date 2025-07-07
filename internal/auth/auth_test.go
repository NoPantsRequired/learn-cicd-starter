package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	/// Test cases for GetAPIKey function
	tests := []struct {
		name     string
		headers  http.Header
		expected string
		err      error
	}{
		{
			name:     "Valid API Key",
			headers:  http.Header{"Authorization": []string{"ApiKey my-api-key"}},
			expected: "my-api-key",
			err:      nil,
		},
		{
			name:     "No Authorization Header",
			headers:  http.Header{},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
		{
			name:     "Malformed Authorization Header",
			headers:  http.Header{"Authorization": []string{"Bearer my-api-key"}},
			expected: "",
			err:      errors.New("malformed authorization header"),
		},
		{
			name:     "Empty Authorization Header",
			headers:  http.Header{"Authorization": []string{""}},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
						}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			if err != nil && tt.err != nil {
				return // If both errors are expected, continue to next test case
			}

			if apiKey != tt.expected {
				t.Errorf("expected (%s, %v), got (%s, %v)", tt.expected, tt.err, apiKey, err)
			}
		})
	}

}
