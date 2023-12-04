package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInputD1 = sanitizeInput(`
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`)

var testInputD1P2 = sanitizeInput(`
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`)

func TestD1P1(t *testing.T) {
	require.Equal(t, 142, doD1P1(testInputD1))
}

func TestD1P2(t *testing.T) {
	s := "abcone2threexyz"
	num, ok, left := parseChunk(s)
	require.True(t, ok)
	require.Equal(t, 1, num)
	require.Equal(t, "2threexyz", left)

	num, ok, left = parseChunk(left)
	require.True(t, ok)
	require.Equal(t, 2, num)
	require.Equal(t, "threexyz", left)

	num, ok, left = parseChunk(left)
	require.True(t, ok)
	require.Equal(t, 3, num)
	require.Equal(t, "xyz", left)

	num, ok, left = parseChunk(left)
	require.False(t, ok)
	require.Equal(t, 0, num)
	require.Equal(t, "", left)

	s = `xtwone3four`
	num, ok, left = parseChunk(s)
	require.True(t, ok)
	require.Equal(t, 2, num)
	require.Equal(t, "one3four", left)

	require.Equal(t, 281, doD1P2(testInputD1P2))
}
