package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	// Part 1
	input := "iwrupvqb"

	for i := 1; true; i++ {
		appended_input := input + strconv.FormatInt(int64(i), 10)
		hash := md5.Sum([]byte(appended_input))
		hash_s := hex.EncodeToString(hash[:])
		beginning := hash_s[:5]
		if beginning == "00000" {
			fmt.Println("Part 1")
			fmt.Println(i)
			break
		}
	}
	for i := 1; true; i++ {
		appended_input := input + strconv.FormatInt(int64(i), 10)
		hash := md5.Sum([]byte(appended_input))
		hash_s := hex.EncodeToString(hash[:])
		beginning := hash_s[:6]
		if beginning == "000000" {
			fmt.Println("Part 2")
			fmt.Println(i)
			break
		}
	}
}
