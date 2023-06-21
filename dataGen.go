// package dataGen
package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/go-gota/gota/dataframe"
	// "github.com/go-gota/gota/series"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Seq struct {
	StartPos   int // Cannot use uints because gota cannot handle them
	Length     int
	BPSequence string
}

func generateData() {
	// Seed for reproducibility
	rand.Seed(0)
	data, err := os.ReadFile("./dataGen/covid.seq")
	check(err)

	// Create the reads
	nAligns := 10
	seqLenMax := 12
	seqLenMin := 8
	// maybe use a dataframe/like object and append the indicies and lengths as well
	seqReads := [][]byte{}
	seqReadStructs := []Seq{}
	for n := 0; n < nAligns; n++ {
		randSeqLen := rand.Intn(seqLenMax-seqLenMin) + seqLenMin
		startIndex := rand.Intn((len(data) - randSeqLen))
		seqRead := data[startIndex : startIndex+randSeqLen]
		// fmt.Println(string(seqRead))
		seqReads = append(seqReads, seqRead)
		seqReadStructs = append(seqReadStructs,
			Seq{
				StartPos:   startIndex,
				Length:     randSeqLen,
				BPSequence: string(seqRead),
			},
		)
	}

	// Write as CSV with dataframes
	df := dataframe.LoadStructs(seqReadStructs)
	fmt.Println(df)
	f, _ := os.Create("./dataGen/covidReads.csv")
	df.WriteCSV(f)
	// Write the reads to a file
	outputWrites, err := os.Create("./dataGen/covidReads.seqs")
	check(err)
	defer outputWrites.Close()
	for n := 0; n < nAligns; n++ {
		outputWrites.Write(seqReads[n])
		outputWrites.WriteString("\n")
	}
	outputWrites.Sync()
}
