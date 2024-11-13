// lesson with anthony gg
package main

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tpc", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	<-s.quitch // 'wait for the quitch channel'
	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("error!", err)
			continue
		}

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error! readLoop! ", err)
			continue // could also close and drop connection, could also check for 'end of files' and other stuff TODO: look into this
		}
		msg := buf[:n]
		fmt.Println(string(msg))
	}
}

func main() {
	server := NewServer(":3000") // "the holy grail of ports"
	server.Start()
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// )

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()

// 	scanner := bufio.NewScanner(conn)
// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		fmt.Printf("Recieved: %s\n", text)
// 		// echo msg to client
// 		conn.Write([]byte(text + "\n"))
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Printf("Error reading from connection: %v", err)
// 	}
// }

// func main() {
// 	listener, err := net.Listen("tcp", ":9000")
// 	if err != nil {
// 		log.Fatalf("Failed to bind to port: %v", err)
// 	}
// 	defer listener.Close()

// 	fmt.Printlin("Server is listening on port 9000...")
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Printf("Failed to accept connection: %v", err)
// 			continue
// 		}
// 		go handleConnection(conn)
// 	}
// }
