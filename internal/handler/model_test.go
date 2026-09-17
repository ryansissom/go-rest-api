package handler_test

import (
	"testing"

	"github.com/ryansissom/go-rest-api/internal/handler"
)

func TestNewsPostRequestBody_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		req         handler.NewsPostRequestBody
		expectedErr bool
	}{
		{
			name:        "author empty",
			req:         handler.NewsPostRequestBody{},
			expectedErr: true,
		},
		{
			name: "title empty",
			req: handler.NewsPostRequestBody{
				Author: "test-author",
			},
			expectedErr: true,
		},
		{
			name: "summary empty",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				CreatedAt: "2024-01-01T00:00:00Z",
			},
			expectedErr: true,
		},
		{
			name: "time invalid",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "invalid-time",
			},
			expectedErr: true,
		},
		{
			name: "tags empty",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-01-01T00:00:00Z",
				Source:    "https://test-site.com",
			},
			expectedErr: true,
		},
		{
			name: "validate",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-01-01T00:00:00Z",
				Source:    "https://test-site.com",
				Tags:      []string{"test-tag"},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.req.Validate()

			if tc.expectedErr && err == nil {
				t.Fatalf("expected error but got nil")
			}

			if !tc.expectedErr && err != nil {
				t.Fatalf("expected nil but got error: %s", err)
			}
		})
	}
}
