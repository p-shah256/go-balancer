package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	proxy_address  = "localhost:3000"
	server_address = "localhost:1973"
)

var logger = log.New(os.Stderr, "[proxy]", log.LstdFlags)

func main() {
	// create a listener at some address
	proxyListener, err := net.Listen("tcp", proxy_address)
	fmt.Println("Proxy is running at ", proxy_address)
	handleErr(err, "[listener]", logger)

	defer proxyListener.Close()

	// for each connection in the listener, accept it
	for {
		clientConn, err := proxyListener.Accept()
		handleErr(err, "[listener] [acceptor]", logger)

		go handleClientConn(clientConn)

	}

}

// ╭────────────────── create proxy client <-> server ───────────────╮
func handleClientConn(clientConn net.Conn) {
	defer clientConn.Close()
	serverConn, err := net.Dial("tcp", server_address)
	handleErr(err, "[listener] [acceptor] [dialor]", logger)
	defer serverConn.Close()
	io.Copy(clientConn, serverConn)
	io.Copy(serverConn, clientConn)

}

func handleErr(err error, message string, logger *log.Logger) {
	logger.Printf(message + err.Error())
}
