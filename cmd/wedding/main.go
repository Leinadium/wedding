package main

import (
	"fmt"
	"os"

	"leinadium.dev/wedding/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
