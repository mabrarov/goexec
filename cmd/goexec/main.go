package main

import (
	"log"
	"os"

	"github.com/mabrarov/goexec/internal/pkg/run"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	code, err := run.Cmd(os.Args[1], os.Args[2:]...)
	if err != nil {
		log.Fatalf("Failed to run command: %s.", err)
	}
	os.Exit(code)
}
