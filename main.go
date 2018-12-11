package main

import (
	"os"

	"github.com/berfinsari/couchdbbeat/cmd"

	_ "github.com/berfinsari/couchdbbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
