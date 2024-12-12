package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	Part1(lines[0])
}

func Part1(input string) {
	var BlockID int

	DiskLayout := make([]int, 0)
	for i, c := range input {
		if i%2 == 0 {
			value, _ := strconv.Atoi(string(c))
			for i := 0; i < value; i++ {
				DiskLayout = append(DiskLayout, []int{BlockID}...)
			}
			BlockID++

		} else {
			value, _ := strconv.Atoi(string(c))
			for i := 0; i < value; i++ {
				DiskLayout = append(DiskLayout, []int{-1}...)
			}
		}
	}
	fmt.Println(DiskLayout)
	// Defrag
	for fileID := slices.Max(DiskLayout); fileID >= 0; fileID-- {
		fmt.Println("Total fileid", fileID)
		indexStart := slices.Index(DiskLayout, fileID)
		var indexOffset int
		for {
			if indexStart+indexOffset == len(DiskLayout) || DiskLayout[indexStart+indexOffset] != fileID {
				break
			} else {
				indexOffset++
			}
		}
		fmt.Println("Matched file, disk layout, indexStart, indexOffset", fileID, DiskLayout[indexStart:indexStart+indexOffset], indexStart, indexOffset)
		fmt.Println(DiskLayout)

		for positionInDiskLayout, characterAtPositionInDiskLayout := range DiskLayout {
			var found bool
			found = true
			if characterAtPositionInDiskLayout == -1 || characterAtPositionInDiskLayout == -2 {

				// We can move up indexOffset spots
				// fmt.Println("Start Loop")
				for k := indexOffset - 1; k > 0; k-- {
					// fmt.Println("LOOP", positionInDiskLayout+k, positionInDiskLayout, k, len(DiskLayout))
					if positionInDiskLayout+k == len(DiskLayout) {
						// fmt.Println("Printing possible", DiskLayout[positionInDiskLayout+k])
					}
					if positionInDiskLayout+k >= len(DiskLayout) || DiskLayout[positionInDiskLayout+k] > -1 {
						found = false
						break
					}
				}
			} else {
				found = false
			}

			// Found a slot! Printing it  it out
			if found {
				fmt.Println("Found a slot!")
				if indexStart < positionInDiskLayout {
					continue
				}
				for z := 0; z < indexOffset; z++ {
					// Set new position
					DiskLayout[indexStart+z] = -1

					fmt.Println("Occupting character", DiskLayout[positionInDiskLayout+z], fileID, positionInDiskLayout+z)
					DiskLayout[positionInDiskLayout+z] = fileID
					fmt.Println(DiskLayout)

				}
				fmt.Println(DiskLayout)

				break
			}
		}
		fmt.Println()

	}
	fmt.Println(DiskLayout)

	var total int
	for i, v := range DiskLayout {
		if v >= 0 {
			total += i * v
		}

	}
	// Too high 8358671598691
	// Too high 8280554925876
	//         15293743965418
	//         6239783302560
	fmt.Println("Total", total)

}

func pop(xs *[]int) int {
	x := (*xs)[len(*xs)-1]
	*xs = (*xs)[:len(*xs)-1]
	return x
}
