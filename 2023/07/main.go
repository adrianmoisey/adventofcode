package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

//go:embed sample.txt
var input string

var (
	total []int
	Cards []Card
)

type Card struct {
	Hand  string
	Value int
}

// Five of a kind, where all five cards have the same label: AAAAA
// Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// High card, where all cards' labels are distinct: 23456

func (c *Card) CardType() {
	required := mapset.NewSet[string]()

	for _, character := range c.Hand {
		required.Add(string(character))
	}
	fmt.Println(required)

}

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		gap := strings.Index(line, " ")
		cards := line[:gap]
		value := line[gap+1:]
		valueInt, _ := strconv.Atoi(value)

		card := Card{Hand: cards, Value: valueInt}
		Cards = append(Cards, card)

	}

	fmt.Println(Cards)
	Card1 := Cards[0]
	Card1.CardType()

}
