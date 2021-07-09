package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 512)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("unable to write data")
		}
	}
}

func basic_echo_server() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		go echo(conn)
	}
}

func buffered_echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes:%s", len(s), s)

	log.Println("Writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}

	writer.Flush()
}

func bufferred_echo_server() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		go buffered_echo(conn)
	}
}

func bufferred_echo_with_copy(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func bufferred_echo_server_improved() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		go bufferred_echo_with_copy(conn)
	}
}
func main() {
	bufferred_echo_server_improved()
}
