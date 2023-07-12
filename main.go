package main

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	// "github.com/jroimartin/gocui"
	// "golang.org/x/image/colornames"
)

func main() {
	// Create the reads
	fmt.Println("Generating Data")
	nAligns := 10
	seqLenMax := 12
	seqLenMin := 8
	sourceFile := "./dataGen/covid.seq"
	outputFilePath := "./dataGen/covidReads"
	generateData(sourceFile, outputFilePath, nAligns, seqLenMax, seqLenMin)

	fmt.Println("Reading Data")
	rawSeqs := readData()
	fmt.Println("Raw Seqs from file", rawSeqs, "---")
	hashedSeqs := generateHashLookup(rawSeqs)
	fmt.Println("Seqs to hash mapping", hashedSeqs, "---")

	pairwiseAligns := pairwiseAlign(hashedSeqs)
	fmt.Println(pairwiseAligns)
}
