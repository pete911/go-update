package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		Fatalf("get working directory: %v", err)
	}

	goDirs, err := listGoDirs(pwd)
	if err != nil {
		Fatalf("list go dirs: %v", err)
	}

	version, err := GoVersion()
	if err != nil {
		Fatalf("get go version: %v", err)
	}
	Printf("local go version %s\n", version)

	for _, d := range goDirs {
		Printf("go project %s\n", d)
		if err := Update(d, version); err != nil {
			Errorf("update %s: %v", d, err)
		}
	}
}

func listGoDirs(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read %s dir: %w", dir, err)
	}

	var out []string
	for _, e := range entries {
		subDir := filepath.Join(dir, e.Name())
		if e.IsDir() && isGoDir(subDir) {
			out = append(out, subDir)
		}
	}
	return out, nil
}

func isGoDir(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if !e.IsDir() && e.Name() == "go.mod" {
			return true
		}
	}
	return false
}
