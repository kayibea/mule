package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kayibea/mule/internal/mule"
)

func runAdd(args []string) int {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)

	appendMode := fs.Bool("a", false, "append FILES to the clipboard")

	fs.SetOutput(os.Stdout)

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 1
	}

	files := fs.Args()
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "add: requires at least one file")
		return 1
	}

	store, err := mule.DefaultStore()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	var resolved []string
	for _, f := range files {
		p, err := filepath.Abs(f)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		resolved = append(resolved, p)
	}

	if *appendMode {
		err = mule.Append(store, resolved)
	} else {
		err = mule.Set(store, resolved)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}
