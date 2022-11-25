package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	pathPtr := flag.String("path", "", "the path to the directory to sort (required)")
	extsPtr := flag.String("exts", "RAF,JPG,JPEG,DNG", "a comma-separated list of extensions to sort")
	flag.Parse()

	path := *pathPtr

	if path == "" {
		fmt.Println(`img-mv: you need to provide a path. Usage:`)
		flag.PrintDefaults()
		os.Exit(1)
	}

	extsSlice := strings.Split(*extsPtr, ",")
	exts := make(map[string]bool)
	for _, ext := range extsSlice {
		ext = strings.TrimSpace(ext)
		ext = strings.TrimLeft(ext, ".")
		exts["."+ext] = true
	}

	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		// if the entry's extension is not one we want
		if !exts[filepath.Ext(entry.Name())] {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			panic(err)
		}
		date := info.ModTime().Format("2006-01-02")
		dateDir := filepath.Join(path, date)
		err = os.MkdirAll(dateDir, 0777)
		if err != nil {
			panic(err)
		}

		os.Rename(filepath.Join(path, entry.Name()), filepath.Join(dateDir, entry.Name()))

	}
}
