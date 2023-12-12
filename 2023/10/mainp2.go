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
var Total int
var Lines []string
var LinesMax int
var CharactersMax int
var History []string

func findNext(LineIndex, CharacterIndex int) (NextLineIndex, NextCharacterIndex int) {
	MyCharacter := string(Lines[LineIndex][CharacterIndex])
	// Up/Down
	if LineIndex-1 >= 0 {
		NextPos := string(Lines[LineIndex-1][CharacterIndex])
		if MyCharacter == "|" || MyCharacter == "L" || MyCharacter == "J" || MyCharacter == "S" {
			if NextPos == "|" || NextPos == "7" || NextPos == "F" {
				Tuple := fmt.Sprint(LineIndex-1) + "," + fmt.Sprint(CharacterIndex)
				if !slices.Contains(History, Tuple) {
					return LineIndex - 1, CharacterIndex
				}
			}
		}
	}
	// Down
	if LineIndex+1 <= LinesMax-1 {
		NextPos := string(Lines[LineIndex+1][CharacterIndex])
		if MyCharacter == "|" || MyCharacter == "7" || MyCharacter == "F" || MyCharacter == "S" {
			if NextPos == "|" || NextPos == "L" || NextPos == "J" {
				Tuple := fmt.Sprint(LineIndex+1) + "," + fmt.Sprint(CharacterIndex)
				if !slices.Contains(History, Tuple) {

					return LineIndex + 1, CharacterIndex
				}
			}
		}
	}
	// Left
	if CharacterIndex-1 >= 0 {
		NextPos := string(Lines[LineIndex][CharacterIndex-1])
		if MyCharacter == "-" || MyCharacter == "J" || MyCharacter == "7" || MyCharacter == "S" {
			if NextPos == "-" || NextPos == "L" || NextPos == "F" {
				Tuple := fmt.Sprint(LineIndex) + "," + fmt.Sprint(CharacterIndex-1)
				if !slices.Contains(History, Tuple) {

					return LineIndex, CharacterIndex - 1
				}
			}
		}
	}
	// Right
	if CharacterIndex+1 <= CharactersMax {
		NextPos := string(Lines[LineIndex][CharacterIndex+1])
		if MyCharacter == "-" || MyCharacter == "F" || MyCharacter == "L" || MyCharacter == "S" {
			if NextPos == "-" || NextPos == "J" || NextPos == "7" {
				Tuple := fmt.Sprint(LineIndex) + "," + fmt.Sprint(CharacterIndex+1)
				if !slices.Contains(History, Tuple) {

					return LineIndex, CharacterIndex + 1
				}
			}
		}
	}
	return 0, 0
}

func main() {
	var StartlineIndex int
	var StartcharacterIndex int

	Lines = strings.Split(strings.TrimSpace(input), "\n")
	LinesMax = len(Lines)
	CharactersMax = len(Lines[0])

	for lineIndex, line := range Lines {
		for characterIndex, _ := range line {
			Character := string(Lines[lineIndex][characterIndex])
			if Character == "S" {
				StartlineIndex = lineIndex
				StartcharacterIndex = characterIndex
			}
		}
	}
	// Move one
	MoveLineIndex, MoveCharacterIndex := StartlineIndex, StartcharacterIndex

	for {
		Tuple := fmt.Sprint(MoveLineIndex) + "," + fmt.Sprint(MoveCharacterIndex)
		History = append(History, Tuple)
		//fmt.Println(Tuple)
		MoveLineIndex, MoveCharacterIndex = findNext(MoveLineIndex, MoveCharacterIndex)
		if MoveLineIndex == 0 && MoveCharacterIndex == 0 {
			break
		}
	}
	fmt.Println(History)

	for line := 0; line < LinesMax; line++ {
		Out := true
		LastBend := ""

		for character := 0; character < CharactersMax; character++ {

			LineS := strconv.Itoa(line)
			CharacterS := strconv.Itoa(character)

			Tuple := LineS + "," + string(CharacterS)
			// We have a character
			if slices.Contains(History, Tuple) {
				Character := string(Lines[line][character])
				if Character == "S" {
					Character = "|"
				}
				fmt.Printf(Character)

				if Character == "F" || Character == "L" {
					LastBend = Character
				}
				if Character == "7" && LastBend == "L" {
					Out = !Out
				}
				if Character == "J" && LastBend == "F" {
					Out = !Out
				}

				if Character == "|" {
					Out = !Out
				}

				// We don't have a character
			} else {
				if Out == false {
					fmt.Printf("I")
					Total++
				} else {
					fmt.Printf("O")
				}

			}
		}
		fmt.Println()
	}
	fmt.Println(Total)
}
