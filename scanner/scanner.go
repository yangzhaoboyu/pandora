package scanner

import (
	"log"
	"os"
	"path/filepath"
)

// Scan all the files in the directory and write them to Chan
func Visit(files chan string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			files <- path
		}
		return nil
	}
}
