package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

var hashAlg string

func main() {
	flag.StringVar(&hashAlg, "hash", "md5", "checksum algorithm (md5, sha1, sha256)")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Usage: checksum -hash=sha1 <File>")
		return
	}
	filePath := flag.Arg(0)
	fmt.Printf("Calculating %s for %s\n", hashAlg, filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	var hash hash.Hash
	switch hashAlg {
	case "md5":
		hash = md5.New()
	case "sha1":
		hash = sha1.New()
	case "sha256":
		hash = sha256.New()
	default:
		fmt.Printf("Unknown hashsum algorithm requested %s\n", hashAlg)
		return
	}

	if _, err := io.Copy(hash, file); err != nil {
		fmt.Printf("Error while calculating hash %s\n", err.Error())
		return
	}

	fmt.Printf("%x", hash.Sum(nil))
}
