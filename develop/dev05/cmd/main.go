package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("grep", os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
