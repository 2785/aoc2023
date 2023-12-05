package main

import (
	"sort"

	"github.com/gammazero/deque"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d5p1)
	root.AddCommand(d5p2)
}

var d5p1 = &cobra.Command{
	Use: "d5p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := loadFileAsString(5)
		s.Infof("min: %d", doD5P1(inp))
	},
}

var d5p2 = &cobra.Command{
	Use: "d5p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := loadFileAsString(5)
		s.Infof("min: %d", doD5P2(inp))
	},
}

func doD5P1(input string) int {
	seeds, maps := d5ParseInput(input)

	return evaluateLowestLoc(seeds, maps)
}

// hot mess? yes, kind of works? also yes
func doD5P2(input string) int {
	// well, it sure doesn't want to do the trivial method, as opposed to tracking each individual
	// seeds we'll need to run this in ranges

	seedRangesRaw, maps := d5ParseInput(input)

	if len(seedRangesRaw)%2 != 0 {
		panic("bad seed ranges")
	}

	chunks := lo.Chunk(seedRangesRaw, 2)

	// let's process these seed ranges into d5 ranges
	seedRanges := lo.Map(chunks, func(chunk []int, _ int) d5Range {
		return d5Range{
			start: chunk[0],
			end:   chunk[0] + chunk[1] - 1,
		}
	})

	// make a lookup, again
	keys := lo.Keys(maps)
	lookup := lo.KeyBy(keys, func(ident d5MapIdent) string { return ident.src })

	// for each of these seed ranges let's process them over maps
	curr := "seed"
	currRanges := seedRanges
	for curr != "location" {
		ident := lookup[curr]

		mapEntries := maps[ident]

		// apply the map to the values
		var newRanges []d5Range
		// do up a deque to hold unprocessed ranges
		unprocessedRanges := deque.New[d5Range]()
		for _, r := range currRanges {
			r := r
			unprocessedRanges.PushBack(r)
		}

	Pop:
		for unprocessedRanges.Len() > 0 {
			// pop one and try to process it
			r := unprocessedRanges.PopFront()
			for _, entry := range mapEntries {
				// try to convert the range
				convertedRange, unconvertedRanges := entry.tryConvertRange(r)

				if convertedRange != nil {
					// we're done with this current record, we can add the unprocessed bits into the new range
					newRanges = append(newRanges, *convertedRange)
					for _, r := range unconvertedRanges {
						unprocessedRanges.PushBack(r)
					}

					// we can break back into the for loop
					continue Pop
				}
			}

			// this means we didn't convert anything, just push the range back into the new ranges
			newRanges = append(newRanges, r)
		}

		// ugh, we've the new ranges, I guess we collapse them together
		currRanges = d5CollapseRanges(newRanges)
		curr = ident.dst
	}

	// now we have the final qualifying ranges, I guess we just... take the min?
	if len(currRanges) == 0 {
		panic("no qualifying ranges")
	}

	min := currRanges[0].start
	return min
}

func evaluateLowestLoc(seeds []int, maps map[d5MapIdent][]d5MapEntry) int {
	// make a lookup of src to the ident
	keys := lo.Keys(maps)
	lookup := lo.KeyBy(keys, func(ident d5MapIdent) string { return ident.src })

	// now we have the lookup

	curr := "seed"
	currVals := seeds
	for curr != "location" {
		// find the next ident
		ident := lookup[curr]

		// find the map
		mapEntries := maps[ident]

		// apply the map to the values
		currVals = parallel.Map(currVals, func(val int, _ int) int {
			for _, entry := range mapEntries {
				if newVal, ok := entry.tryConvert(val); ok {
					return newVal
				}
			}

			return val
		})

		curr = ident.dst
	}

	// now we have the final values
	return lo.Min(currVals)
}

