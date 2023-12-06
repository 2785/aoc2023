package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInputD6 = sanitizeInput(`
Time:      7  15   30
Distance:  9  40  200
`)

func TestD6P1(t *testing.T) {
	require.Equal(t, 288, doD6P1(testInputD6))
}

func TestD6P2(t *testing.T) {
	require.Equal(t, 71503, doD6P2(testInputD6))
}

func TestD6ParseInput(t *testing.T) {
	parsed := parseD6Input(testInputD6)
	require.ElementsMatch(t, []d6Race{
		{7, 9}, {15, 40}, {30, 200},
	}, parsed)
}
