package storage

import (
	"context"
	"github.com/homaderaka/peersmsg"
)

type Storage interface {
	GetMessages(c context.Context) ([]*peersmsg.Message, error)
	AddMessage(c context.Context, m *peersmsg.Message) error
}
