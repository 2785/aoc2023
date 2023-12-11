package main

import (
	"math"
	"sort"

	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d10p1)
	root.AddCommand(d10p2)
}

var infinity = math.MaxInt

var d10p1 = &cobra.Command{
	Use: "d10p1",
	Run: func(cmd *cobra.Command, args []string) {
		s.Infof("result: %d", doD10P1(loadLines(10)))
	},
}

var d10p2 = &cobra.Command{
	Use: "d10p2",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func doD10P1(inp []string) int {
	// shrug, dijkstra
	_, start := d10ParseMap(inp)

	// we are looking for start, again

	// explore start's neighbors

	q := []*d10Point{}

	push := func(p *d10Point) {
		p.seen = true
		q = append(q, p)
		sort.Slice(q, func(i, j int) bool {
			return q[i].dist < q[j].dist
		})
	}

	take := func() *d10Point {
		p := q[0]
		p.visited = true
		q = q[1:]
		return p
	}

	// start is special, we need to figure out which one is a valid neighbor

	// we need to make start the end
	start.dist = infinity
	start.seen = false
	start.visited = false

	replaceStart := func(p *d10Point) {
		p.seen = true
		p.visited = true
		p.dist = 0

		for _, n := range p.neighbors() {
			n := n
			if n.val != "S" {
				n.dist = 1
				push(n)
			}
		}
	}

	// top can be |, 7, F
	t, ok := start.t()
	if ok && t.val == "|" || t.val == "7" || t.val == "F" {
		replaceStart(t)
	} else if b, ok := start.b(); ok && b.val == "|" || b.val == "J" || b.val == "L" {
		replaceStart(b)
	} else if l, ok := start.l(); ok && l.val == "-" || l.val == "F" || l.val == "L" {
		replaceStart(l)
	} else if r, ok := start.r(); ok && r.val == "-" || r.val == "J" || r.val == "7" {
		replaceStart(r)
	} else {
		panic("no valid neighbor")
	}

	for len(q) > 0 {
		p := take()

		// if we are S, we're done
		if p.val == "S" {
			return (p.dist + 1) / 2
		}

		// explore neighbors
		if p.val == "." {
			panic("what are we doing on a .")
		}

		neighbors := p.neighbors()
		for _, n := range neighbors {
			n := n // just copy the pointer
			if n.visited {
				continue
			}

			if !n.seen {
				push(n)
			}

			// update the distance
			if n.dist > p.dist+1 {
				n.dist = p.dist + 1
			}
		}
	}

	panic("no path found")
}

func d10ParseMap(inp []string) (d10Map, *d10Point) {
	res := make(d10Map, len(inp))

	ymax := len(inp) - 1

	var start *d10Point

	// iterate it backward so that my brain can process the y axis
	for i, line := range inp {
		y := ymax - i
		res[y] = make([]d10Point, len(line))
		for j, char := range line {
			res[y][j] = d10Point{
				m:    res,
				x:    j,
				y:    y,
				val:  string(char),
				dist: infinity,
			}

			if char == 'S' {
				res[y][j].dist = 0
				res[y][j].seen = true
				start = &res[y][j]
			}
		}
	}

	return res, start
}

type d10Map [][]d10Point

func (d d10Map) Get(x, y int) d10Point {
	return d[y][x]
}

func (d d10Map) Set(x, y int, p d10Point) {
	d[y][x] = p
}

type d10Point struct {
	m       d10Map
	x       int
	y       int
	val     string
	visited bool
	seen    bool
	dist    int
}

func (p *d10Point) t() (*d10Point, bool) {
	if p.y == len(p.m)-1 {
		return p, false
	}

	return &p.m[p.y+1][p.x], true
}

func (p *d10Point) b() (*d10Point, bool) {
	if p.y == 0 {
		return p, false
	}

	return &p.m[p.y-1][p.x], true
}

func (p *d10Point) l() (*d10Point, bool) {
	if p.x == 0 {
		return p, false
	}

	return &p.m[p.y][p.x-1], true
}

func (p *d10Point) r() (*d10Point, bool) {
	if p.x == len(p.m[0])-1 {
		return p, false
	}

	return &p.m[p.y][p.x+1], true
}

func (p *d10Point) neighbors() []*d10Point {
	res := make([]*d10Point, 0, 4)

	// t
	if p.val == "|" || p.val == "J" || p.val == "L" {
		t, ok := p.t()
		if ok {
			res = append(res, t)
		}
	}

	// b
	if p.val == "|" || p.val == "F" || p.val == "7" {
		b, ok := p.b()
		if ok {
			res = append(res, b)
		}
	}

	// l
	if p.val == "-" || p.val == "J" || p.val == "7" {
		l, ok := p.l()
		if ok {
			res = append(res, l)
		}
	}

	// r
	if p.val == "-" || p.val == "F" || p.val == "L" {
		r, ok := p.r()
		if ok {
			res = append(res, r)
		}
	}

	return res
}
