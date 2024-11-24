package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/agnivade/levenshtein"
)

//go:embed input.txt
var input string
var RowsAbove int
var RowsLeft int

func CalculateBlock(block []string) (outcome int) {
	var row int

	for i, _ := range block {
		var totalDistance int
		row = 0

		if i+1 < len(block) {
			//fmt.Println("NEW ROW STARTING ========", i)
			row = i + 1

			// Walk backwards and forwards
			top := i
			bottom := i + 1
			for {

				//fmt.Println("Start loop")
				if top < 0 || bottom > len(block)-1 {
					//fmt.Println("break", top, bottom)
					break
				}
				//fmt.Println(block[top], block[bottom])
				distance := levenshtein.ComputeDistance(block[top], block[bottom])
				totalDistance += distance
				//if totalDistance > 1 {
				//	break
				//}
				top--
				bottom++
			}
		}
		//fmt.Println("totalDistance", totalDistance, row)
		if totalDistance == 1 {
			return row
		}
	}

	return row
}

func main() {
	linesInput := strings.Split(strings.TrimSpace(input), "\n\n")
	for _, block := range linesInput {
		blockSplit := strings.Split(block, "\n")
		rows := CalculateBlock(blockSplit)
		RowsAbove += rows
		if rows == 0 {
			newblockSplit := []string{}

			for j := 0; j < len(blockSplit[0]); j++ {
				newblockSplit = append(newblockSplit, "")

				newblockSplit[j] = ""
				for i := 0; i < len(blockSplit); i++ {
					newblockSplit[j] = newblockSplit[j] + string(blockSplit[i][j])
				}
			}
			blockSplit = newblockSplit
			rows = CalculateBlock(blockSplit)
			RowsLeft += rows
		}
		fmt.Println(rows)

	}
	fmt.Println(RowsLeft + 100*RowsAbove)

}
