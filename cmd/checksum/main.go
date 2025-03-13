// SPDX-FileCopyrightText: 2021 Eric Neidhardt
// SPDX-License-Identifier: MIT
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"hash"
	"io"
	"log"
	"os"
	"strings"
)

var (
	hashAlg  string = "md5"
	filePath string = ""
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

func main() {
	// check arguments
	flag.StringVar(&hashAlg, "check", "md5", "checksum algorithm (md5, sha1, sha256)")
	flag.StringVar(&hashAlg, "c", "md5", "checksum algorithm (md5, sha1, sha256)")
	flag.StringVar(&filePath, "file", filePath, "optional file to check")
	flag.StringVar(&filePath, "f", filePath, "optional file to check")
	flag.Parse()

	// prepare input reader
	var input io.Reader
	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalln("Error reading file", err.Error())
			return
		}
		defer file.Close()
		input = file
	} else if len(flag.Args()) > 0 {
		data := flag.Arg(0)
		input = strings.NewReader(data)
	} else {
		log.Println("Usage: checksum -check=md5 stringTocheck")
		return
	}

	// select hash algorithm
	var hash hash.Hash
	switch hashAlg {
	case "md5":
		hash = md5.New()
	case "sha1":
		hash = sha1.New()
	case "sha256":
		hash = sha256.New()
	case "sha512":
		hash = sha512.New()
	default:
		log.Fatalf("Unknown checksum algorithm requested: %s\n", hashAlg)
		return
	}

	// apply hash algorithm
	if _, err := io.Copy(hash, input); err != nil {
		log.Fatalln("Error while calculating hash", err.Error())
		return
	}

	log.Printf("%s  %x", hashAlg, hash.Sum(nil))
}
