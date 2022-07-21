package main

import (
	"github.com/romeq/testaustime-cli/apiEngine"
	"github.com/romeq/testaustime-cli/args"
	"github.com/romeq/testaustime-cli/config"
	"github.com/romeq/testaustime-cli/datahelper"
	"github.com/romeq/testaustime-cli/logger"
)

func main() {
	// parse arguments
	arguments := args.Parse()
	if arguments.DisableColors {
		logger.ColorsEnabled = false
	}

	// parse configuration file
	cfg := config.GetConfiguration(arguments.AlternateConfigFile)

	// apiengine
	api := apiEngine.New(cfg.Token, cfg.ApiUrl)
	switch arguments.Command {
	case "profile":
		userProfile := api.GetProfile()
		datahelper.ShowProfile(userProfile)

		break
	case "statistics":
		datahelper.ShowStatistics(api.GetStatistics())
		break
	default:
		args.Usage()
		return
	}
}
