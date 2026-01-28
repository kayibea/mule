package mule

import (
	"os"
	"path/filepath"
	"strings"
)

type Store struct {
	Path string
}

func DefaultStore() (*Store, error) {
	cache := os.Getenv("XDG_CACHE_HOME")
	if cache == "" {
		cache, _ = os.UserHomeDir()
		cache = filepath.Join(cache, ".cache")
	}
	path := filepath.Join(cache, "mulefile")

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	if _, err := os.OpenFile(path, os.O_CREATE, 0o644); err != nil {
		return nil, err
	}

	return &Store{Path: path}, nil
}

func (s *Store) Load() ([]string, error) {
	b, err := os.ReadFile(s.Path)
	if err != nil {
		return nil, err
	}
	lines := strings.Fields(string(b))
	return lines, nil
}

func (s *Store) Save(files []string) error {
	data := strings.Join(files, "\n") + "\n"
	return os.WriteFile(s.Path, []byte(data), 0o644)
}

func (s *Store) Prune() error {
	return os.WriteFile(s.Path, nil, 0o644)
}
