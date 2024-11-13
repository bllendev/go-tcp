package main

import (
	"bufio"
	"net"
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	go main()               // Start the server in a goroutine
	time.Sleep(time.Second) // give time

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	message := "Hello world server"
	_, err = conn.Write([]byte(message + "\n"))
	if err != nil {
		t.Fatalf("Failed to read response from server: %v", err)
	}

	// read response
	response, err := bufio.NewReader(conn).ReadSTring("\n")
	if err != nil {
		t.Fatalf("")
	}
}
