package main

import (
	"fmt"
	// "github.com/go-gota/gota/dataframe"
	// "github.com/jroimartin/gocui"
	// "golang.org/x/image/colornames"
)

func main() {
	// generateData()
	fmt.Println("Joe")
	rawSeqs := readData()
	fmt.Println("Raw Seqs from file", rawSeqs, "---")
	hashedSeqs := generateHashLookup(rawSeqs)
	fmt.Println("Seqs to hash mapping", hashedSeqs, "---")

	pairwiseAligns := pairwiseAlign(hashedSeqs)
	fmt.Println(pairwiseAligns)
}
