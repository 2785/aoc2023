package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var d7TestInput = sanitizeInput(`
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`)

func TestD7P1(t *testing.T) {
	require.Equal(t, 6440, doD7P1(d7TestInput))
}

func TestD7P2(t *testing.T) {
	require.Equal(t, 5905, doD7P2(d7TestInput))
}

func TestD7GetHandType(t *testing.T) {
	require.Equal(t, 0, d7GetHandType("AAAAA"))
	require.Equal(t, -1, d7GetHandType("AAAAK"))
	require.Equal(t, -2, d7GetHandType("AAABB"))
	require.Equal(t, -3, d7GetHandType("AAABC"))
	require.Equal(t, -4, d7GetHandType("AABBD"))
	require.Equal(t, -5, d7GetHandType("AABCD"))
	require.Equal(t, -6, d7GetHandType("ABCDE"))

	require.True(t, d7HighCardSmaller1("12345", "12346"))
	require.False(t, d7HighCardSmaller1("12345", "12345"))
	require.False(t, d7HighCardSmaller1("A1234", "T1234"))
	require.True(t, d7HighCardSmaller1("A1234", "A1235"))
	require.True(t, d7HighCardSmaller1("5967J", "8A745"))
	require.True(t, d7HighCardSmaller1("83J27", "T73K6"))
	require.False(t, d7HighCardSmaller1("83J27", "T73K6"))
}
