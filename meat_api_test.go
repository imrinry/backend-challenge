package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// MockHTTPClient for testing
type MockHTTPClient struct {
	response string
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return &http.Response{
		Body: io.NopCloser(strings.NewReader(m.response)),
	}, nil
}

func TestCountMeat(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "Basic test",
			input: "t-bone t-bone, fatback. pastrami",
			expected: map[string]int{
				"t-bone":   2,
				"fatback":  1,
				"pastrami": 1,
			},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:  "Multiple spaces and punctuation",
			input: "t-bone,  t-bone.  fatback",
			expected: map[string]int{
				"t-bone":  2,
				"fatback": 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countMeat(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestBeefSummaryEndpoint(t *testing.T) {
	// Set up mock response
	mockResponse := "t-bone t-bone, fatback. pastrami"
	mockClient := &MockHTTPClient{response: mockResponse}
	httpClient = mockClient

	app := fiber.New()
	app.Get("/beef/summary", getBeefSummary)

	// Test the endpoint
	resp, err := app.Test(httptest.NewRequest("GET", "/beef/summary", nil))
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Read response body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	// Parse response
	var summary BeefSummary
	err = json.Unmarshal(body, &summary)
	assert.NoError(t, err)

	// Verify response structure and content
	assert.NotNil(t, summary.Beef)
	assert.Equal(t, 2, summary.Beef["t-bone"])
	assert.Equal(t, 1, summary.Beef["fatback"])
	assert.Equal(t, 1, summary.Beef["pastrami"])
}
