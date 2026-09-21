package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"uuid"

	"github.com/ryansissom/go-rest-api/internal/logger"
	"github.com/ryansissom/go-rest-api/internal/store"
)

type NewsStorer interface {
	// Create stores a new news item.
	Create(store.News) (store.News, error)
	// FindByID returns one news item by ID.
	FindByID(uuid.UUID) (store.News, error)
	// FindAll returns every news item.
	FindAll() ([]store.News, error)
	// DeleteByID removes a news item by ID.
	DeleteByID(uuid.UUID) error
	// UpdateByID replaces a news item by ID.
	UpdateByID(store.News) error
}

// AllNewsResponse wraps the collection returned by GET /news.
type AllNewsResponse struct {
	News []store.News `json:"news"`
}

// PostNews validates and creates a news item from the request body.
func PostNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")
		var newsRequestBody NewsPostReqBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode the request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		n, err := newsRequestBody.Validate()
		if err != nil {
			logger.Error("request validation failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, err := ns.Create(n); err != nil {
			logger.Error("error creating news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// GetAllNews returns every news item in the store.
func GetAllNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")
		news, err := ns.FindAll()

		if err != nil {
			logger.Error("failed to fetch all news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		allNewsResponse := AllNewsResponse{News: news}
		if err := json.NewEncoder(w).Encode(allNewsResponse); err != nil {
			logger.Error("failed to write response", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// GetNewsByID returns the news item identified by the route parameter.
func GetNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")
		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id not a valid uuid", "newsId", newsID)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		news, err := ns.FindByID(newsUUID)
		if err != nil {
			logger.Error("failed to find news", "newsId", newsID, "error", err)
			if errors.Is(err, store.ErrNotFound) {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		if err := json.NewEncoder(w).Encode(&news); err != nil {
			logger.Error("failed to encode", "newsId", newsID)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}

// UpdateNewsByID validates the body and replaces the item identified by the URL.
func UpdateNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")
		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id not a valid uuid", "newsId", newsID, "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var newsRequestBody NewsPostReqBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode the request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		n, err := newsRequestBody.Validate()
		if err != nil {
			logger.Error("request validation failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		n.ID = newsUUID

		if err := ns.UpdateByID(n); err != nil {
			logger.Error("error updating news", "newsId", newsID, "error", err)
			if errors.Is(err, store.ErrNotFound) {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}
}

// DeleteNewsByID removes the item identified by the route parameter.
func DeleteNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")
		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id not a valid uuid", "newsId", newsID, "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := ns.DeleteByID(newsUUID); err != nil {
			logger.Error("failed to delete news", "newsId", newsID, "error", err)
			if errors.Is(err, store.ErrNotFound) {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
