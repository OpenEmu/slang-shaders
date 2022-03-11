package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pierrre/archivefile/zip"
)

func main() {
	log.SetFlags(0)

	matches, err := filepath.Glob("shaders/*")
	if err != nil {
		log.Fatalf("Error glob: %v", err)
	}

	var files []string
	for _, m := range matches {
		fi, err := os.Stat(m)
		if err != nil || !fi.IsDir() {
			continue
		}
		files = append(files, m)
	}

	if err := os.MkdirAll("_build", os.ModePerm); err != nil {
		log.Fatalf("Error creating _build directory: %v", err)
	}

	for _, fi := range files {
		name := filepath.Join("_build", filepath.Base(fi)+".oeshaderplugin")
		log.Printf("Building shader %q to %q\n", fi, name)
		err := zip.ArchiveFile(fi+string(filepath.Separator), name, nil)
		if err != nil {
			log.Fatalf("Error archiving %s: %v", name, err)
		}
	}
}
