package server

import (
	"log"
	"net"
	"testing"
)

func TestServer_AcceptConnections(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn.Write([]byte("Hello!\x00"))
		conn.Close()
	}

}
