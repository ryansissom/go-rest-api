package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryansissom/go-rest-api/internal/handler"
)

func Test_PostNews(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.PostNews()(w, r)

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
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.PostNews()(w, r)

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
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.PostNews()(w, r)

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
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.PostNews()(w, r)

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
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.PostNews()(w, r)

			// Assort
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("Expected %d but got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})

	}

}
