package args

import (
	"flag"
	"fmt"
)

var colors = true
var Commands = []Command{
	{
		Name: "account",
		Info: "show account information",
		SubCommands: []SubCommand{
			{
				Name: "token",
				Info: "shows your authtoken",
			},
			{
				Name: "newToken",
				Info: "regenerates your authorization token",
			},
			{
				Name: "newFriendcode",
				Info: "regenerates your friend code",
			},
		},
	},
	{
		Name: "statistics",
		Info: "show coding statistics",
		SubCommands: []SubCommand{
			{
				Name: "latest",
				Info: "show latest languages and projects",
			},
		},
	},
	{
		Name: "friends",
		Info: "show friends' statistics",
		SubCommands: []SubCommand{
			{
				Name: "pastMonth",
				Info: "show friends' coding time during past month",
			},
			{
				Name: "pastWeek",
				Info: "show friends' coding time during past week",
			},
		},
	},
	{
		Name:        "user",
		Info:        "show specific friend's statistics",
		SubCommands: []SubCommand{},
	},
}

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
