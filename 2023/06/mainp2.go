package main

import (
	_ "embed"
	"fmt"
)

//go:embed sample.txt
var input string

var (
	total []int
)

func main() {
	/*
		races := [][]int{
			[]int{71530, 94200},
		}
	*/
	races := [][]int{
		[]int{46689866, 358105418071080},
	}

	for _, race := range races {
		RaceTime := race[0]
		RaceDistance := race[1]

		LongerThanDistance := 0

		fmt.Println(RaceTime, RaceDistance)
		i := 0
		Below := true
		Above := false
		for true {
			i++
			ButtonHold := i
			CarMoved := (RaceTime - i) * ButtonHold
			if CarMoved > RaceDistance {
				LongerThanDistance++
			}

			//fmt.Println(ButtonHold, CarMoved)
			if Below == true && Above == false && CarMoved > RaceDistance {
				Above = true
				Below = false
			} else if Below == false && Above == true && CarMoved < RaceDistance {
				break

			}
		}
		total = append(total, LongerThanDistance)
	}
	totalMultiplied := 1
	for _, v := range total {
		totalMultiplied = totalMultiplied * v
	}
	fmt.Println("Total: ", totalMultiplied)

}
