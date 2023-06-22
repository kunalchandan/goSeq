package main

import (
	"fmt"
	"testing"
)

func TestDrawMatricies_LargeMismatched(t *testing.T) {
	sequence1 := []byte{Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu}
	sequence2 := []byte{Gu, Gu, Gu, Gu, Gu, Ad, Ad, Ad, Ad, Ad, Ad, Gu, Gu, Gu, Gu, Gu, Gu}

	scoringMatrix, directionMatrix := GenNWMatrix(sequence1, sequence2)
	drawMatricies(scoringMatrix, directionMatrix, sequence1, sequence2)
	alignSeq1, alignSeq2, score1, score2 := alignedSeqsFromMatricies(scoringMatrix, directionMatrix, sequence1, sequence2)

	fmt.Println(string(alignSeq1), string(alignSeq2), score1, score2)
}

func TestDrawMatricies_Medium(t *testing.T) {
	sequence1 := []byte{Gu, Cy, Ad, Th, Gu, Cy, Gu}
	sequence2 := []byte{Gu, Ad, Th, Th, Ad, Cy, Ad}

	scoringMatrix, directionMatrix := GenNWMatrix(sequence1, sequence2)
	drawMatricies(scoringMatrix, directionMatrix, sequence1, sequence2)
	alignSeq1, alignSeq2, score1, score2 := alignedSeqsFromMatricies(scoringMatrix, directionMatrix, sequence1, sequence2)

	fmt.Println(string(alignSeq1), string(alignSeq2), score1, score2)
}

func TestDrawMatricies_Small(t *testing.T) {
	sequence1 := []byte{Ad, Th, Cy, Gu}
	sequence2 := []byte{Ad, Th, Cy, Ad}

	scoringMatrix, directionMatrix := GenNWMatrix(sequence1, sequence2)
	drawMatricies(scoringMatrix, directionMatrix, sequence1, sequence2)
	alignSeq1, alignSeq2, score1, score2 := alignedSeqsFromMatricies(scoringMatrix, directionMatrix, sequence1, sequence2)

	fmt.Println(string(alignSeq1), string(alignSeq2), score1, score2)
}
