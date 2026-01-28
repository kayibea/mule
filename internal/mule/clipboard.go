package mule

import (
	"fmt"
	"os"
	"sort"
)

func Set(store *Store, entries []string) error {
	return write(store, entries)
}

func Append(store *Store, entries []string) error {
	existing, err := store.Load()
	if err != nil {
		return err
	}

	entries = append(entries, existing...)
	return write(store, entries)
}

func write(store *Store, entries []string) error {
	sort.Strings(entries)
	entries = unique(entries)

	for _, f := range entries {
		if _, err := os.Stat(f); err != nil {
			return fmt.Errorf("no such file: %s", f)
		}
	}

	return store.Save(entries)
}

func unique(xs []string) []string {
	out := xs[:0]
	for i, x := range xs {
		if i == 0 || x != xs[i-1] {
			out = append(out, x)
		}
	}
	return out
}
