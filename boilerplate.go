package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"text/template"

	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(newDay)
}

var ft = `package main

import "github.com/spf13/cobra"

func init() {
	root.AddCommand(d{{.}}p1)
	root.AddCommand(d{{.}}p2)
}

var d{{.}}p1 = &cobra.Command{
	Use: "d{{.}}p1",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var d{{.}}p2 = &cobra.Command{
	Use: "d{{.}}p2",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
`

var tt = `package main

import "testing"

func TestD{{.}}P1(t *testing.T) {
	
}

func TestD{{.}}P2(t *testing.T) {
	
}
`

var newDay = &cobra.Command{
	Use: "new",
	Run: func(cmd *cobra.Command, args []string) {
		// look at what days we already have

		// list curr dir
		files, err := filepath.Glob("d*.go")
		c(err)

		re := regexp.MustCompile(`d(\d+).go`)

		// find the highest number
		highest := 0
		for _, f := range files {
			m := re.FindStringSubmatch(f)
			if m == nil {
				continue
			}

			if len(m) != 2 {
				continue
			}

			n, err := strconv.Atoi(m[1])
			c(err)
			if n > highest {
				highest = n
			}
		}

		next := highest + 1

		// create the files
		mainTemplate, err := template.New("main").Parse(ft)
		c(err)

		mainFile, err := os.Create(fmt.Sprintf("d%d.go", next))
		c(err)

		err = mainTemplate.Execute(mainFile, next)
		c(err)

		testTemplate, err := template.New("test").Parse(tt)
		c(err)

		testFile, err := os.Create(fmt.Sprintf("d%d_test.go", next))
		c(err)

		err = testTemplate.Execute(testFile, next)
		c(err)

		// make the data file
		_, err = os.Create(fmt.Sprintf("data/d%d.txt", next))
		c(err)
	},
}
