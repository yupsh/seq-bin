package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/seq"
)

const (
	flagSeparator  = "separator"
	flagFormat     = "format"
	flagEqualWidth = "equal-width"
)

func main() {
	app := &cli.App{
		Name:  "seq",
		Usage: "print a sequence of numbers",
		UsageText: `seq [OPTIONS] LAST
   seq [OPTIONS] FIRST LAST
   seq [OPTIONS] FIRST INCREMENT LAST

   Print numbers from FIRST to LAST, in steps of INCREMENT.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagSeparator,
				Aliases: []string{"s"},
				Usage:   "use STRING to separate numbers (default: \\n)",
			},
			&cli.StringFlag{
				Name:    flagFormat,
				Aliases: []string{"f"},
				Usage:   "use printf style floating-point FORMAT",
			},
			&cli.BoolFlag{
				Name:    flagEqualWidth,
				Aliases: []string{"w"},
				Usage:   "equalize width by padding with leading zeroes",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "seq: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add all arguments - convert strings to float64
	for i := 0; i < c.NArg(); i++ {
		val, err := strconv.ParseFloat(c.Args().Get(i), 64)
		if err != nil {
			return fmt.Errorf("invalid number: %s", c.Args().Get(i))
		}
		params = append(params, val)
	}

	// Add flags based on CLI options
	if c.IsSet(flagSeparator) {
		params = append(params, Separator(c.String(flagSeparator)))
	}
	if c.IsSet(flagFormat) {
		params = append(params, Format(c.String(flagFormat)))
	}
	if c.Bool(flagEqualWidth) {
		params = append(params, EqualWidth)
	}

	// Create and execute the seq command
	cmd := Seq(params...)
	return gloo.Run(cmd)
}
