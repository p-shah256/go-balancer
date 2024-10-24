package main

import (
	"fmt"
	"log"
	"net"
)

const server_address = "localhost:1973"

func main() {
	listener, err := net.Listen("tcp", server_address)
	if err != nil {
		log.Fatalf("Failed to set up server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Echo server is running at", server_address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected from", conn.LocalAddr())

	// Echo the text back to the client
	_, err := conn.Write([]byte("from server: " + "\n"))
	if err != nil {
		log.Printf("Failed to send data back: %v", err)
	}
}
