package main

import (
	"os"

	"yext/m4/confcode/cmd/sud/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
