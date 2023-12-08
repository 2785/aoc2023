package main

import (
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var l *zap.Logger
var s *zap.SugaredLogger

func init() {
	var err error
	l, err = zap.NewDevelopment(zap.WithCaller(false))
	c(err)

	zap.ReplaceGlobals(l)

	s = l.Sugar()
}

var start time.Time

var root = &cobra.Command{
	Use: "aoc2023",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		start = time.Now()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		s.Infof("took %s", time.Since(start))
	},
}

func c(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	c(root.Execute())
}
