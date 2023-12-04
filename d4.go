package main

import (
	"github.com/gammazero/deque"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d4p1)
	root.AddCommand(d4p2)
}

var d4p1 = &cobra.Command{
	Use: "d4p1",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("sum: %d", doD4P1(loadLines(4)))
	},
}

var d4p2 = &cobra.Command{
	Use: "d4p2",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("sum: %d", doD4P2(loadLines(4)))
	},
}

func doD4P1(inp []string) int {
	cards := d4ParseCards(inp)

	sum := 0

	for _, card := range cards {

		matches := 0

		win := lo.KeyBy(card.winningNumbers, func(n int) int { return n })
		for _, n := range card.yourNumbers {
			if _, ok := win[n]; ok {
				matches++
			}
		}

		if matches != 0 {
			sum += pow(2, matches-1)
		}
	}

	return sum
}

// god damn scratch cards
func doD4P2(inp []string) int {
	cards := d4ParseCards(inp)

	// let's work out how many winnings each of these cards has
	matches := make([]int, len(cards))

	for i, card := range cards {
		win := lo.KeyBy(card.winningNumbers, func(n int) int { return n })
		for _, n := range card.yourNumbers {
			if _, ok := win[n]; ok {
				matches[i]++
			}
		}
	}

	tot := 0

	// now that's sorted out, time for a fifo queue
	q := deque.New[int]()
	// push all the original cards into the queue
	for i := range cards {
		tot += 1
		q.PushBack(i)
	}

	// now we start popping and pushing
	for q.Len() > 0 {
		thing := q.PopFront()
		matches := matches[thing]

		if matches > 0 {
			for i := 0; i < matches; i++ {
				tot += 1
				q.PushBack(thing + i + 1)
			}
		}
	}

	return tot
}

func d4ParseCards(inp []string) []d4Card {
	cards := make([]d4Card, len(inp))

	for i, s := range inp {
		cards[i] = d4ParseCard(s)
	}

	return cards
}

type d4Card struct {
	winningNumbers []int
	yourNumbers    []int
}

func d4ParseCard(input string) d4Card {
	// chop off the card number
	split := sp(input, ":")
	reqLen(split, 2)

	input = split[1]

	split = sp(input, "|")
	reqLen(split, 2)

	// lhs is the winning numbers
	lhs := sp(split[0], " ")
	rhs := sp(split[1], " ")

	card := d4Card{
		winningNumbers: lo.Map(lhs, func(s string, _ int) int {
			return atoi(s)
		}),
		yourNumbers: lo.Map(rhs, func(s string, _ int) int {
			return atoi(s)
		}),
	}

	return card
}
