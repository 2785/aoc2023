package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInputD2 = sanitizeInput(`
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`)

func TestD2P1(t *testing.T) {
	require.Equal(t, 8, doD2P1(testInputD2))
}

func TestD2P2(t *testing.T) {
	require.Equal(t, 2286, doD2P2(testInputD2))
}

func TestD2ParseBag(t *testing.T) {
	bag := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	parsed := d2parseBag(bag)
	require.Equal(t, []d2Draw{
		{R: 4, G: 0, B: 3},
		{R: 1, G: 2, B: 6},
		{R: 0, G: 2, B: 0},
	}, parsed)
}
