package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	log.Println("application up")

	//парсим все аргументы
	var t time.Duration
	flag.DurationVar(&t, "timeout", time.Second*10, "таймаут на подключение к серверу (по умолчанию 10s)")
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatalln("need 2 arg")
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	//пытаемся приконектиться по tcp с timeout'ом
	conn, err := net.DialTimeout("tcp", host+":"+port, t)
	if err != nil {
		log.Fatalln(err)
	}
	//по завершению закрываем коннект
	defer func(conn net.Conn) {
		log.Println("close connection")
		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	errCh := make(chan error)

	//из сокета в STDOUT
	go func() {
		_, err = io.Copy(os.Stdout, conn)
		errCh <- err
	}()

	//STDIN в сокет
	go func() {
		_, err = io.Copy(conn, os.Stdin)
		errCh <- err
	}()

	//ловим ошибку
	log.Println(<-errCh)
	log.Println("application down")
}
