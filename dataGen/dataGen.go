package main

import (
	"fmt"
	"math/rand"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	data, err := os.ReadFile("./covid.seq")
	check(err)

	// Create the reads
	nAligns := 10
	seqLenMax := 12
	seqLenMin := 8
	// maybe use a dataframe/like object and append the indicies and lengths as well
	seqReads := [][]byte{}
	for n := 0; n < nAligns; n++ {
		randSeqLen := rand.Intn(seqLenMax-seqLenMin) + seqLenMin
		startIndex := rand.Intn((len(data) - randSeqLen))
		seqRead := data[startIndex : startIndex+randSeqLen]
		fmt.Println(string(seqRead))
		seqReads = append(seqReads, seqRead)
	}
	fmt.Println("")

	// Write the reads to a file
	outputWrites, err := os.Create("./covidReads.seqs")
	check(err)
	defer outputWrites.Close()
	for n := 0; n < nAligns; n++ {
		outputWrites.Write(seqReads[n])
		outputWrites.WriteString("\n")
	}
	outputWrites.Sync()
}
