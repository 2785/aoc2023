package main

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d9p1)
	root.AddCommand(d9p2)
}

var d9p1 = &cobra.Command{
	Use: "d9p1",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("Result: %d", doD9P1(loadLines(9)))
	},
}

var d9p2 = &cobra.Command{
	Use: "d9p2",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("Result: %d", doD9P2(loadLines(9)))
	},
}

func doD9P1(inp []string) int {
	inputs := d9ParseInput(inp)
	var sum int
	for _, line := range inputs {
		val := d9ExtrapolateRight(line)
		sum += val
	}

	return sum
}

func doD9P2(inp []string) int {
	inputs := d9ParseInput(inp)
	var sum int
	for _, line := range inputs {
		val := d9ExtrapolateLeft(line)
		sum += val
	}

	return sum
}

func d9ParseInput(inp []string) [][]int {
	res := make([][]int, len(inp))
	for i, line := range inp {
		split := sp(line, " ")
		res[i] = lo.Map(split, func(item string, _ int) int { return atoi(item) })
	}
	return res
}

func d9ExtrapolateRight(s []int) int {
	if len(s) < 2 {
		panic("series must be at least 2 elements long")
	}

	// if all items are 0, we can confidently extrapolate the series
	if lo.EveryBy(s, func(item int) bool { return item == 0 }) {
		return 0
	}

	diffs := make([]int, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		diffs[i] = s[i+1] - s[i]
	}

	additionalDiffItem := d9ExtrapolateRight(diffs)
	return s[len(s)-1] + additionalDiffItem
}

func d9ExtrapolateLeft(s []int) int {
	if len(s) < 2 {
		panic("series must be at least 2 elements long")
	}

	// if all items are 0, we can confidently extrapolate the series
	if lo.EveryBy(s, func(item int) bool { return item == 0 }) {
		return 0
	}

	diffs := make([]int, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		diffs[i] = s[i+1] - s[i]
	}

	additionalDiffItem := d9ExtrapolateLeft(diffs)
	return s[0] - additionalDiffItem
}
