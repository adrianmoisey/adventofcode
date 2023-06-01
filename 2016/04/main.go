package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Each room consists of an encrypted name (lowercase letters separated by dashes)
// followed by a dash, a sector ID, and a checksum in square brackets.
type Room struct {
	Name           string
	SectorID       int64
	Checksum       string
	ChecksumLength int
}

func main() {
	// Open file by line
	readFile, _ := os.Open("input.txt")
	//readFile, _ := os.Open("sample.txt")
	//readFile, _ := os.Open("sample1.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var total int64

	for fileScanner.Scan() {
		fmt.Println("=====")
		line := fileScanner.Text()
		room := StringToRoom(line)
		//fmt.Println(room.Name, room.SectorID, room.Checksum)

		mapping := make(map[string]int)
		for _, l := range room.Name {
			//fmt.Printf("%c\n", l)
			if string(l) != "-" {
				mapping[string(l)] += 1
			}
		}
		reverse_mapping := make(map[int][]string)
		keys := make([]int, 0, len(reverse_mapping))
		for k, v := range mapping {
			reverse_mapping[v] = append(reverse_mapping[v], k)
		}
		for k := range reverse_mapping {
			keys = append(keys, k)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))

		our_compare := ""

		for _, k := range keys {
			a := reverse_mapping[k]
			sort.Strings(a)
			//fmt.Println(k, a)

			for _, l := range a {
				our_compare = our_compare + l
			}
		}
		our_compare = our_compare[:room.ChecksumLength]
		if our_compare == room.Checksum {
			// Part 1
			total = total + room.SectorID
			// Part 2
			//fmt.Println(room.Name)

			rotatedName := ""

			for _, c := range room.Name {
				//fmt.Println(c)
				if c == 45 {
					rotatedName = rotatedName + " "
					//fmt.Println(" ")
				} else {
					rotatedName = rotatedName + string(caesar(c, int(room.SectorID)))
					//fmt.Println(string(caesar(c, int(room.SectorID))))
				}
			}
			fmt.Println(rotatedName, room.SectorID)
		}
	}

	readFile.Close()
	fmt.Println(total)
}

func StringToRoom(in string) (room Room) {
	lastIndex := strings.LastIndex(in, "-")

	room.Name = in[:lastIndex]

	remaining := in[lastIndex+1:]
	a := strings.FieldsFunc(remaining, Split)

	room.SectorID, _ = strconv.ParseInt(a[0], 0, 64)
	room.Checksum = a[1]
	room.ChecksumLength = len(room.Checksum)
	return
}

func Split(r rune) bool {
	return r == '[' || r == ']'
}

func caesar(r rune, shift int) rune {
	// Shift character by specified number of places.
	// ... If beyond range, shift backward or forward.
	shift = shift % 26
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}
