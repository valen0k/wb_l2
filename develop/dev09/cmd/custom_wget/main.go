package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("need 1 arg")
	}
	fullURL := flag.Arg(0)

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(fullURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}(resp.Body)

	segments := strings.Split(fullURL, "/")
	fileName := "index.html"
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		fileName = segments[len(segments)-1]
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}(file)

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("Downloaded a file %s with size %d", fileName, size)
}
