package service

import (
	"context"
	"github.com/homaderaka/peersmsg"
	"pnode/internal/storage"
)

// Service is a high-level type that wraps the storage.
type Service struct {
	Storage storage.Storage
}

// NewService creates a new Service with initialized storage.
func NewService(s storage.Storage) *Service {
	return &Service{
		Storage: s,
	}
}

// GetMessages gets all stored messages from the service's storage.
// Note: The returned messages should NOT be mutated.
func (s *Service) GetMessages(c context.Context) (m []*peersmsg.Message, err error) {
	m, err = s.Storage.GetMessages(c)
	return
}

// AddMessage adds a new message to the service's storage.
func (s *Service) AddMessage(c context.Context, m *peersmsg.Message) (err error) {
	err = s.Storage.AddMessage(c, m)
	return
}
