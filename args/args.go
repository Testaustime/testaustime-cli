package args

import (
	"flag"
	"fmt"
)

type Args struct {
	Command             string
	DisableColors       bool
	AlternateConfigFile string
}

func Parse() (args Args) {
	flag.Usage = Usage

	flag.StringVar(&args.AlternateConfigFile, "c", "", "Set alternate config location")
	flag.BoolVar(&args.DisableColors, "no-colors", false, "Disable colors in output")
	flag.Parse()

	args.Command = flag.Arg(0)
	return args
}

func Usage() {
	fmt.Println(
		"usage: ./testaustime-cli [flags] [command] \n",
		"\n",
		"flags: \n",
		"  -c FILE: read configuration from FILE \n",
		"  -no-colors: don't include colors in output \n",
		"  -h, --help: show help menu \n",
		"\n",
		"commands: \n",
		"  profile: show account information \n",
		"  statistics: show coding statistics",
	)
}
