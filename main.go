package main

import (
	"os"

	"github.com/abhi-kr-2100/dojo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}