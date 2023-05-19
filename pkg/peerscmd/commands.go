package peerscmd

import (
	"fmt"
	"strings"
)

type CommandHandler func(args []string) string

var CommandHandlers = map[string]CommandHandler{
	"send": func(args []string) string {
		if len(args) < 2 {
			return "Error: send command expects 2 arguments"
		}
		// Do something with args[0] and args[1]
		return fmt.Sprintf("Message sent to %s: %s", args[0], args[1])
	},
	"echo": func(args []string) string {
		if len(args) == 0 {
			return "Error: echo command expects at least 1 argument"
		}
		return strings.Join(args, " ")
	},
	"ping": func(args []string) string {
		return "Pong!"
	},
	"trace": func(args []string) string {
		return "Tracing..."
	},
}
