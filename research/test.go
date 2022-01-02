package main

import (
	"github.com/kr/pretty"
	"github.com/nint8835/monodraw"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./research/example.monopic")
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}

	data, err := monodraw.DecodeMonopic(f)
	if err != nil {
		log.Fatalf("error decoding file: %v\n", err)
	}

	pretty.Print(data)
}
