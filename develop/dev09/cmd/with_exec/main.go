package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("need 1 arg")
	}
	cmd := exec.Command("wget", flag.Arg(0))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
