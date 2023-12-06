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
			[]int{7, 9},
			[]int{15, 40},
			[]int{30, 200},
		}
	*/
	races := [][]int{
		[]int{46, 358},
		[]int{68, 1054},
		[]int{98, 1807},
		[]int{66, 1080},
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

			fmt.Println(ButtonHold, CarMoved)
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
