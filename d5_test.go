package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInputD5 = `
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func TestD5P1(t *testing.T) {
	require.Equal(t, 35, doD5P1(testInputD5))
}

func TestD5P2(t *testing.T) {
	require.Equal(t, 46, doD5P2(testInputD5))
}

func TestD5ParseMap(t *testing.T) {

	inp := `
seed-to-soil map:
50 98 2
52 50 48
	`

	ident, entries := d5ParseMap(inp)
	require.Equal(t, "seed", ident.src)
	require.Equal(t, "soil", ident.dst)

	require.ElementsMatch(t, []d5MapEntry{
		{srcStart: 50, dstStart: 98, length: 2},
		{srcStart: 52, dstStart: 50, length: 48},
	}, entries)
}

func TestD5ConvertRange(t *testing.T) {
	e := d5MapEntry{srcStart: 2, dstStart: 5, length: 3}

	// if we give it 1 ~ 5, we should bave 2, 3, 4 full over lap, and 1 and 5 as non overlap blocks
	overlap, nonOverlap := e.tryConvertRange(d5Range{start: 1, end: 5})
	require.EqualValues(t, &d5Range{start: 5, end: 7}, overlap)
	require.ElementsMatch(t, []d5Range{
		{start: 1, end: 1},
		{start: 5, end: 5},
	}, nonOverlap)

	// if we give it 1 ~ 2 we should have 1 as non overlap, and 2 as overlap
	overlap, nonOverlap = e.tryConvertRange(d5Range{start: 1, end: 2})
	require.EqualValues(t, &d5Range{start: 5, end: 5}, overlap)
	require.ElementsMatch(t, []d5Range{
		{start: 1, end: 1},
	}, nonOverlap)

	// if we give it 1 ~ 1 we should have 1 as non overlap
	overlap, nonOverlap = e.tryConvertRange(d5Range{start: 1, end: 1})
	require.Nil(t, overlap)
	require.ElementsMatch(t, []d5Range{
		{start: 1, end: 1},
	}, nonOverlap)

	// if we give it 3 ~ 4 we should just overlap everything
	overlap, nonOverlap = e.tryConvertRange(d5Range{start: 3, end: 4})
	require.EqualValues(t, &d5Range{start: 6, end: 7}, overlap)
	require.Empty(t, nonOverlap)
}

func TestD5CollapseRange(t *testing.T) {
	// non overlapping ranges should be returned as is
	require.ElementsMatch(t, []d5Range{
		{1, 2},
		{4, 5},
	}, d5CollapseRanges([]d5Range{
		{1, 2},
		{4, 5},
	}))

	// adjacent ranges should be collapsed
	require.ElementsMatch(t, []d5Range{
		{1, 5},
	}, d5CollapseRanges([]d5Range{
		{1, 2},
		{3, 5},
	}))

	// overlapping ranges should be collapsed
	require.ElementsMatch(t, []d5Range{
		{1, 5},
	}, d5CollapseRanges([]d5Range{
		{2, 3},
		{1, 5},
	}))
}
