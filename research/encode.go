package main

import (
	"encoding/json"
	"flag"
	"github.com/nint8835/monodraw"
	"log"
	"os"
)

var inFile = flag.String("in", "", ".json file to encode")
var outFile = flag.String("out", "", ".monopic file to write encoded data to")

func main() {
	flag.Parse()

	f, err := os.Open(*inFile)
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer f.Close()

	var data interface{}
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Fatalf("error decoding json: %v", err)
	}

	outF, err := os.Create(*outFile)
	if err != nil {
		log.Fatalf("error opening output file: %v", err)
	}
	defer outF.Close()

	err = monodraw.EncodeMonopic(outF, data)
	if err != nil {
		log.Fatalf("error writing monopic file: %v", err)
	}
}