func d5ParseInput(inp string) ([]int, map[d5MapIdent][]d5MapEntry) {
	chunks := sp(inp, "\n\n")
	// first chunk is the seeds
	seedsRaw := chunks[0]

	// chop off the "seeds:" part
	seedsRaw = mustTrimPrefix(seedsRaw, "seeds:")
	// now we split by space
	seedSplit := sp(seedsRaw, " ")
	// and we turn it into numbers
	seeds := lo.Map(seedSplit, func(s string, _ int) int { return atoi(s) })

	// now we parse the maps
	ret := make(map[d5MapIdent][]d5MapEntry)
	for _, chunk := range chunks[1:] {
		ident, entries := d5ParseMap(chunk)
		ret[ident] = entries
	}

	return seeds, ret
}

func d5ParseMap(inp string) (ident d5MapIdent, entries []d5MapEntry) {
	lines := sp(inp, "\n")
	// first line is the ident
	identRaw := lines[0]
	lines = lines[1:]

	identSplit := sp(identRaw, " ")
	reqLen(identSplit, 2)

	// split by "-"
	identSplit = sp(identSplit[0], "-")
	reqLen(identSplit, 3)

	ident.src = identSplit[0]
	ident.dst = identSplit[2]

	// done processing ident, now process entries
	for _, line := range lines {
		split := sp(line, " ")
		reqLen(split, 3)

		ent := d5MapEntry{
			dstStart: atoi(split[0]),
			srcStart: atoi(split[1]),
			length:   atoi(split[2]),
		}

		entries = append(entries, ent)
	}

	return ident, entries
}

type d5MapIdent struct {
	src string
	dst string
}

type d5MapEntry struct {
	srcStart, dstStart, length int
}

func (e d5MapEntry) tryConvert(val int) (int, bool) {
	if val >= e.srcStart && val < e.srcStart+e.length {
		return val + e.dstStart - e.srcStart, true
	}

	return val, false
}

type d5Range struct {
	start, end int // end inclusive
}

func (e d5MapEntry) tryConvertRange(inp d5Range) (convertedRange *d5Range, unconvertedRange []d5Range) {
	il, ir := inp.start, inp.end
	ml, mr := e.srcStart, e.srcStart+e.length-1
	entryDelta := e.dstStart - e.srcStart

	if il < ml {
		if ir < ml {
			// no overlap
			return nil, []d5Range{inp}
		}

		// ir is at least equal to ml
		if ir <= mr {
			// partial overlap
			return &d5Range{
					start: ml + entryDelta,
					end:   ir + entryDelta,
				}, []d5Range{
					{start: il, end: ml - 1},
				}
		}

		// otherwise we have a rhs non overlap range
		return &d5Range{
				start: ml + entryDelta,
				end:   mr + entryDelta,
			}, []d5Range{
				{start: il, end: ml - 1},
				{start: mr + 1, end: ir},
			}
	}

	// otherwise il is at least equal to ml
	if il <= mr {
		// we at least have some overlap
		if ir <= mr {
			// full overlap
			return &d5Range{
				start: il + entryDelta,
				end:   ir + entryDelta,
			}, []d5Range{}
		}

		// otherwise we have a rhs non overlap range
		return &d5Range{
				start: il + entryDelta,
				end:   mr + entryDelta,
			}, []d5Range{
				{start: mr + 1, end: ir},
			}
	}

	// otherwise it's a complete non overlap
	return nil, []d5Range{inp}
}

func d5CollapseRanges(ranges []d5Range) []d5Range {
	// we'll need to sort the ranges first
	sort.Slice(ranges, func(a, b int) bool { return ranges[a].start < ranges[b].start })

	ret := make([]d5Range, 0, len(ranges))
	for _, r := range ranges {
		if len(ret) == 0 {
			ret = append(ret, r)
			continue
		}

		last := &ret[len(ret)-1]
		if last.end+1 >= r.start {
			// we can extend if we should
			if last.end < r.end {
				last.end = r.end
			}
		} else {
			ret = append(ret, r)
		}
	}

	return ret
}
