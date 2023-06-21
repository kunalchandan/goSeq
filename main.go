package main

import (
	"fmt"
	// "github.com/go-gota/gota/dataframe"
	// "github.com/jroimartin/gocui"
	// "golang.org/x/image/colornames"
)

func main() {
	sequence1 := []byte{Ad, Th, Cy, Gu}
	sequence2 := []byte{Ad, Th, Cy, Ad}

	scoringMatrix, directionMatrix := GenNWMatrix(sequence1, sequence2)
	drawMatricies(scoringMatrix, directionMatrix, sequence1, sequence2)

	fmt.Println("-------------------------")
	fmt.Println("-------------------------")
	fmt.Println("-------------------------")

	sequence1 = []byte{Gu, Cy, Ad, Th, Gu, Cy, Gu}
	sequence2 = []byte{Gu, Ad, Th, Th, Ad, Cy, Ad}

	scoringMatrix, directionMatrix = GenNWMatrix(sequence1, sequence2)
	drawMatricies(scoringMatrix, directionMatrix, sequence1, sequence2)

	fmt.Println("-------------------------")
	fmt.Println("-------------------------")
	fmt.Println("-------------------------")

	sequence1 = []byte{Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, Gu, }
	sequence2 = []byte{Gu, Gu, Gu, Gu, Gu, Ad, Ad, Ad, Ad, Ad, Ad, Gu, Gu, Gu, Gu, Gu, Gu, }

	scoringMatrix, directionMatrix = GenNWMatrix(sequence1, sequence2)
	drawMatricies(scoringMatrix, directionMatrix, sequence1, sequence2)
	// fmt.Println(NWScore("banana", "the banna man"))
	// generateData()
	// readData()
}
