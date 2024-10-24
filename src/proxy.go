package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// TODO abstract this away into command line args later
const (
	proxy_address  = "localhost:3211"
	server_address = "localhost:1973"
)

var logger = log.New(os.Stderr, "[proxy]", log.LstdFlags)

// accepts all connections at port 3000 as client connection
// forwards all bytes BLINDLY
//
//	client -> proxy -> server
//	client <- proxy <- server
func main() {
	// create a listener at some address
	proxyListener, err := net.Listen("tcp", proxy_address)
	fmt.Println("Proxy is running at ", proxy_address)
	handleErr(err, "[listener]", logger)

	defer proxyListener.Close()

	// for each connection in the listener, accept it and spawn a thread to handle that connection
	for {
		clientConn, err := proxyListener.Accept()
		handleErr(err, "[listener] [acceptor]", logger)

		go handleClientConn(clientConn)
	}
}

// ╭────────────────── create proxy client <-> server ───────────────╮
//  1. connect to server
//  2. copy bytes
//     from client -> server and
//     then server <- client
func handleClientConn(clientConn net.Conn) {
	defer clientConn.Close()
	serverConn, err := net.Dial("tcp", server_address)
	handleErr(err, "[listener] [acceptor] [dialor]", logger)
	defer serverConn.Close()
	io.Copy(clientConn, serverConn)
	io.Copy(serverConn, clientConn)
}

func handleErr(err error, message string, logger *log.Logger) {
	if err != nil {
		logger.Printf(message + ": " + err.Error())
		os.Exit(1)
	}
}
