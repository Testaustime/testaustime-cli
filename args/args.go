package args

import (
	"flag"
	"fmt"
)

var colors = true

type Args struct {
	Command             string
	SubCommand          string
	DisableColors       bool
	AlternateConfigFile string
}

func Parse() (args Args) {
	flag.Usage = Usage

	flag.StringVar(&args.AlternateConfigFile, "c", "", "Set alternate config location")
	flag.BoolVar(&args.DisableColors, "no-colors", false, "Disable colors in output")
	flag.Parse()

	colors = !args.DisableColors
	args.Command = flag.Arg(0)
	args.SubCommand = flag.Arg(1)

	return args
}

func coloredString(str string, color int) string {
	if colors {
		return fmt.Sprintf("\033[%dm%s\033[0m", color, str)
	}
	return str
}
