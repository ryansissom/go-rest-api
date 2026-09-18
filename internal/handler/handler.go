package handler

import (
	"encoding/json"
	"net/http"
	"uuid"

	"github.com/ryansissom/go-rest-api/internal/logger"
)

type NewsStorer interface {
	// Create news from post request body
	Create(NewsPostReqBody) (NewsPostReqBody, error)
	// Find news by its ID
	FindByID(uuid.UUID) (NewsPostReqBody, error)
	// Return all news in the store
	FindAll() ([]NewsPostReqBody, error)
	// Deletes a news item by its ID
	DeleteByID(uuid.UUID) error
	// Updates a news resource by its ID
	UpdateByID(NewsPostReqBody) error
}

type AllNewsResponse struct {
	News []NewsPostReqBody `json:"news"`
}

func PostNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request recieved")
		var newsRequestBody NewsPostReqBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode the request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := newsRequestBody.Validate(); err != nil {
			logger.Error("request validation failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, err := ns.Create(newsRequestBody); err != nil {
			logger.Error("error creating news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetAllNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request recieved")
		news, err := ns.FindAll()

		if err != nil {
			logger.Error("failed to fetch all news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		allNewsResponse := AllNewsResponse{News: news}
		if err := json.NewEncoder(w).Encode(allNewsResponse); err != nil {
			logger.Error("failed to write response", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func GetNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request recieved")
		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id not a valid uuid", "newsId", newsID)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		news, err := ns.FindByID(newsUUID)
		if err != nil {
			logger.Error("news not found", "newsId", newsID, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(&news); err != nil {
			logger.Error("failed to encode", "newsId", newsID)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}

func UpdateNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request received")

		var newsRequestBody NewsPostReqBody
		if err := json.NewDecoder(r.Body).Decode(&newsRequestBody); err != nil {
			logger.Error("failed to decode the request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := newsRequestBody.Validate(); err != nil {
			logger.Error("request validation failed", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if err := ns.UpdateByID(newsRequestBody); err != nil {
			logger.Error("error updating news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func DeleteNewsByID(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		logger.Info("request recieved")
		newsID := r.PathValue("news_id")
		newsUUID, err := uuid.Parse(newsID)
		if err != nil {
			logger.Error("news id not a valid uuid", "newsId", newsID, "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := ns.DeleteByID(newsUUID); err != nil {
			logger.Error("news not found", "newsId", newsID, "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
