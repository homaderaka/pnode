package storage

import (
	"context"
	"github.com/homaderaka/peersmsg"
	"sync"
)

type RAM struct {
	Messages []*peersmsg.Message
	M        sync.RWMutex
}

// GetMessages returns a slice of all stored messages.
// Note: The returned messages should NOT be mutated.
func (r *RAM) GetMessages(c context.Context) (m []*peersmsg.Message, err error) {
	select {
	case <-c.Done():
		return nil, c.Err()
	default:
	}

	r.M.RLock()
	defer r.M.RUnlock()

	m = r.Messages
	return
}

func (r *RAM) AddMessage(c context.Context, m *peersmsg.Message) (err error) {
	select {
	case <-c.Done():
		return c.Err()
	default:
	}

	r.M.Lock()
	defer r.M.Unlock()

	r.Messages = append(r.Messages, m)
	return
}
