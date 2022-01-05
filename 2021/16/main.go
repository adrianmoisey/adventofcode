package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

var m = map[string]string{
	"0": "0000", "1": "0001", "2": "0010", "3": "0011",
	"4": "0100", "5": "0101", "6": "0110", "7": "0111",
	"8": "1000", "9": "1001", "A": "1010", "B": "1011",
	"C": "1100", "D": "1101", "E": "1110", "F": "1111",
}

func hexToBinary(hex string) string {
	hx := s.Split(hex, "")
	bx := make([]string, len(hx))
	for i, h := range hx {
		bx[i] = m[h]
	}
	return s.Join(bx, "")
}

func splitBin(binary string, i int) (string, string) {
	bx := s.Split(binary, "")
	return s.Join(bx[0:i], ""), s.Join(bx[i:], "")
}

func binToInt(bin string) int64 {
	if len(bin) < 4 {
		bx := s.Split(bin, "")
		zeroPad := make([]string, 4)
		zeroPad[0] = "0"
		zeroPad[1] = bx[0]
		zeroPad[2] = bx[1]
		zeroPad[3] = bx[2]
		bin = s.Join(zeroPad, "")
	}
	num, _ := strconv.ParseInt(bin, 2, 64)
	return num
}

func readLiteral(binary string) (int64, string) {
	bx := s.Split(binary, "")
	bin := make([]string, 0)
	for bx[0] != "0" {
		bin = append(bin, bx[1:5]...)
		bx = bx[5:]
	}
	bin = append(bin, bx[1:5]...)
	binary = s.Join(bx[5:], "")
	return binToInt(s.Join(bin, "")), binary
}

func calc(values []int64, action int64) int64 {
	var sum int64 = 0
	switch action {
	case 0:
		for _, v := range values {
			sum += v
		}
		return sum
	case 1:
		sum = 1
		for _, v := range values {
			sum *= v
		}
		return sum
	case 2:
		sum = values[0]
		for _, v := range values {
			if v < sum {
				sum = v
			}
		}
		return sum
	case 3:
		for _, v := range values {
			if v > sum {
				sum = v
			}
		}
		return sum
	case 5:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case 6:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case 7:
		if values[0] == values[1] {
			return 1
		}
		return 0
	}
	return 0
}

func readPacket(binary string) (int64, int64, string) {
	// version count that we return
	var v int64
	var version int64 = 0
	var val int64

	fst, binary := splitBin(binary, 3)
	version += binToInt(fst)

	fst, binary = splitBin(binary, 3)
	operatorType := binToInt(fst)
	if operatorType == 4 {
		val, binary = readLiteral(binary)
		return version, val, binary
	} else {
		fst, binary = splitBin(binary, 1)

		values := make([]int64, 0)
		if fst == "0" {
			fst, binary = splitBin(binary, 15)
			bitLength := binToInt(fst)

			prevBinary := binary // to track diff
			v, val, binary = readPacket(binary)
			values = append(values, val)
			version += v

			diff := len(prevBinary) - len(binary)
			for diff < int(bitLength) {
				v, val, binary = readPacket(binary)
				values = append(values, val)
				version += v

				diff = len(prevBinary) - len(binary)
			}

			val = calc(values, operatorType)
		} else if fst == "1" {
			fst, binary = splitBin(binary, 11)

			count := int(binToInt(fst))
			for i := 0; i < count; i++ {
				v, val, binary = readPacket(binary)
				values = append(values, val)
				version += v
			}

			val = calc(values, operatorType)
		}
	}
	return version, val, binary
}

func main() {
	file, _ := os.ReadFile("input.txt")
	binary := hexToBinary(string(file))
	versionCount, value, _ := readPacket(binary)

	fmt.Println("Part 1:", versionCount)
	fmt.Println("Part 2:", value)
}
