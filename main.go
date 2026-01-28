package main

import (
	"os"

	"github.com/kayibea/mule/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args))
}
