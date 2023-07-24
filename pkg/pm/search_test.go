package pm

import (
	"net/http/httptest"
	"testing"
)

func TestHeadersValidation(t *testing.T) {
	type testCase struct {
		name           string
		requestHeaders map[string]string
		target         string
		expectedOutput string
	}

	testCases := []testCase{
		{
			name: "TestHeadersValidation",
			// requestHeaders: map[string]string{"X-Public": "true"},
			target: "/search",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.target, nil)
			// adicionar los headers al request
			req.Header = addHeadersToRequest(tt.requestHeaders)
			HeadersValidation(req)
			// Validar el assert que no haya error

		})
	}
}
