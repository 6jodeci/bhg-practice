package main

import (
	"io"
	"log"
	"net"
)

// a handler function that simply echoes received data
func echo(conn net.Conn) {
	defer conn.Close()

	// a buffer to store recived data
	b := make([]byte, 512)
	for {
		// receive data via conn.Read into a buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		// send data via conn.Write
		log.Println("Write data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	// bind to TCP port 20080 on all interfaces
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("listening on 0.0.0.0:20080")

	for {
		// wait for conntecions. Create net.Conn on connection established
		conn, err := listener.Accept()
		log.Println("Recived connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// handle the connection. Using gorutine for concurrency
		go echo(conn)
	}
}
