package store

import (
	"errors"
	"sync"
	"uuid"
)

// ErrNotFound lets handlers distinguish missing items from server errors.
var ErrNotFound = errors.New("news not found")

// Store is a concurrency-safe in-memory collection of news items.
type Store struct {
	l sync.Mutex
	n []News
}

func New() *Store {
	return &Store{
		l: sync.Mutex{},
		n: []News{},
	}
}

// Create assigns a new ID and appends the item to the collection.
func (s *Store) Create(news News) (News, error) {
	s.l.Lock()
	defer s.l.Unlock()
	news.ID = uuid.New()
	s.n = append(s.n, news)
	return news, nil
}

// FindAll returns the collection of stored news items.
func (s *Store) FindAll() ([]News, error) {
	s.l.Lock()
	defer s.l.Unlock()
	return s.n, nil
}

// FindByID returns the item matching id or ErrNotFound.
func (s *Store) FindByID(id uuid.UUID) (News, error) {
	s.l.Lock()
	defer s.l.Unlock()
	for _, n := range s.n {
		if n.ID == id {
			return n, nil
		}
	}
	return News{}, ErrNotFound
}

// UpdateByID replaces the item with the same ID or returns ErrNotFound.
func (s *Store) UpdateByID(news News) error {
	s.l.Lock()
	defer s.l.Unlock()
	for idx, n := range s.n {
		if n.ID == news.ID {
			s.n[idx] = news
			return nil
		}
	}
	return ErrNotFound
}

// DeleteByID removes the item matching id or returns ErrNotFound.
func (s *Store) DeleteByID(id uuid.UUID) error {
	s.l.Lock()
	defer s.l.Unlock()

	idx := func(id uuid.UUID) int {
		for i, n := range s.n {
			if n.ID == id {
				return i
			}
		}
		return -1
	}(id)

	if idx == -1 {
		return ErrNotFound
	}
	s.n = append(s.n[:idx], s.n[idx+1:]...)
	return nil
}
