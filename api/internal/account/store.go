package account

import (
	"errors"
	"sync"
)

var (
	ErrNotFound      = errors.New("account not found")
	ErrAlreadyExists = errors.New("account already exists")
)

type Store struct {
	mu       sync.RWMutex
	accounts map[string]Account
}

func NewStore() *Store {
	return &Store{accounts: make(map[string]Account)}
}

func (s *Store) Create(account Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[account.ID]; exists {
		return ErrAlreadyExists
	}
	s.accounts[account.ID] = account

	return nil
}

func (s *Store) Get(id string) (Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	account, exists := s.accounts[id]
	if !exists {
		return Account{}, ErrNotFound
	}

	return account, nil
}
