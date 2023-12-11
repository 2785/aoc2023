package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var d10TestInput = sanitizeInput(`
..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`)

func TestD10P1(t *testing.T) {
	require.Equal(t, 8, doD10P1(d10TestInput))
}

func TestD10P2(t *testing.T) {

}

func TestD10ParseMap(t *testing.T) {
	m, start := d10ParseMap(d10TestInput)
	require.Equal(t, 5, len(m))
	require.Equal(t, 5, len(m[0]))
	require.Equal(t, "S", start.val)
	require.Equal(t, 0, start.x)
	require.Equal(t, 2, start.y)

	top, ok := start.t()
	require.True(t, ok)

	require.Equal(t, 0, top.x)
	require.Equal(t, 3, top.y)
	require.Equal(t, ".", top.val)

	// if we modify the map, it should be reflected in the point
	m[3][0].val = "X"
	require.Equal(t, "X", top.val)
}
