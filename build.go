package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pierrre/archivefile/zip"
)

func main() {
	matches, err := filepath.Glob("shaders/*")
	if err != nil {
		fmt.Printf("error glob: %v\n", err)
		os.Exit(1)
	}

	var files []string
	for _, m := range matches {
		fi, err := os.Stat(m)
		if err != nil || !fi.IsDir() {
			continue
		}
		files = append(files, m)
	}

	os.MkdirAll("_build", os.ModePerm)

	for _, fi := range files {
		name := filepath.Join("_build", filepath.Base(fi)+".oeshaderplugin")
		fmt.Printf("building shader %q to %q\n", fi, name)
		err := zip.ArchiveFile(fi+string(filepath.Separator), name, nil)
		if err != nil {
			// boo
		}
	}
}
