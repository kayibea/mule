package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kayibea/mule/internal/mule"
	"github.com/kayibea/mule/internal/util"
)

func runMove(args []string) int {
	fs := flag.NewFlagSet("move", flag.ContinueOnError)

	verbose := fs.Bool("v", false, "explain what being done")

	fs.SetOutput(os.Stdout)

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 1
	}

	if fs.NArg() > 0 {
		fmt.Fprintln(os.Stderr, "move: unexpected arguments")
		return 1
	}

	store, err := mule.DefaultStore()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	files, err := store.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	for _, src := range files {
		base := filepath.Base(src)
		dst := filepath.Join(".", base)

		if err := util.Copy(src, dst); err != nil {
			fmt.Fprintf(os.Stderr, "move %s: %v\n", src, err)
			return 1
		}

		if err := os.RemoveAll(src); err != nil {
			fmt.Fprintf(os.Stderr, "move %s: %v\n", src, err)
			return 1
		}

		if *verbose {
			fmt.Printf("moved: '%s'\n", src)
		}
	}

	return 0
}
