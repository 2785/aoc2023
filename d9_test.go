package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var d9TestInput = sanitizeInput(`
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`)

func TestD9P1(t *testing.T) {
	require.Equal(t, 114, doD9P1(d9TestInput))
}

func TestD9P2(t *testing.T) {
	require.Equal(t, 2, doD9P2(d9TestInput))
}

func TestD9ExtrapolateSeries(t *testing.T) {
	require.Equal(t, 18, d9ExtrapolateRight([]int{0, 3, 6, 9, 12, 15}))
	require.Equal(t, 28, d9ExtrapolateRight([]int{1, 3, 6, 10, 15, 21}))
	require.Equal(t, 68, d9ExtrapolateRight([]int{10, 13, 16, 21, 30, 45}))

	require.Equal(t, -3, d9ExtrapolateLeft([]int{0, 3, 6, 9, 12, 15}))
	require.Equal(t, 0, d9ExtrapolateLeft([]int{1, 3, 6, 10, 15, 21}))
}
