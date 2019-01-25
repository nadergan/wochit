package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {

	var directoryToScan string
	hashes := make(map[string][]string)

	// Take the directory from Command line and assume
	// current directory if not provided.
	if len(os.Args) > 1 {
		directoryToScan = os.Args[1]
	} else {
		directoryToScan, _ = os.Getwd()
	}

	// Scan and find duplicate files based on MD5 hash calculation
	fmt.Println("Scanning Files In Directory ", directoryToScan, ":")
	files, _ := filepath.Glob(directoryToScan + "/*")
	for _, f := range files {
		fileHash := hashTheFile(f)
		value, exist := hashes[fileHash]
		if exist {
			hashes[fileHash] = append(value, f)
		} else {
			hashes[fileHash] = []string{f}
		}
	}

	// fmt.Println(hashes)

	for h, f := range hashes {
		// Search and print the identical files
		if len(f) > 1 {
			fmt.Printf("\n%d identical files found with hash %40s \n", len(f), h)
		} else {
			break
		}
		// print the actual files list
		for _, fn := range f {
			fmt.Println(fn)
		}

	}
}

func hashTheFile(path string) string {
	// Open the file
	file, err := os.Open(path)
	// Report and Exit on case of error!
	if err != nil {
		panic(err)
	}
	// Init SHA1 generator
	sha1 := sha1.New()
	// Stream file to SHA1 engine
	io.Copy(sha1, file)
	// Convert the []byte to string
	hash := fmt.Sprintf("%x", sha1.Sum(nil))
	// Done with the file
	file.Close()
	// Return the hash
	return hash
}
