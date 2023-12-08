package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d7p1)
	root.AddCommand(d7p2)
}

var d7p1 = &cobra.Command{
	Use: "d7p1",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("score: %d", doD7P1(loadLines(7)))
	},
}

var d7p2 = &cobra.Command{
	Use: "d7p2",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("score: %d", doD7P2(loadLines(7)))
	},
}

type d7HandAndBid struct {
	hand      string
	bid       int
	typeScore int
}

func doD7P1(input []string) int {
	hands := lo.Map(input, func(s string, _ int) d7HandAndBid {
		split := sp(s, " ")
		reqLen(split, 2)
		hand := split[0]
		bid := atoi(split[1])
		typeScore := d7GetHandType(hand)

		return d7HandAndBid{
			hand:      hand,
			bid:       bid,
			typeScore: typeScore,
		}
	})

	// sort the hands by strength
	sort.Slice(hands, func(i, j int) bool {
		left, right := hands[i], hands[j]
		if left.typeScore != right.typeScore {
			return left.typeScore < right.typeScore
		}

		return d7HighCardSmaller1(left.hand, right.hand)
	})

	var tot int

	for i, hand := range hands {
		rank := i + 1
		score := hand.bid * rank

		tot += score
	}

	return tot
}

func doD7P2(input []string) int {
	hands := lo.Map(input, func(s string, _ int) d7HandAndBid {
		split := sp(s, " ")
		reqLen(split, 2)
		hand := split[0]
		bid := atoi(split[1])
		typeScore := d7GetHandType2(hand)

		return d7HandAndBid{
			hand:      hand,
			bid:       bid,
			typeScore: typeScore,
		}
	})

	// sort the hands by strength
	sort.Slice(hands, func(i, j int) bool {
		left, right := hands[i], hands[j]
		if left.typeScore != right.typeScore {
			return left.typeScore < right.typeScore
		}

		return d7HighCardSmaller2(left.hand, right.hand)
	})

	var tot int

	for i, hand := range hands {
		rank := i + 1
		score := hand.bid * rank

		tot += score
	}

	return tot
}

func d7GetHandType(hand string) int {
	cards := sp(hand, "")
	reqLen(cards, 5)

	cardMap := make(map[string]int)
	for _, card := range cards {
		cardMap[card]++
	}

	// if all 5 are the same, 5 of a kind
	if len(cardMap) == 1 {
		return 0
	}

	if len(cardMap) == 2 {
		// either 4 of a kind, or full house
		for _, v := range cardMap {
			if v == 4 {
				return -1
			}
		}

		return -2
	}

	if len(cardMap) == 3 {
		// either 3 of a kind, or 2 pair
		for _, v := range cardMap {
			if v == 3 {
				return -3
			}
		}

		return -4
	}

	if len(cardMap) == 4 {
		return -5
	}

	return -6
}

func d7HighCardSmaller1(hand1, hand2 string) bool {
	cards1 := sp(hand1, "")
	reqLen(cards1, 5)

	cards2 := sp(hand2, "")
	reqLen(cards2, 5)

	for i := 0; i < 5; i++ {
		if cards1[i] == cards2[i] {
			continue
		}

		return d7CardScore1(cards1[i]) < d7CardScore1(cards2[i])
	}

	return false
}

func d7HighCardSmaller2(hand1, hand2 string) bool {
	cards1 := sp(hand1, "")
	reqLen(cards1, 5)

	cards2 := sp(hand2, "")
	reqLen(cards2, 5)

	for i := 0; i < 5; i++ {
		if cards1[i] == cards2[i] {
			continue
		}

		return d7CardScore2(cards1[i]) < d7CardScore2(cards2[i])
	}

	return false
}

func d7CardScore1(c string) int {
	// if we are a number, take the number
	n, err := strconv.Atoi(c)
	if err == nil {
		return n
	}

	switch c {
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	}

	panic("unknown card")
}

func d7CardScore2(c string) int {
	// if we are a number, take the number
	n, err := strconv.Atoi(c)
	if err == nil {
		return n
	}

	switch c {
	case "T":
		return 10
	case "J":
		return 1
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	}

	panic("unknown card")
}

func d7GetHandType2(hand string) int {
	// if there are jokers, turning them into cards with the most count is always better

	cards := sp(hand, "")
	reqLen(cards, 5)

	cardMap := make(map[string]int)
	for _, card := range cards {
		cardMap[card]++
	}

	joker := cardMap["J"]
	if joker == 5 {
		// well
		return 0
	}

	if joker > 0 {
		// which one to add them to
		delete(cardMap, "J")
		var maxCard string
		var maxCount int
		for card, count := range cardMap {
			if count > maxCount {
				maxCard = card
				maxCount = count
			}
		}

		// max card is found
		hand = strings.ReplaceAll(hand, "J", maxCard)
	}

	return d7GetHandType(hand)
}
