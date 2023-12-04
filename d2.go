package main

import (
	"strconv"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d2p1)
	root.AddCommand(d2p2)
}

var d2p1 = &cobra.Command{
	Use: "d2p1",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("sum: %d", doD2P1(loadLines(2)))
	},
}

var d2p2 = &cobra.Command{
	Use: "d2p2",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("sum: %d", doD2P2(loadLines(2)))
	},
}

func doD2P1(inp []string) int {
	games := lo.Map(inp, func(s string, _ int) []d2Draw {
		draws := d2parseBag(s)
		return draws
	})

	possibleGames := []int{}

	r, g, b := 12, 13, 14
outer:
	for i, game := range games {
		for _, draw := range game {
			if draw.R > r || draw.G > g || draw.B > b {
				continue outer
			}
		}

		possibleGames = append(possibleGames, i+1)
	}

	return lo.Sum(possibleGames)
}

func doD2P2(inp []string) int {
	games := lo.Map(inp, func(s string, _ int) []d2Draw {
		draws := d2parseBag(s)
		return draws
	})

	var sum int

	for _, game := range games {
		var r, g, b int

		for _, draw := range game {
			if draw.R > r {
				r = draw.R
			}

			if draw.G > g {
				g = draw.G
			}

			if draw.B > b {
				b = draw.B
			}
		}

		sum += r * g * b
	}

	return sum
}

type d2Draw struct {
	R, G, B int
}

func d2parseBag(bag string) []d2Draw {
	// split by ":"
	split := sp(bag, ":")
	reqLen(split, 2)

	bag = split[1]

	// split by ";"
	split = sp(bag, ";")

	var res []d2Draw
	for _, d := range split {
		// split by ","
		split = sp(d, ",")

		draw := d2Draw{}

		for _, part := range split {
			// split by " "
			split = sp(part, " ")
			reqLen(split, 2)

			num, err := strconv.Atoi(split[0])
			c(err)

			switch split[1] {
			case "blue":
				draw.B = num
			case "red":
				draw.R = num
			case "green":
				draw.G = num
			default:
				s.Panicf("unknown color: %s", split[1])
			}
		}

		res = append(res, draw)
	}

	return res
}
