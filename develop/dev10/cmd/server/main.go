package main

import (
	"log"
	"net"
)

const (
	network = "tcp"
	address = ":8081"
)

func main() {
	listen, err := net.Listen(network, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			log.Println(err)
		}
	}(listen)
	log.Printf("Server is listening %s", address)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	input := make([]byte, 1024)

	for {
		n, err := conn.Write([]byte("What do you want? "))
		if err != nil || n == 0 {
			log.Println("Write error:", err)
			return
		}

		n, err = conn.Read(input)
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		n, err = conn.Write([]byte("You want: " + string(input)))
		if err != nil || n == 0 {
			log.Println("Write error:", err)
			return
		}
	}
}
