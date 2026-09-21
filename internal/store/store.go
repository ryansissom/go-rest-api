package store

import (
	"errors"
	"sync"
	"uuid"
)

var ErrNotFound = errors.New("news not found")

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

func (s *Store) Create(news News) (News, error) {
	s.l.Lock()
	defer s.l.Unlock()
	news.ID = uuid.New()
	s.n = append(s.n, news)
	return news, nil
}

func (s *Store) FindAll() ([]News, error) {
	s.l.Lock()
	defer s.l.Unlock()
	return s.n, nil
}

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
