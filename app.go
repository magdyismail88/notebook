package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/magdyismail88/notebook/cmd"
)

const (
	help = `
Usage:

	go run app.go [--argument]

The commands are:
	--build        build dependencies
	--server	     start server
`

	runServer = `
Server is listening on 8888
Ctrl + C
`
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		switch arg {
		case "--help":

			fmt.Print(help)

			break

		case "--build":

			if err := install(); err != nil {
				log.Panic(err)
			}

			fmt.Println("Done")

			break

		case "--server":
			if err := run(); err != nil {
				log.Panic(err)
			}

			break

		default:
			fmt.Print(help)

			break
		}

	}

}

func install() error {

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}

	return nil
}

func run() error {

	a := exec.Command("./bin/notebook-bin")

	log.Print(runServer)

	if err := a.Run(); err != nil {
		return fmt.Errorf("Server error: %s", err)
	}

	return nil

}
