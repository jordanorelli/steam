package main

import (
	"fmt"
	"github.com/jordanorelli/steam"
	"io"
	"os"
)

func bail(code int, t string, args ...interface{}) {
	var out io.Writer
	if code == 0 {
		out = os.Stdout
	} else {
		out = os.Stderr
	}
	fmt.Fprintf(out, t+"\n", args...)
	os.Exit(code)
}

func main() {
	key := os.Getenv("STEAM_KEY")
	if key == "" {
		bail(1, "no steam key provided. use the environment variable STEAM_KEY to provide an api key")
	}

	client := steam.NewClient(key)
	if len(os.Args) < 2 {
		bail(1, "supply a subcommand pl0x")
	}
	cmd, ok := commands[os.Args[1]]
	if !ok {
		bail(1, "no such subcommand %s", os.Args[1])
	}
	cmd.handler(client)
}
