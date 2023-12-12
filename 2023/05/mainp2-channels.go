package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

//go:embed input.txt
var input string

var (
	locations []int
	Seeds     []string
	Mappings  = make(map[string]Mapping)
)

type Row struct {
	Destination int
	Source      int
	Range       int
}

type Mapping struct {
	Type string
	Rows []Row
}

func (m *Mapping) FindPosition(number int) (position int) {
	position = number
	for _, row := range m.Rows {
		if number >= row.Source && number <= row.Source+row.Range {
			max := number - row.Source
			index := row.Destination + max

			return index
		}
	}
	return position
}

func main() {
	lines := strings.Split(input, "\n")
	var NewMapping Mapping
	for _, line := range lines {
		if line == "" {
			Mappings[NewMapping.Type] = NewMapping
			continue
		}
		// Populate Seeds
		if strings.Contains(line, "seeds") {
			seeds := strings.SplitN(line, " ", 2)[1]
			Seeds = strings.Fields(seeds)

			result := make([]string, 0, len(Seeds)/2)
			for i := 1; i < len(Seeds); i += 2 {
				result = append(result, Seeds[i-1]+","+Seeds[i])
			}

			Seeds = result

			continue
		} else if strings.Contains(line, "map") {
			mapName := strings.Split(line, " ")[0]
			NewMapping = Mapping{Type: mapName}
			continue
		}
		mapKey := strings.Split(line, " ")

		DestinationStart, _ := strconv.Atoi(mapKey[0])
		SourceStart, _ := strconv.Atoi(mapKey[1])
		RangeLength, _ := strconv.Atoi(mapKey[2])

		row := Row{}

		row.Destination = DestinationStart
		row.Source = SourceStart
		row.Range = RangeLength
		NewMapping.Rows = append(NewMapping.Rows, row)
	}

	var smol int
	smol = math.MaxInt16

	for _, seedTuple := range Seeds {
		seedStart := strings.Split(seedTuple, ",")[0]
		seedStartInt, _ := strconv.Atoi(seedStart)
		seedRange := strings.Split(seedTuple, ",")[1]
		seedRangeInt, _ := strconv.Atoi(seedRange)

		resultChan := make(chan int, 100000)
		doneChan := make(chan struct{})

		var wg sync.WaitGroup

		go func() {
			for result := range resultChan {
				if result < smol {
					smol = result
				}
			}
		}()

		for seed := seedStartInt; seed < seedStartInt+seedRangeInt+1; seed += 1 {
			fmt.Println("h")
			wg.Add(1)

			go FindLocation(seed, resultChan, &wg)
		}

		go func() {
			wg.Wait()
			close(resultChan)
			close(doneChan)
		}()

		<-doneChan

	}

	fmt.Println(smol)
}

func FindLocation(seed int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	m := Mappings["seed-to-soil"]
	soil := m.FindPosition(seed)
	m = Mappings["soil-to-fertilizer"]
	fertilizer := m.FindPosition(soil)

	m = Mappings["fertilizer-to-water"]
	water := m.FindPosition(fertilizer)

	m = Mappings["water-to-light"]
	light := m.FindPosition(water)

	m = Mappings["light-to-temperature"]
	temperature := m.FindPosition(light)

	m = Mappings["temperature-to-humidity"]
	humidity := m.FindPosition(temperature)

	m = Mappings["humidity-to-location"]
	location := m.FindPosition(humidity)

	resultChan <- location
}
