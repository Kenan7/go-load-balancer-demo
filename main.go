package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var (
	counter int 

	listenAddr = "localhost:8080"

	backendServers = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
)

func main() {
	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatalf("Error listening: %s", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s", err)
		}

		backend := chooseBackend()
		fmt.Printf("counter=%d, backend=%s\n", counter, backend)

		go func() {
			err := proxy(conn, backend)

			if err != nil {
				log.Printf("Failed to proxy connection: %s", err)
			}
		}()
	}
}

func proxy(conn net.Conn, backend string) error {
	backendConn, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("Failed to connect to backend: %s: %v", backend, err)
	}

	go io.Copy(backendConn, conn)
	go io.Copy(conn, backendConn)

	return nil
}

func chooseBackend() string {
	// roun robin

	s := backendServers[counter%len(backendServers)]
	counter++
	return s
}