package main

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var l *zap.Logger
var s *zap.SugaredLogger

func init() {
	var err error
	l, err = zap.NewProduction()
	c(err)

	zap.ReplaceGlobals(l)

	s = l.Sugar()
}

var root = &cobra.Command{
	Use: "aoc2023",
}

func c(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	c(root.Execute())
}
