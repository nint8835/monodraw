package main

import (
	"encoding/json"
	"flag"
	"github.com/nint8835/monodraw"
	"log"
	"os"
)

var inFile = flag.String("in", "", ".monopic file to decode")
var outFile = flag.String("out", "", ".json file to write decoded data to")

func main() {
	flag.Parse()

	f, err := os.Open(*inFile)
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer f.Close()

	data, err := monodraw.DecodeMonopic(f)
	if err != nil {
		log.Fatalf("error decoding file: %v", err)
	}

	outF, err := os.Create(*outFile)
	if err != nil {
		log.Fatalf("error opening output file: %v", err)
	}
	defer outF.Close()

	err = json.NewEncoder(outF).Encode(data)
	if err != nil {
		log.Fatalf("error encoding json: %v", err)
	}
}
