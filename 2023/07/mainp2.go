package main

import (
	_ "embed"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

//go:embed input.txt
var input string

var (
	total        int
	Ranking      map[int][]Card
	RankingSlice []Card
)

type Card struct {
	Hand         string
	AlphaHand    string
	Value        int
	CardMappings map[string]int
	CardType     int
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
// A, B, C, D, E, F, G, H, I, J, K, L, or M

// 7 - Five of a kind, where all five cards have the same label: AAAAA
// 6 - Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// 5 - Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// 4 - Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// 3 - Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// 2 - One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// 1 - High card, where all cards' labels are distinct: 23456

func (c *Card) PopulateCardType() {

	mapping := map[string]string{
		"A": "A",
		"K": "B",
		"Q": "C",
		"T": "D",
		"9": "E",
		"8": "F",
		"7": "G",
		"6": "H",
		"5": "I",
		"4": "J",
		"3": "K",
		"2": "L",
		"J": "M",
	}

	c.CardMappings = make(map[string]int)

	var AlpheHandSlice []string

	J := 0

	for _, character := range c.Hand {
		if string(character) == "J" {
			J++
		} else {
			c.CardMappings[string(character)]++
		}
		AlpheHandSlice = append(AlpheHandSlice, mapping[string(character)])
	}

	c.AlphaHand = strings.Join(AlpheHandSlice, "")

	a := maps.Values(c.CardMappings)
	slices.Sort(a)
	if len(a) != 0 {
		a[len(a)-1] = a[len(a)-1] + J

		if len(c.CardMappings) == 1 {
			c.CardType = 1
		} else if len(c.CardMappings) == 2 && slices.Contains(a, 1) && slices.Contains(a, 4) {
			c.CardType = 2
		} else if len(c.CardMappings) == 2 && slices.Contains(a, 2) && slices.Contains(a, 3) {
			c.CardType = 3
		} else if len(c.CardMappings) == 3 && slices.Contains(a, 3) && slices.Contains(a, 1) {
			c.CardType = 4
		} else if len(c.CardMappings) == 3 && slices.Contains(a, 2) && slices.Contains(a, 1) {
			c.CardType = 5
		} else if len(c.CardMappings) == 4 && slices.Contains(a, 2) && slices.Contains(a, 1) {
			c.CardType = 6
		} else if len(c.CardMappings) == 5 {
			c.CardType = 7
		}
	} else {
		c.CardType = 1
		AlpheHandSlice = []string{"M", "M", "M", "M", "M"}
		c.AlphaHand = strings.Join(AlpheHandSlice, "")
	}

}

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	Ranking = make(map[int][]Card)

	for _, line := range lines {
		gap := strings.Index(line, " ")
		cards := line[:gap]
		value := line[gap+1:]
		valueInt, _ := strconv.Atoi(value)

		card := Card{Hand: cards, Value: valueInt}
		card.PopulateCardType()
		Ranking[card.CardType] = append(Ranking[card.CardType], card)

	}

	for i := 7; i > 0; i-- {

		ranking := Ranking[i]
		sort.Slice(ranking, func(i, j int) bool {
			return ranking[i].AlphaHand < ranking[j].AlphaHand
		})
		slices.Reverse(ranking)
		for _, card := range ranking {
			RankingSlice = append(RankingSlice, card)
		}
	}
	for i, card := range RankingSlice {
		i++
		total += i * card.Value
		if card.CardType == 1 {
			fmt.Println(card)
		}
	}
	fmt.Println(total)
	// Too low: 250212379
	// Too low: 250212695
}
