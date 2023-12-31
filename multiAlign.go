// package multiAlign
package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"math"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	// . "github.com/samber/mo"
)

type BasePair byte

const (
	Adenine  BasePair = 97
	Thymine  BasePair = 116
	Cytosine BasePair = 99
	Guanine  BasePair = 103
)

const (
	Ad  byte = 97
	Th  byte = 116
	Cy  byte = 99
	Gu  byte = 103
	GAP byte = 61
)

func repr(b byte) string {
	if b == byte(Ad) {
		return "Ad"
	} else if b == byte(Th) {
		return "Th"
	} else if b == byte(Cy) {
		return "Cy"
	} else if b == byte(Gu) {
		return "Gu"
	} else {
		return "??"
	}
}

type Direction int

const (
	North    Direction = 1
	Diagonal Direction = 0
	West     Direction = -1
)

func SubsCost(base1 byte, base2 byte) int {
	// https://en.wikipedia.org/wiki/Needleman%E2%80%93Wunsch_algorithm#Scoring_systems
	if base1 == base2 {
		return 1
	} else {
		return -1
	}
}

func GenNWMatrix(seq1 []byte, seq2 []byte) ([][]int, [][]int) {
	gapPenalty := -2
	// Initialize scoring matrix with null value
	fmt.Println("Initializing scoring matrix")
	scoringMatrix := make([][]int, len(seq1)+1)
	directionMatrix := make([][]int, len(seq1)+1)
	for i := range scoringMatrix {
		scoringMatrix[i] = make([]int, len(seq2)+1)
		directionMatrix[i] = make([]int, len(seq2)+1)
	}
	// Initialize scoring matrix with null values
	for ii := 0; ii < len(seq1)+1; ii++ {
		for jj := 0; jj < len(seq2)+1; jj++ {
			scoringMatrix[ii][jj] = -1
			directionMatrix[ii][jj] = 0
		}
	}

	fmt.Println("Scoring...")
	for ii := 0; ii < len(seq1)+1; ii++ {
		for jj := 0; jj < len(seq2)+1; jj++ {
			scoringMatrix[ii][jj] = 0
			CDiag := math.MinInt
			CWest := math.MinInt
			CNorth := math.MinInt
			Dir := Diagonal
			CMax := CDiag

			if ii == 0 && jj == 0 {
				CDiag = 0
				CWest = 0
				CNorth = 0
			} else if ii == 0 {
				CWest = scoringMatrix[ii][jj-1] + gapPenalty
			} else if jj == 0 {
				CNorth = scoringMatrix[ii-1][jj] + gapPenalty
			} else {
				CDiag = scoringMatrix[ii-1][jj-1] + SubsCost(seq1[ii-1], seq2[jj-1])
				CWest = scoringMatrix[ii][jj-1] + gapPenalty
				CNorth = scoringMatrix[ii-1][jj] + gapPenalty
			}

			CMax = CDiag
			if CNorth > CMax {
				CMax = CNorth
				Dir = North
			}
			if CWest > CMax {
				CMax = CWest
				Dir = West
			}
			scoringMatrix[ii][jj] = CMax
			directionMatrix[ii][jj] = int(Dir)
		}
	}
	return scoringMatrix, directionMatrix
}

