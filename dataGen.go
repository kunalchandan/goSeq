// package dataGen
package main

import (
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

func generateData(sourceFile string, outputFilePath string, nAligns int, seqLenMax int, seqLenMin int) {
	// Seed for reproducibility
	rand.Seed(0)
	data, err := os.ReadFile(sourceFile)
	check(err)

	// maybe use a dataframe/like object and append the indices and lengths as well
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
	// This CSV is the entire dataframe with the start index and length
	df := dataframe.LoadStructs(seqReadStructs)
	f, _ := os.Create(outputFilePath + ".csv")
	df.WriteCSV(f)
	// Write the reads to a file
	// This file is the raw reads themselves as if you didn't know where the sequences actually came from
	outputWrites, err := os.Create(outputFilePath + ".seqs")
	check(err)
	defer outputWrites.Close()
	for n := 0; n < nAligns; n++ {
		outputWrites.Write(seqReads[n])
		outputWrites.WriteString("\n")
	}
	outputWrites.Sync()
}
