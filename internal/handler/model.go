package handler

import (
	"errors"
	"fmt"
	"net/url"
	"time"
	"uuid"

	"github.com/ryansissom/go-rest-api/internal/store"
)

type NewsPostReqBody struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	CreatedAt string    `json:"created_at"`
	Content   string    `json:"content"`
	Source    string    `json:"source"`
	Tags      []string  `json:"tags"`
}

func (n NewsPostReqBody) Validate() (news store.News, errs error) {
	if n.Author == "" {
		errs = errors.Join(errs, fmt.Errorf("author is empty: %s", n.Author))
	}
	if n.Title == "" {
		errs = errors.Join(errs, fmt.Errorf("title is empty: %s", n.Title))
	}
	if n.Content == "" {
		errs = errors.Join(errs, fmt.Errorf("content is empty: %s", n.Content))
	}
	if n.Summary == "" {
		errs = errors.Join(errs, fmt.Errorf("summary is empty: %s", n.Summary))
	}
	t, err := time.Parse(time.RFC3339, n.CreatedAt)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	if n.Source == "" {
		errs = errors.Join(errs, fmt.Errorf("source is empty: %s", n.Source))
	}
	url, err := url.Parse(n.Source)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	if len(n.Tags) == 0 {
		errs = errors.Join(errs, errors.New("tags cannot be empty"))
	}
	if errs != nil {
		return news, errs
	}
	return store.News{
		ID:        n.ID,
		Author:    n.Author,
		Title:     n.Title,
		Summary:   n.Summary,
		Content:   n.Content,
		CreatedAt: t,
		Source:    url,
		Tags:      n.Tags,
	}, nil
}
