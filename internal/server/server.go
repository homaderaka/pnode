package server

import (
	"fmt"
	"github.com/homaderaka/peersmsg"
	"io"
	"log"
	"net"
	"os"
	"pnode/pkg/peerscmd"
	"strings"
)

type Server struct {
	p peersmsg.Parser
}

func NewServer(p peersmsg.Parser) *Server {
	return &Server{
		p: p,
	}
}

func (s *Server) AcceptConnections() {
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
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		message, err := s.p.NextMessage(conn)
		if err != nil {
			fmt.Printf("Message (%s) received with error %v:", message, err)
			break
		}
		if err == io.EOF {
			log.Println("read error:", err)
			break
		}
		if message == nil {
			// log.Println("message is not nill")
			continue
		}
		log.Printf("Message (%s) received", message)

		parts := strings.Split(message.String(), " ")
		command := parts[0]
		args := parts[1:]

		handler, ok := peerscmd.CommandHandlers[command]
		if !ok {
			conn.Write([]byte(fmt.Sprintf("Error: unknown command %s\x00", command)))
			continue
		}
		response := handler(args)
		_, err = conn.Write([]byte(fmt.Sprintf("Response: %s\x00", response)))
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("written a response wo errors")

		// Here you can add your message handling logic
	}
}
