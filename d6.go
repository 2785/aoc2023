package main

import (
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d6p1)
	root.AddCommand(d6p2)
}

var d6p1 = &cobra.Command{
	Use: "d6p1",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("ways: %d", doD6P1(loadLines(6)))
	},
}

var d6p2 = &cobra.Command{
	Use: "d6p2",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("ways: %d", doD6P2(loadLines(6)))
	},
}

// this is massively suspicious
func doD6P1(input []string) int {
	out := 1

	races := parseD6Input(input)
	for _, race := range races {
		ways := 0

		for i := 0; i <= race.Time; i++ {
			speed := i
			timeLeft := race.Time - i

			if speed*timeLeft > race.DistToBeat {
				ways++
			}
		}

		out *= ways
	}

	return out
}

// I knew it
func doD6P2(input []string) int {
	// let's binary search this
	reqLen(input, 2)

	timeStr := input[0]
	distStr := input[1]

	// chop off the header
	timeStr = mustTrimPrefix(timeStr, "Time:")
	distStr = mustTrimPrefix(distStr, "Distance:")

	timeStr = strings.ReplaceAll(timeStr, " ", "")
	distStr = strings.ReplaceAll(distStr, " ", "")

	t := atoi(timeStr)
	d := atoi(distStr)

	win := func(n int) bool {
		return n*(t-n) > d
	}

	// grab the center point
	center := t / 2
	// this must beat the record
	if center*(t-center) < d {
		panic("bad input")
	}

	// to the left, we are searching for an item that beats it, but the left of it doesn't beat it

	var left, right int

	cl, cr := 0, center
	sel := center / 2

	validLeft := func(n int) bool {
		// is it valid itself?
		return win(n) && !win(n-1)
	}

	for {
		if validLeft(sel) {
			break
		}

		if win(sel) {
			// we need to go to the left
			cr = sel
		} else {
			// we need to go to the right
			cl = sel
		}

		sel = (cl + cr) / 2
	}

	left = sel

	cl, cr = center, t
	sel = (cl + cr) / 2

	validRight := func(n int) bool {
		// is it valid itself?
		return win(n) && !win(n+1)
	}

	for {
		if validRight(sel) {
			break
		}

		if win(sel) {
			// we need to go to the right
			cl = sel
		} else {
			// we need to go to the left
			cr = sel
		}

		sel = (cl + cr) / 2
	}

	right = sel

	return right - left + 1
}

type d6Race struct {
	Time       int
	DistToBeat int
}

func parseD6Input(inp []string) []d6Race {
	reqLen(inp, 2)

	timeStr := inp[0]
	distStr := inp[1]

	// chop off the header
	timeStr = mustTrimPrefix(timeStr, "Time:")
	distStr = mustTrimPrefix(distStr, "Distance:")

	times := lo.Map(sp(timeStr, " "), func(s string, _ int) int {
		return atoi(s)
	})

	dists := lo.Map(sp(distStr, " "), func(s string, _ int) int {
		return atoi(s)
	})

	if len(times) != len(dists) {
		panic("bad input")
	}

	ret := make([]d6Race, len(times))
	for i, time := range times {
		ret[i] = d6Race{
			Time:       time,
			DistToBeat: dists[i],
		}
	}

	return ret
}
