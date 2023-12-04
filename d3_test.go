package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInputD3 = sanitizeInput(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`)

func TestD3P1(t *testing.T) {
	require.Equal(t, 4361, doD3P1(testInputD3))
}

func TestD3P2(t *testing.T) {
	require.Equal(t, 467835, doD3P2(testInputD3))
}

func TestD3Around(t *testing.T) {
	loc := d3Loc{ln: 0, start: 0, end: 1}
	// this should have no neighbors in a 1x1 grid
	require.Equal(t, []d3Loc{}, loc.around(1, 1))

	// this should have 3 neighbors in a 2x2 grid
	require.ElementsMatch(t, []d3Loc{
		{ln: 0, start: 1, end: 2},
		{ln: 1, start: 0, end: 1},
		{ln: 1, start: 1, end: 2},
	}, loc.around(2, 2))

	// this should have 3 neighbors in a 5x5 grid
	require.ElementsMatch(t, []d3Loc{
		{ln: 0, start: 1, end: 2},
		{ln: 1, start: 0, end: 1},
		{ln: 1, start: 1, end: 2},
	}, loc.around(5, 5))

	// grab a 2 length thing in a 5x5 grind
	loc = d3Loc{ln: 2, start: 2, end: 4}
	// this should have 10 neighbors in a 5x5 grid
	require.ElementsMatch(t, []d3Loc{
		{ln: 1, start: 1, end: 2},
		{ln: 1, start: 2, end: 3},
		{ln: 1, start: 3, end: 4},
		{ln: 1, start: 4, end: 5},
		{ln: 2, start: 1, end: 2},
		{ln: 2, start: 4, end: 5},
		{ln: 3, start: 1, end: 2},
		{ln: 3, start: 2, end: 3},
		{ln: 3, start: 3, end: 4},
		{ln: 3, start: 4, end: 5},
	}, loc.around(5, 5))

	// 3 length thing in a 5x5 grind
	loc = d3Loc{ln: 2, start: 2, end: 5}
	// this should have 9 neighbors in a 5x5 grid
	require.ElementsMatch(t, []d3Loc{
		{ln: 1, start: 1, end: 2},
		{ln: 1, start: 2, end: 3},
		{ln: 1, start: 3, end: 4},
		{ln: 1, start: 4, end: 5},
		{ln: 2, start: 1, end: 2},
		{ln: 3, start: 1, end: 2},
		{ln: 3, start: 2, end: 3},
		{ln: 3, start: 3, end: 4},
		{ln: 3, start: 4, end: 5},
	}, loc.around(5, 5))
}

func TestD3ParseLine(t *testing.T) {
	inp := "617*......"
	parsed := d3ParseLine(inp, 0)
	require.Equal(t, map[d3Loc]d3Thing{
		{ln: 0, start: 0, end: 3}: {typ: "num", val: 617},
		{ln: 0, start: 3, end: 4}: {typ: "thing"},
		{ln: 0, start: 0, end: 1}: {typ: "thing"},
		{ln: 0, start: 1, end: 2}: {typ: "thing"},
		{ln: 0, start: 2, end: 3}: {typ: "thing"},
	}, parsed)
}
