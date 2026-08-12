package device

import (
	"errors"
	"sync"
)

var (
	ErrNotFound      = errors.New("device not found")
	ErrAlreadyExists = errors.New("device already exists")
)

type Store struct {
	mu      sync.RWMutex
	devices map[string]Device
}

func NewStore() *Store {
	return &Store{
		devices: make(map[string]Device),
	}
}

func (s *Store) Create(device Device) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.devices[device.ID]; exists {
		return ErrAlreadyExists
	}

	s.devices[device.ID] = device
	return nil
}

func (s *Store) Get(id string) (Device, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	device, exists := s.devices[id]
	if !exists {
		return Device{}, ErrNotFound
	}

	return device, nil
}
