package handler_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/ryansissom/go-rest-api/internal/handler"
	"github.com/ryansissom/go-rest-api/internal/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewsPostRequestBody_Validate(t *testing.T) {
	type expectation struct {
		err  string
		news store.News
	}
	testCases := []struct {
		name         string
		req          handler.NewsPostReqBody
		expectations expectation
	}{
		{
			name: "author empty",
			req:  handler.NewsPostReqBody{},
			expectations: expectation{
				err: "author is empty",
			},
		},
		{
			name: "title empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
			},
			expectations: expectation{
				err: "title is empty",
			},
		},
		{
			name: "content empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title:  "test-title",
			},
			expectations: expectation{
				err: "content is empty",
			},
		},
		{
			name: "summary empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title:  "test-title",
			},
			expectations: expectation{
				err: "summary is empty",
			},
		},
		{
			name: "time invalid",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "invalid-time",
			},
			expectations: expectation{
				err: `parsing time "invalid-time"`,
			},
		},
		{
			name: "source invalid",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-01-01T00:00:00Z",
			},
			expectations: expectation{
				err: "source is empty",
			},
		},
		{
			name: "tags empty",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-01-01T00:00:00Z",
				Source:    "https://test-site.com",
			},
			expectations: expectation{
				err: "tags cannot be empty",
			},
		},
		{
			name: "validate",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				Content:   "test-content",
				CreatedAt: "2024-01-01T00:00:00Z",
				Source:    "https://test-site.com",
				Tags:      []string{"test-tag"},
			},
			expectations: expectation{
				news: store.News{
					Author:  "test-author",
					Title:   "test-title",
					Summary: "test-summary",
					Content: "test-content",
					Tags:    []string{"test-tag"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			news, err := tc.req.Validate()

			if tc.expectations.err != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectations.err)
			} else {
				assert.NoError(t, err)

				parsedTime, parseErr := time.Parse(time.RFC3339, tc.req.CreatedAt)
				require.NoError(t, parseErr)
				tc.expectations.news.CreatedAt = parsedTime

				parsedSource, err := url.Parse(tc.req.Source)
				require.NoError(t, err)
				tc.expectations.news.Source = parsedSource

				assert.Equal(t, tc.expectations.news, news)
			}
		})
	}
}
