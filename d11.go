package main

import (
	"sort"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d11p1)
	root.AddCommand(d11p2)
}

var d11p1 = &cobra.Command{
	Use: "d11p1",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("Answer: %d", doD11(loadLines(11), 2))
	},
}

var d11p2 = &cobra.Command{
	Use: "d11p2",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("Answer: %d", doD11(loadLines(11), 1000000))
	},
}

func doD11(inp []string, growthFactor int) int {
	galaxies := d11ParseMap(inp)

	galaxies = d11GrowMap(galaxies, growthFactor)

	l := len(galaxies)

	var sum int

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			sum += dist(galaxies[i], galaxies[j])
		}
	}

	return sum
}

// I'm gonna assume doing this dense is a very very bad idea, so let's do this sparse
type d11Galaxy struct {
	x, y int
}

func d11ParseMap(inp []string) []d11Galaxy {
	galaxies := []d11Galaxy{}

	maxY := len(inp) - 1

	// again, let's invert y so my brain is happy
	for lid, line := range inp {
		y := maxY - lid

		for x, c := range line {
			if c == '#' {
				galaxies = append(galaxies, d11Galaxy{x, y})
			}
		}
	}

	return galaxies
}

func d11GrowMap(galaxies []d11Galaxy, growthFactor int) []d11Galaxy {
	// deal with one direction at a time

	// let's just always make new copies of these galaxy things

	// x direction
	sort.Slice(galaxies, func(i, j int) bool {
		return galaxies[i].x < galaxies[j].x
	})

	xMin, xMax := galaxies[0].x, galaxies[len(galaxies)-1].x

	// group by x
	groupedByX := lo.GroupBy(galaxies, func(g d11Galaxy) int { return g.x })

	var xOffset int

	xExpanded := make([]d11Galaxy, 0)

	for x := xMin; x <= xMax; x++ {
		gals, ok := groupedByX[x]
		if !ok {
			xOffset += (growthFactor - 1)
			continue
		}

		// found galaxies
		for _, g := range gals {
			xExpanded = append(xExpanded, d11Galaxy{g.x + xOffset, g.y})
		}
	}

	// now do y
	sort.Slice(xExpanded, func(i, j int) bool {
		return xExpanded[i].y < xExpanded[j].y
	})

	yMin, yMax := xExpanded[0].y, xExpanded[len(xExpanded)-1].y

	// group by y
	groupedByY := lo.GroupBy(xExpanded, func(g d11Galaxy) int { return g.y })

	var yOffset int

	yExpanded := make([]d11Galaxy, 0)

	for y := yMin; y <= yMax; y++ {
		gals, ok := groupedByY[y]
		if !ok {
			yOffset += (growthFactor - 1)
			continue
		}

		// found galaxies
		for _, g := range gals {
			yExpanded = append(yExpanded, d11Galaxy{g.x, g.y + yOffset})
		}
	}

	return yExpanded
}

func dist(a, b d11Galaxy) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
