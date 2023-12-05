package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func loadFileAsString(d int) string {
	fn := fmt.Sprintf("data/d%d.txt", d)
	content, err := os.ReadFile(fn)
	c(err)

	return string(content)
}

func loadLines(d int) []string {
	return sanitizeInput(loadFileAsString(d))
}

func sanitizeInput(in string) []string {
	return sp(in, "\n")
}

func sp(in string, sep string) []string {
	lines := strings.Split(in, sep)
	lines = lo.FilterMap(lines, func(s string, _ int) (string, bool) {
		s = strings.TrimSpace(s)
		return s, s != ""
	})

	return lines
}

func reqLen[T any](thing []T, req int) {
	if len(thing) != req {
		panic(fmt.Sprintf("expected %d, got %d", req, len(thing)))
	}
}

func pow(base, exp int) int {
	res := 1

	for i := 0; i < exp; i++ {
		res *= base
	}

	return res
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	c(err)

	return n
}

func mustTrimPrefix(s, prefix string) string {
	res := strings.TrimPrefix(s, prefix)
	if res == s {
		panic(fmt.Sprintf("expected %q to have prefix %q", s, prefix))
	}

	return strings.TrimSpace(res)
}
