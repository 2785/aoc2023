package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var d8TestInput = `
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

var d8TestInput2 = `
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

func TestD8P1(t *testing.T) {
	require.Equal(t, 6, doD8P1(d8TestInput))
}

func TestD8P2(t *testing.T) {
	require.Equal(t, 6, doD8P2(d8TestInput2))
}

func TestD8ParseMap(t *testing.T) {
	require.Equal(t, d8Map{
		curr: "BBB",
		l:    "DDD",
		r:    "EEE",
	}, d8ParseMapEntry("BBB = (DDD, EEE)"))
}
