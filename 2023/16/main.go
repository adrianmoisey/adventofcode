package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/gammazero/deque"
)

//go:embed input.txt
var input string
var VisitedPoints map[string]bool
var StartBeams []LightBeam

type LightBeam struct {
	Line      int
	Character int
	Direction string
}

func (l *LightBeam) CurrentLocation() (location string) {
	LineString := strconv.Itoa(l.Line)
	CharacterString := strconv.Itoa(l.Character)
	return LineString + "," + CharacterString
}

func main() {

	linesInput := strings.Split(strings.TrimSpace(input), "\n")

	for i := 0; i < len(linesInput[0]); i++ {
		StartBeam := LightBeam{
			Line:      0,
			Character: i,
			Direction: "S",
		}
		StartBeam2 := LightBeam{
			Line:      len(linesInput) - 1,
			Character: i,
			Direction: "N",
		}

		StartBeams = append(StartBeams, StartBeam)
		StartBeams = append(StartBeams, StartBeam2)
	}

	for i := 0; i < len(linesInput); i++ {
		StartBeam := LightBeam{
			Line:      i,
			Character: 0,
			Direction: "E",
		}
		StartBeam2 := LightBeam{
			Line:      i,
			Character: len(linesInput[0]) - 1,
			Direction: "W",
		}

		StartBeams = append(StartBeams, StartBeam)
		StartBeams = append(StartBeams, StartBeam2)

	}

	/*StartBeam2 := LightBeam{
		Line:      0,
		Character: 3,
		Direction: "S",
	}
	StartBeams = append(StartBeams, StartBeam2)
	*/

	Max := 0
	for _, StartBeam := range StartBeams {
		fmt.Println("Start Beam", StartBeam)
		Deque := deque.New[LightBeam]()
		VisitedPoints := make(map[string]bool)

		//Max = 0

		Deque.PushBack(StartBeam)
		/*
			for _, block := range linesInput {
				fmt.Println(block)
			}
		*/
		history := 0
		counter := 0
		for {
			if len(VisitedPoints) > history {
				history = len(VisitedPoints)
				counter = 0
			} else {
				counter++
			}
			if counter > 500000 {
				if Max < len(VisitedPoints) {
					Max = len(VisitedPoints)
				}
				break
			}
			/*
				fmt.Println()
				for x := range linesInput {
					for y := range linesInput[x] {

						LineString := strconv.Itoa(x)
						CharacterString := strconv.Itoa(y)
						Mapping := LineString + "," + CharacterString
						if VisitedPoints[Mapping] {
							fmt.Printf("#")
						} else {
							fmt.Printf(string(linesInput[x][y]))
						}

					}
					fmt.Println()
				}
			*/

			//fmt.Println()
			//fmt.Println("\033[2J")

			if Deque.Len() == 0 {
				break
			}
			CurrentBeam := Deque.PopFront()
			CurrentSquare := string(linesInput[CurrentBeam.Line][CurrentBeam.Character])

			//if VisitedPoints[CurrentBeam.CurrentLocation()] == true && CurrentSquare == "." {
			//	break
			//}

			//CurrentBeam, _ := CurrentBeamI.(LightBeam)
			VisitedPoints[CurrentBeam.CurrentLocation()] = true

			if (CurrentBeam.Direction == "E" || CurrentBeam.Direction == "W") && CurrentSquare == "|" {
				CurrentBeam.Direction = "S"
				NewBeam := LightBeam{
					Line:      CurrentBeam.Line,
					Character: CurrentBeam.Character,
					Direction: "N",
				}
				MoveBeam(NewBeam, Deque, len(linesInput[0]), len(linesInput))
			}

			if (CurrentBeam.Direction == "N" || CurrentBeam.Direction == "S") && CurrentSquare == "-" {
				CurrentBeam.Direction = "E"
				NewBeam := LightBeam{
					Line:      CurrentBeam.Line,
					Character: CurrentBeam.Character,
					Direction: "W",
				}
				MoveBeam(NewBeam, Deque, len(linesInput[0]), len(linesInput))
			}

			//fmt.Println(CurrentBeam.Direction, CurrentSquare, Deque.Len(), len(VisitedPoints))
			if CurrentBeam.Direction == "N" {
				if CurrentSquare == "\\" {
					CurrentBeam.Direction = "W"
				} else if CurrentSquare == "/" {
					CurrentBeam.Direction = "E"
				}
			} else if CurrentBeam.Direction == "S" {
				if CurrentSquare == "\\" {
					CurrentBeam.Direction = "E"
				} else if CurrentSquare == "/" {
					CurrentBeam.Direction = "W"
				}
			} else if CurrentBeam.Direction == "E" {
				if CurrentSquare == "\\" {
					CurrentBeam.Direction = "S"
				} else if CurrentSquare == "/" {
					CurrentBeam.Direction = "N"
				}
			} else if CurrentBeam.Direction == "W" {
				if CurrentSquare == "\\" {
					CurrentBeam.Direction = "N"
				} else if CurrentSquare == "/" {
					CurrentBeam.Direction = "S"
				}
			}

			MoveBeam(CurrentBeam, Deque, len(linesInput[0]), len(linesInput))

		}
		//fmt.Println(VisitedPoints)
		/*
			for x := range linesInput {
				for y := range linesInput[x] {

					LineString := strconv.Itoa(x)
					CharacterString := strconv.Itoa(y)
					Mapping := LineString + "," + CharacterString
					if VisitedPoints[Mapping] {
						fmt.Printf("#")
					} else {
						fmt.Printf(string(linesInput[x][y]))
					}

				}

			}
		*/
		fmt.Println()

		fmt.Println(len(VisitedPoints))
	}
	fmt.Println("Max", Max)
}

func MoveBeam(lightBeam LightBeam, stack *deque.Deque[LightBeam], width, length int) {
	if lightBeam.Direction == "E" {
		lightBeam.Character++
		if lightBeam.Character < width {
			stack.PushBack(lightBeam)
		}
	}
	if lightBeam.Direction == "W" {
		lightBeam.Character--
		if lightBeam.Character >= 0 {
			stack.PushBack(lightBeam)
		}
	}

	if lightBeam.Direction == "N" {
		lightBeam.Line--
		if lightBeam.Line >= 0 {
			stack.PushBack(lightBeam)
		}
	}

	if lightBeam.Direction == "S" {
		lightBeam.Line++
		if lightBeam.Line < length {
			stack.PushBack(lightBeam)
		}
	}
}
