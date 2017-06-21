package main

import (
	"os"
	"fmt"
	"encoding/json"
	"github.com/onitake/dashidator/manifest"
)

func main() {
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "manifest.mpd"
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		os.Exit(1)
	}

	mpd, err := manifest.ReadMpd(file)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		os.Exit(1)
	}

	js, _ := json.MarshalIndent(mpd, "", "  ")
	fmt.Println(string(js))
}
