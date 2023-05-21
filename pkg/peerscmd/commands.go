package peerscmd

import (
	"context"
	"fmt"
	"github.com/homaderaka/peersmsg"
	"strings"

	"pnode/internal/service" // Replace with your actual service package
)

type Command struct {
	Args []string
	Func CommandHandler
}

type CommandHandler interface {
	Execute(ctx context.Context, args []string) (string, error)
}

type SendCommand struct {
	S *service.Service
	P peersmsg.Parser
}

func (s SendCommand) Execute(ctx context.Context, args []string) (out string, err error) {
	if len(args) < 1 {
		return "", fmt.Errorf("Error: send command expects 1 argument")
	}
	// Here, you would normally send the message

	message, err := s.P.FromString(strings.Join(args, " "))
	if err != nil {
		return
	}

	err = s.S.AddMessage(ctx, &message)
	return fmt.Sprintf("Message sent: %v", args), nil
}

type GetCommand struct {
	s *service.Service
}

func (g GetCommand) Execute(ctx context.Context, args []string) (string, error) {
	messages, err := g.s.GetMessages(ctx)
	if err != nil {
		return "", err
	}

	derefStringsArray := make([]string, len(messages))
	for i, strPtr := range messages {
		derefStringsArray[i] = (*strPtr).String()
	}

	return strings.Join(derefStringsArray, "\n"), nil
}

type EchoCommand struct {
}

func (e EchoCommand) Execute(ctx context.Context, args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("Error: echo command expects at least 1 argument")
	}
	return strings.Join(args, " "), nil
}

func NewCommandHandlers(s *service.Service, p peersmsg.Parser) map[string]CommandHandler {
	return map[string]CommandHandler{
		"send": SendCommand{s, p},
		"get":  GetCommand{s},
		"echo": EchoCommand{},
	}
}
