package router

import (
	"net/http"

	"github.com/ryansissom/go-rest-api/internal/handler"
)

func New(ns handler.NewsStorer) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /news", handler.PostNews(ns))
	r.HandleFunc("GET /news", handler.GetAllNews(ns))
	r.HandleFunc("GET /news/{news_id}", handler.GetNewsByID(ns))
	r.HandleFunc("PUT /news/{news_id}", handler.UpdateNewsByID(ns))
	r.HandleFunc("DELETE /news/{news_id}", handler.DeleteNewsByID(ns))

	return r
}
