package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/kayibea/mule/internal/mule"
)

func runList(args []string) int {
	if len(args) > 0 {
		fmt.Fprintln(os.Stderr, "list: unexpected arguments")
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

	fmt.Println(strings.Join(files, "\n"))

	return 0
}
