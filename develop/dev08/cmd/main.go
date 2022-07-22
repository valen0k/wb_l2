package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type shell struct {
	path    string
	cmdLine string
}

func newShell() (*shell, error) {
	abs, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}
	return &shell{
		path:    abs,
		cmdLine: filepath.Base(abs) + "$ ",
	}, nil
}

func (s *shell) cd(path string) error {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("cd: no such file or directory:", path)
	}
	s.path, err = filepath.Abs(".")
	if err != nil {
		return err
	}
	s.cmdLine = filepath.Base(s.path) + "$ "
	return nil
}

func (s *shell) pwd() {
	fmt.Println(s.path)
}

func (s *shell) echo(str []string) {
	fmt.Println(strings.Join(str, " "))
}

func (s *shell) kill(pids []string) error {
	for _, pid := range pids {
		atoi, err := strconv.Atoi(pid)
		if err != nil {
			return err
		}
		process, err := os.FindProcess(atoi)
		if err != nil {
			fmt.Printf("kill: kill %d failed: no such process\n", atoi)
		}
		err = process.Kill()
		if err != nil {
			fmt.Printf("kill: kill %d failed: no kill procecc\n", atoi)
		}
	}
	return nil
}

func (s *shell) ps() error {
	cmd := exec.Command("ps")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	mShell, err := newShell()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var cmd []string

	for {
		fmt.Print(mShell.cmdLine)
		scanner.Scan()
		cmd = strings.Split(scanner.Text(), " ")

		switch cmd[0] {
		case "cd":
			err = mShell.cd(cmd[1])
			if err != nil {
				log.Fatalln(err)
			}
		case "pwd":
			mShell.pwd()
		case "echo":
			mShell.echo(cmd[1:])
		case "kill":
			err = mShell.kill(cmd[1:])
			if err != nil {
				log.Fatalln(err)
			}
		case "ps":
			err = mShell.ps()
			if err != nil {
				log.Fatalln(err)
			}
		case "exit":
			fmt.Println("exit")
			return
		default:
			fmt.Printf("Command '%s' not found\n", cmd[0])
		}
	}
}
