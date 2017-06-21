package main

import (
	"os"
	"fmt"
	"github.com/onitake/dashidator"
)

func main() {
	filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "manifest.mpd"
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Can't open manifest.mpd: %s\n", err)
		os.Exit(1)
	}

	mpd, err := ReadMpd(nil)
	if err != nil {
		fmt.Printf("Error reading manifest: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(mpd)
}
