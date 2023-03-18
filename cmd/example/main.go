package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	lab2 "github.com/AlinaDubchak/Lab-2-Go"
)

var (
	inputExp    = flag.String("e", "", "Expression to compute")
	fileWithExp = flag.String("f", "", "File containing expression to compute")
	outFileRes  = flag.String("o", "", "File to output result")
)

func main() {
	flag.Parse()
	var reader io.Reader
	var writer io.Writer

	if *inputExp != "" {
		reader = strings.NewReader(*inputExp)
	} else if *fileWithExp != "" {
		file, err := os.Open(*fileWithExp)
		if err != nil {
			log.Fatalln("Error opening file:\n", err)
		}
		defer file.Close()
		reader = file
	} else {
		log.Fatalln("Either -e or -f flag must be provided")
	}

	if *outFileRes != "" {
		file, err := os.Create(*outFileRes)
		if err != nil {
			log.Fatalln("Error creating file:\n", err)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	computeHandler := lab2.NewComputeHandler(reader, writer)

	err := computeHandler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error computing expression:", err)
		os.Exit(1)
	}
}
