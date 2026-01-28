package cli

import (
	"fmt"
	"os"

	"github.com/kayibea/mule/internal/mule"
)

func runPrune(args []string) int {
	if len(args) > 0 {
		fmt.Fprintln(os.Stderr, "clear: unexpected arguments")
		return 1
	}

	store, err := mule.DefaultStore()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	err = store.Prune()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}
