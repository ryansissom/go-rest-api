package handler_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"uuid"

	"github.com/ryansissom/go-rest-api/internal/handler"
)

func Test_PostNews(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewsStorer
		body           io.Reader
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			body:           strings.NewReader(``),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid request body ",
			body: strings.NewReader(`
			{
			"id" : "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author" : "test-author"
			"title" : "test-title"
			"summary" : "test-summary"
			"created_at" : "2024-01-01T00:00:00Z"
			"source" : "https://example.com"
			}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "db error",
			body: strings.NewReader(`
			{
			"id" : "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author" : "test-author",
			"title" : "test-title",
			"summary" : "test-summary",
			"created_at" : "2024-01-01T00:00:00Z",
			"source" : "https://example.com",
			"tags" : ["test"]
			}`),
			store:          mockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			body: strings.NewReader(`
			{
			"id" : "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author" : "test-author",
			"title" : "test-title",
			"summary" : "test-summary",
			"created_at" : "2024-01-01T00:00:00Z",
			"source" : "https://example.com",
			"tags" : ["test"]
			}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", tc.body)

			// Act
			handler.PostNews(tc.store)(w, r)

			// Assort
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected %d but got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})

	}

}

func Test_GetAllNews(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "db error",
			store:          mockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			store:          mockNewsStore{},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.GetAllNews(tc.store)(w, r)

			// Assort
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected %d but got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})

	}

}

func Test_GetNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewsStorer
		newsID         string
		expectedStatus int
	}{
		{
			name:           "invalid news id",
			store:          mockNewsStore{},
			newsID:         "invalid-uuid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "db error",
			store:          mockNewsStore{errState: true},
			newsID:         uuid.New().String(),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			store:          mockNewsStore{},
			newsID:         uuid.New().String(),
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			r.SetPathValue("news_id", tc.newsID)
			// Act
			handler.GetNewsByID(tc.store)(w, r)

			// Assort
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected %d but got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})

	}

}
func Test_UpdateNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		body           io.Reader
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			body:           strings.NewReader(`{`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid request body ",
			body: strings.NewReader(`
			{
			"id" : "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author" : "test-author"
			"title" : "test-title"
			"summary" : "test-summary"
			"created_at" : "2024-01-01T00:00:00Z"
			"source" : "https://example.com"
			}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "db error",
			body: strings.NewReader(`
			{
			"id" : "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author" : "test-author",
			"title" : "test-title",
			"summary" : "test-summary",
			"created_at" : "2024-01-01T00:00:00Z",
			"source" : "https://example.com",
			"tags" : ["test"]
			}`),
			store:          mockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "success",
			body: strings.NewReader(`
			{
			"id" : "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author" : "test-author",
			"title" : "test-title",
			"summary" : "test-summary",
			"created_at" : "2024-01-01T00:00:00Z",
			"source" : "https://example.com",
			"tags" : ["test"]
			}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", tc.body)

			// Act
			handler.UpdateNewsByID(tc.store)(w, r)

			// Assort
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected %d but got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})

	}

}
func Test_DeleteNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		store          mockNewsStore
		newsID         string
		expectedStatus int
	}{
		{
			name:           "invalid news id",
			store:          mockNewsStore{},
			newsID:         "invalid-uuid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "db error",
			store:          mockNewsStore{errState: true},
			newsID:         uuid.New().String(),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			store:          mockNewsStore{},
			newsID:         uuid.New().String(),
			expectedStatus: http.StatusNoContent,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			r.SetPathValue("news_id", tc.newsID)

			// Act
			handler.DeleteNewsByID(tc.store)(w, r)

			// Assort
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected %d but got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})

	}

}

type mockNewsStore struct {
	errState bool
}

func (m mockNewsStore) Create(_ handler.NewsPostReqBody) (news handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) FindByID(_ uuid.UUID) (news handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) FindAll() (news []handler.NewsPostReqBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) DeleteByID(_ uuid.UUID) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}

func (m mockNewsStore) UpdateByID(_ handler.NewsPostReqBody) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}
