package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(d1p1)
	root.AddCommand(d1p2)
}

var d1p1 = &cobra.Command{
	Use: "d1p1",
	Run: func(cmd *cobra.Command, args []string) {
		input := loadLines(1)
		s.Infof("sum: %d", doD1P1(input))
	},
}

var d1p2 = &cobra.Command{
	Use: "d1p2",
	Run: func(cmd *cobra.Command, args []string) {
		input := loadLines(1)
		s.Infof("sum: %d", doD1P2(input))
	},
}

func doD1P1(input []string) int {
	var sum int

	for _, line := range input {
		var fd, ld string

		lineSplit := strings.Split(line, "")
		for _, char := range lineSplit {
			if _, err := strconv.Atoi(char); err == nil {
				if fd == "" {
					fd = char
				}

				ld = char
			}
		}

		num, err := strconv.Atoi(fd + ld)
		c(err)

		sum += num
	}

	return sum
}

func doD1P2(input []string) int {
	var sum int

	for _, line := range input {
		var nums []int

		num, ok, left := parseChunk(line)
		for ok {
			nums = append(nums, num)
			num, ok, left = parseChunk(left)
		}

		if len(nums) == 0 {
			c(fmt.Errorf("no numbers found in line: %s", line))
		}

		sum += nums[0]*10 + nums[len(nums)-1]
	}

	return sum
}

var wordMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parseChunk(chunk string) (num int, ok bool, left string) {
	// if first thing in the chunk is an actual number, we're done
	if len(chunk) == 0 {
		return 0, false, ""
	}

	char1 := chunk[0]
	if num, err := strconv.Atoi(string(char1)); err == nil {
		return num, true, chunk[1:]
	}

	// otherwise if it starts with any number text spelled out, we parse that & we're done
	for word, num := range wordMap {
		if strings.HasPrefix(chunk, word) {
			return num, true, chunk[1:]
		}
	}

	// otherwise we chop off the first letter & try again
	return parseChunk(chunk[1:])
}
