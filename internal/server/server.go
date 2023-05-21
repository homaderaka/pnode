package server

import (
	"context"
	"fmt"
	"github.com/homaderaka/peersmsg"
	"io"
	"log"
	"net"
	"os"
	"pnode/internal/service"
	"pnode/pkg/peerscmd"
	"strings"
)

type Server struct {
	p peersmsg.Parser
	s *service.Service
	h map[string]peerscmd.CommandHandler
}

func NewServer(p peersmsg.Parser, s *service.Service) *Server {
	return &Server{
		p: p,
		s: s,
		h: peerscmd.NewCommandHandlers(s, p),
	}
}

func (s *Server) AcceptConnections(c context.Context) {
	// TODO: replace with config

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" // Default port number if PORT environment variable is not set
	}

	listen, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go s.handleConnection(c, conn)
	}
}

func (s *Server) handleConnection(c context.Context, conn net.Conn) {
	defer conn.Close()

	for {
		message, err := s.p.NextMessage(conn)
		if err != nil {
			fmt.Printf("Message (%s) received with error %v:", message, err)
			return
		}
		if err == io.EOF {
			log.Println("read error:", err)
			return
		}
		if message == nil {
			continue
		}
		log.Printf("Message (%s) received", message)

		parts := strings.Split(message.String(), " ")
		command := parts[0]
		args := parts[1:]

		handler, ok := s.h[command]
		if !ok {
			conn.Write([]byte(fmt.Sprintf("Error: unknown command (%s)\x00", command)))
			continue
		}

		// TODO: add custom context
		response, err := handler.Execute(c, args)
		if err != nil {
			log.Println(err)
			return
		}

		_, err = conn.Write([]byte(fmt.Sprintf("Response: %s\x00", response)))
		if err != nil {
			log.Println(err)
			return
		}

		// Here you can add your message handling logic
	}
}
