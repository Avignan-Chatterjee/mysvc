package memstore

import (
	"fmt"

	"mysvc/src/domain/entity"
)

type MemStoreRepository interface {
	GetByID(id int64) (*entity.Entity, error)
	CreateEntity(e entity.Entity) error
}

type memStore struct {
	store map[int64]*entity.Entity
}

func NewMemStoreRepository() MemStoreRepository {
	return &memStore{
		store: make(map[int64]*entity.Entity),
	}
}

func (s *memStore) GetByID(id int64) (*entity.Entity, error) {
	if e, ok := s.store[id]; ok {
		return e, nil
	}
	return nil, fmt.Errorf("No entry found!")
}

func (s *memStore) CreateEntity(e entity.Entity) error {
	s.store[e.Id] = &e
	return nil
}
