package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

func loadLines(d int) []string {
	fn := fmt.Sprintf("data/d%d.txt", d)
	content, err := os.ReadFile(fn)
	c(err)

	return sanitizeInput(string(content))
}

func sanitizeInput(in string) []string {
	lines := strings.Split(string(in), "\n")
	lines = lo.FilterMap(lines, func(s string, _ int) (string, bool) {
		s = strings.TrimSpace(s)
		return s, s != ""
	})

	return lines
}
