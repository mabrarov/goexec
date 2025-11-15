package main

import (
	"flag"
	"log"
	"os"

	"github.com/mabrarov/goexec/internal/pkg/run"
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if flag.NArg() == 0 {
		return
	}
	code, err := run.Cmd(flag.Arg(0), flag.Args()[1:]...)
	if err != nil {
		log.Printf("Failed to run command: %s.", err)
		os.Exit(-1)
	}
	os.Exit(code)
}
