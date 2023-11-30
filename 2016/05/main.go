package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.Trim(input, "\n")

	var answer []string
	part2 := make([]string, 8)
	var part2_count int

	// 3231929
	for i := 0; i <= 323193000; i++ {
		// var position string
		var letter string
		h := md5.New()
		next := input + strconv.FormatInt(int64(i), 10)
		//fmt.Println(next)
		io.WriteString(h, next)
		hash := hex.EncodeToString(h.Sum(nil))
		if hash[:5] == "00000" {
			// fmt.Println(hash)
			answer = append(answer, hash[5:6])
			// position = hash[5:6]
			position_i, err := strconv.Atoi(hash[5:6])
			if err == nil {
				letter = hash[6:7]
				if position_i < len(part2) && part2[position_i] == "" {
					part2[position_i] = letter
					// fmt.Println(position, position_i, letter)
					part2_count += 1
				}
			}

		}
		if len(answer) >= 8 && part2_count >= 8 {
			fmt.Println("Part1: ")
			fmt.Println(strings.Join(answer[:8], ""))
			fmt.Println("Part2: ")
			fmt.Println(strings.Join(part2, ""))
			break
		}
	}

}
