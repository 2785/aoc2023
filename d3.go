package main

import (
	"strconv"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d3p1)
	root.AddCommand(d3p2)
}

var d3p1 = &cobra.Command{
	Use: "d3p1",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("sum: %d", doD3P1(loadLines(3)))
	},
}

var d3p2 = &cobra.Command{
	Use: "d3p2",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("sum: %d", doD3P2(loadLines(3)))
	},
}

func doD3P1(inp []string) int {
	graph := d3ParseGraph(inp)

	var sum int

allThings:
	for loc, thing := range graph {
		if thing.typ == "thing" {
			continue
		}

		around := loc.around(len(inp), len(inp[0]))

		for _, l := range around {
			if graph[l].typ == "thing" {
				sum += thing.val
				continue allThings
			}
		}
	}

	return sum
}

func doD3P2(inp []string) int {
	// go through the numbers anyway, but let's do a reverse lookup
	lookup := map[d3Loc][]int{}
	graph := d3ParseGraph(inp)

	for loc, thing := range graph {
		if thing.typ == "thing" {
			continue
		}

		around := loc.around(len(inp), len(inp[0]))

		for _, l := range around {
			if graph[l].typ == "thing" && graph[l].symbol == "*" {
				lookup[l] = append(lookup[l], thing.val)
			}
		}
	}

	var sum int
	for _, nums := range lookup {
		if len(nums) == 2 {
			sum += nums[0] * nums[1]
		}
	}

	return sum
}

func d3ParseGraph(inp []string) map[d3Loc]d3Thing {
	ret := map[d3Loc]d3Thing{}

	for i, line := range inp {
		parsed := d3ParseLine(line, i)
		for k, v := range parsed {
			ret[k] = v
		}
	}

	return ret
}

func d3ParseLine(l string, ln int) map[d3Loc]d3Thing {
	ret := map[d3Loc]d3Thing{}

	split := sp(l, "")
	var currNum *string
	var currStart int

	for i, char := range split {
		i := i
		char := char
		// if we're on a number
		if _, err := strconv.Atoi(char); err == nil {
			if currNum == nil {
				currNum = &char
				currStart = i
			} else {
				currNum = lo.ToPtr(*currNum + char)
			}

			continue
		}

		// we are not a number, if we currently have a number we need to terminate it and add it to
		// the set
		if currNum != nil {
			num, err := strconv.Atoi(*currNum)
			c(err)

			ret[d3Loc{
				ln:    ln,
				start: currStart,
				end:   i,
			}] = d3Thing{
				typ: "num",
				val: num,
			}

			currNum = nil
		}

		// let's see what thing we are on, if we're on a "." we can skip
		if char == "." {
			continue
		}

		// otherwise we're on a thing
		ret[d3Loc{
			ln:    ln,
			start: i,
			end:   i + 1,
		}] = d3Thing{
			typ:    "thing",
			symbol: char,
		}
	}

	// if at this point we have a number still, wrap that up
	if currNum != nil {
		num, err := strconv.Atoi(*currNum)
		c(err)

		ret[d3Loc{
			ln:    ln,
			start: currStart,
			end:   len(split),
		}] = d3Thing{
			typ: "num",
			val: num,
		}
	}

	return ret
}

type d3Loc struct {
	ln    int
	start int
	end   int // not inclusive
}

type d3Thing struct {
	typ    string
	val    int
	symbol string
}

func (l d3Loc) around(lc, rc int) []d3Loc {
	ret := []d3Loc{}

	// left & right
	if l.start > 0 {
		ret = append(ret, d3Loc{
			ln:    l.ln,
			start: l.start - 1,
			end:   l.start,
		})
	}

	if l.end < rc {
		ret = append(ret, d3Loc{
			ln:    l.ln,
			start: l.end,
			end:   l.end + 1,
		})
	}

	// consider top if not first line
	if l.ln > 0 {
		// corners

		// consider left
		if l.start > 0 {
			ret = append(ret, d3Loc{
				ln:    l.ln - 1,
				start: l.start - 1,
				end:   l.start,
			})
		}

		// consider right
		if l.end < rc {
			ret = append(ret, d3Loc{
				ln:    l.ln - 1,
				start: l.end,
				end:   l.end + 1,
			})
		}

		// in between
		for i := l.start; i < l.end; i++ {
			ret = append(ret, d3Loc{
				ln:    l.ln - 1,
				start: i,
				end:   i + 1,
			})
		}
	}

	// consider bottom if not last line
	if l.ln < lc-1 {
		// consider left
		if l.start > 0 {
			ret = append(ret, d3Loc{
				ln:    l.ln + 1,
				start: l.start - 1,
				end:   l.start,
			})
		}

		// consider right
		if l.end < rc {
			ret = append(ret, d3Loc{
				ln:    l.ln + 1,
				start: l.end,
				end:   l.end + 1,
			})
		}

		// in between
		for i := l.start; i < l.end; i++ {
			ret = append(ret, d3Loc{
				ln:    l.ln + 1,
				start: i,
				end:   i + 1,
			})
		}
	}

	return ret
}
