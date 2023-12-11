package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var d11TestInput = sanitizeInput(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`)

func TestD11P1(t *testing.T) {
	require.Equal(t, 374, doD11(d11TestInput, 2))
}

func TestD11P2(t *testing.T) {
	require.Equal(t, 1030, doD11(d11TestInput, 10))
	require.Equal(t, 8410, doD11(d11TestInput, 100))
}

func TestD1ParseMap(t *testing.T) {
	galaxies := d11ParseMap(d11TestInput)
	require.ElementsMatch(t, []d11Galaxy{
		{0, 0},
		{4, 0},
		{7, 1},
		{9, 3},
		{1, 4},
		{6, 5},
		{0, 7},
		{7, 8},
		{3, 9},
	}, galaxies)

	grew := d11GrowMap(galaxies, 1)
	require.ElementsMatch(t, d11ParseMap(sanitizeInput(`
	....#........
	.........#...
	#............
	.............
	.............
	........#....
	.#...........
	............#
	.............
	.............
	.........#...
	#....#.......
	`)), grew)
}

func TestDist(t *testing.T) {
	require.Equal(t, 0, dist(d11Galaxy{0, 0}, d11Galaxy{0, 0}))
	require.Equal(t, 1, dist(d11Galaxy{0, 0}, d11Galaxy{0, 1}))
	require.Equal(t, 3, dist(d11Galaxy{0, 0}, d11Galaxy{2, 1}))
}