func drawMatricies(scoringMatrix [][]int, directionMatrix [][]int, seq1 []byte, seq2 []byte) {
	fmt.Println("Drawing Score Matrix")
	fmt.Print("       ")
	for i := 0; i < len(seq2); i++ {
		fmt.Printf("  %s", repr(seq2[i]))
	}
	fmt.Println("")
	for i := 0; i < len(scoringMatrix); i++ {
		if i == 0 {
			fmt.Printf("   ")
		} else {
			fmt.Printf("%s ", repr(seq1[i-1]))
		}
		for j := 0; j < len(scoringMatrix[i]); j++ {
			fmt.Printf("\033[31;1;4m%4d\033[0m", scoringMatrix[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("Drawing Direction Matrix")
	fmt.Print("      ")
	for i := 0; i < len(seq2); i++ {
		fmt.Printf("%s ", repr(seq2[i]))
	}
	fmt.Println("")
	for i := 0; i < len(directionMatrix); i++ {
		if i == 0 {
			fmt.Printf("   ")
		} else {
			fmt.Printf("%s ", repr(seq1[i-1]))
		}
		for j := 0; j < len(directionMatrix[i]); j++ {
			if Direction(directionMatrix[i][j]) == Diagonal {
				fmt.Printf("\033[31;1;4m %s \033[0m", "\\")
			} else if Direction(directionMatrix[i][j]) == West {
				fmt.Printf("\033[31;1;4m %s \033[0m", "-")
			} else if Direction(directionMatrix[i][j]) == North {
				fmt.Printf("\033[31;1;4m %s \033[0m", "|")
			}
		}
		fmt.Println("")
	}
}

func alignedSeqsFromMatricies(scoringMatrix [][]int, directionMatrix [][]int, seq1 []byte, seq2 []byte) ([]byte, []byte, int, int) {
	alignedLen := len(seq1)
	if len(seq2) > alignedLen {
		alignedLen = len(seq2)
	}
	index1 := len(seq1) - 1
	index2 := len(seq2) - 1
	reversedSeq1 := make([]byte, alignedLen)
	reversedSeq2 := make([]byte, alignedLen)
	alignmentScore1 := 0
	alignmentScore2 := 0
	for i := 0; i < alignedLen; i++ {
		reversedSeq1[i] = 0
		reversedSeq2[i] = 0
		alignmentScore1 += scoringMatrix[index1+1][index2+1]
		alignmentScore2 += scoringMatrix[index1+1][index2+1]
		if Direction(directionMatrix[index1+1][index2+1]) == Diagonal {
			// Implies alignment of sequences here
			reversedSeq1 = append(reversedSeq1, seq1[index1])
			reversedSeq2 = append(reversedSeq2, seq2[index2])
			index1--
			index2--
		} else if Direction(directionMatrix[index1+1][index2+1]) == North {
			// Implies gap in the top sequence
			fmt.Println("Gap in top sequence")
			reversedSeq1 = append(reversedSeq1, GAP)
			reversedSeq2 = append(reversedSeq2, seq2[index2])
			index1--
		} else if Direction(directionMatrix[index1+1][index2+1]) == West {
			// Implies gap in the left sequence
			fmt.Println("Gap in left sequence")
			reversedSeq1 = append(reversedSeq1, seq1[index1])
			reversedSeq2 = append(reversedSeq2, GAP)
			index2--
		}
	}
	// Reverse the reversed sequences
	for i, j := 0, len(reversedSeq1)-1; i < j; i, j = i+1, j-1 {
		reversedSeq1[i], reversedSeq1[j] = reversedSeq1[j], reversedSeq1[i]
		reversedSeq2[i], reversedSeq2[j] = reversedSeq2[j], reversedSeq2[i]
	}
	fmt.Println("Sequences 1 and 2")
	fmt.Println(string(reversedSeq1))
	fmt.Println(string(reversedSeq2))
	fmt.Println("Alignment Scores")
	fmt.Println(alignmentScore1)
	fmt.Println(alignmentScore2)
	return reversedSeq1, reversedSeq2, alignmentScore1, alignmentScore2
}

func AlignmentScore(seq1 []byte, seq2 []byte) int {
	if len(seq1) != len(seq2) {
		fmt.Printf("Sequences 1 and 2 are of different lengths\n")
		fmt.Printf("Length 1: %d\n", len(seq1))
		fmt.Printf("Length 2: %d\n", len(seq2))
		fmt.Printf("Lengths must be equal, try forced compare where ends are padded with GAPs")
		panic("Sequences must be of equal length")
	}
	alignedLen := len(seq1)
	alignmentScore := 0
	sequentialGap := false
	for i := 0; i < alignedLen; i++ {
		if (seq1[i] == GAP || seq2[i] == GAP) && !sequentialGap {
			alignmentScore += -5
			sequentialGap = true
		} else if (seq1[i] == GAP || seq2[i] == GAP) && sequentialGap {
			alignmentScore += -1
		} else if seq1[i] != seq2[i] {
			alignmentScore += -1
			sequentialGap = false
		} else {
			alignmentScore += 1
			sequentialGap = false
		}

	}
	return alignmentScore
}

func AlignmentScoreForced(seq1 []byte, seq2 []byte) int {
	// warn that this function should not be used most of the time
	paddingSeq := len(seq2) - len(seq1)
	if paddingSeq == 0 {
		return AlignmentScore(seq1, seq2)
	} else if paddingSeq > 0 {
		return AlignmentScore(append(seq1, make([]byte, paddingSeq)...), seq2)
	} else {
		return AlignmentScore(seq1, append(seq2, make([]byte, -paddingSeq)...))
	}
}

func readData() dataframe.DataFrame {
	content, _ := ioutil.ReadFile("./dataGen/covidReads.csv")
	ioContent := strings.NewReader(string(content))

	dataRaw := dataframe.ReadCSV(
		ioContent,
	)
	return dataRaw
}

func generateHashLookup(rawSeqs dataframe.DataFrame) dataframe.DataFrame {
	// Note hash length could be increased to reduce likelihood of collisions
	// IMO n/2^64 is enough
	// hasher := fnv.New64()
	hashLambda := func(s series.Series) series.Series {
		// Get the sequence
		hasher := fnv.New64a()
		str := []byte(s.Elem(2).String())
		hasher.Write(str)
		hashInt := hasher.Sum64()
		hash := fmt.Sprintf("%016X", hashInt)

		// new_s := s.Copy()
		// new_s.Append(hash)
		// return new_s
		return series.New(hash, series.String, "BPSequenceHash")
	}
	ndf := dataframe.New(
		rawSeqs.Col("BPSequence"),
		rawSeqs.Rapply(hashLambda).Col("X0"),
	)
	ndf = ndf.Rename("BPSequenceHash", "X0")
	return ndf
}

func pairwiseAlign(df dataframe.DataFrame) dataframe.DataFrame {
	// Generate pairwise scores for alignment
	sequences := df.Col("BPSequence")
	for ii := 0; ii < df.Nrow(); ii++ {
		seq1 := []byte(sequences.Elem(ii).String())
		for jj := ii + 1; jj < df.Nrow(); jj++ {
			seq2 := []byte(sequences.Elem(ii).String())
			_ = AlignmentScoreForced(
				seq1,
				seq2,
			)
		}
	}
	return df
}
