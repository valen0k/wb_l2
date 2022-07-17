package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	var host string

	fmt.Print("➜ ")
	_, err := fmt.Scan(&host)
	if err != nil {
		os.Exit(1)
	}

	//отправляем полученного хоста для запроса текущего времени у сервера
	time, err := ntp.Time(host)
	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, err.Error())
		if err != nil {
			os.Exit(2)
		}
		os.Exit(3)
	}

	fmt.Println(time)
}
