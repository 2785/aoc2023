package main

import (
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d8p1)
	root.AddCommand(d8p2)
}

var d8p1 = &cobra.Command{
	Use: "d8p1",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("steps: %d", doD8P1(loadFileAsString(8)))
	},
}

var d8p2 = &cobra.Command{
	Use: "d8p2",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("steps: %d", doD8P2(loadFileAsString(8)))
	},
}

func doD8P1(input string) int {
	inp := d8ParseInput(input)

	curr := "AAA"
	steps := 0

	currInstIndex := 0
	nextInst := func() bool {
		if currInstIndex >= len(inp.instruction) {
			currInstIndex = 0
		}

		defer func() { currInstIndex++ }()
		return inp.instruction[currInstIndex]
	}

	for curr != "ZZZ" {
		m, ok := inp.maps[curr]
		if !ok {
			panic("invalid map, can't find " + curr)
		}

		left := nextInst()
		if left {
			curr = m.l
		} else {
			curr = m.r
		}

		steps++
	}

	return steps
}

func doD8P2(input string) int {
	// ofc it wants you to do lcm
	inp := d8ParseInput(input)

	keys := lo.Keys(inp.maps)
	curr := lo.Filter(keys, func(s string, _ int) bool { return strings.HasSuffix(s, "A") })

	steps := lo.Map(curr, func(s string, _ int) int {
		curr := s
		steps := 0

		currInstIndex := 0
		nextInst := func() bool {
			if currInstIndex >= len(inp.instruction) {
				currInstIndex = 0
			}

			defer func() { currInstIndex++ }()
			return inp.instruction[currInstIndex]
		}

		done := func() bool {
			return strings.HasSuffix(curr, "Z")
		}

		for !done() {
			m, ok := inp.maps[curr]
			if !ok {
				panic("invalid map, can't find " + curr)
			}

			left := nextInst()
			if left {
				curr = m.l
			} else {
				curr = m.r
			}

			steps++
		}

		return steps
	})

	return lcm(steps...)
}

type d8Map struct {
	curr string
	l, r string
}

func d8ParseMapEntry(input string) d8Map {
	split := sp(input, " = ")
	reqLen(split, 2)

	l := split[0]
	r := split[1]

	// r needs to be stripped of brackets
	r = r[1 : len(r)-1]
	split = sp(r, ",")
	reqLen(split, 2)

	return d8Map{
		curr: l,
		l:    split[0],
		r:    split[1],
	}
}

type d8Input struct {
	instruction []bool
	maps        map[string]d8Map
}

func d8ParseInput(input string) d8Input {
	split := sp(input, "\n\n")
	reqLen(split, 2)

	instRaw := split[0]
	mapsRaw := split[1]

	instsRaw := sp(instRaw, "")
	insts := lo.Map(instsRaw, func(s string, _ int) bool {
		switch s {
		case "L":
			return true
		case "R":
			return false
		default:
			panic("invalid instruction")
		}
	})

	mapsLines := sp(mapsRaw, "\n")

	maps := lo.Map(mapsLines, func(s string, _ int) d8Map {
		return d8ParseMapEntry(s)
	})

	return d8Input{
		instruction: insts,
		maps:        lo.KeyBy(maps, func(m d8Map) string { return m.curr }),
	}
}

func lcm(nums ...int) int {
	if len(nums) < 2 {
		panic("need at least 2 numbers")
	}

	a := nums[0]
	b := nums[1]

	if len(nums) > 2 {
		return lcm(a, lcm(nums[1:]...))
	}

	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
