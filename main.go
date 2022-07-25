package main

import (
	"fmt"

	"github.com/romeq/testaustime-cli/apiEngine"
	"github.com/romeq/testaustime-cli/args"
	"github.com/romeq/testaustime-cli/config"
	"github.com/romeq/testaustime-cli/datahelper"
	"github.com/romeq/testaustime-cli/logger"
	"github.com/romeq/testaustime-cli/utils"
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
	// account
	case args.Commands[0].Name:
		switch arguments.SubCommand {
		case "":
			userProfile := api.GetProfile()
			datahelper.ShowAccount(userProfile)

		case args.Commands[0].SubCommands[0].Name:
			utils.ColoredPrint(35, api.GetAuthtoken())

		case args.Commands[0].SubCommands[1].Name:
			token := api.NewAuthtoken()
			utils.ColoredPrint(35, token)
			cfg.UpdateField(&cfg.Token, token)

		case args.Commands[0].SubCommands[2].Name:
			token := api.NewFriendcode()
			utils.ColoredPrint(35, fmt.Sprintf("ttfc_%s", token))

		default:
			args.AccountUsage()
		}

	// statistics
	case args.Commands[1].Name:
		switch arguments.SubCommand {
		case "":
			datahelper.ShowStatistics(api.GetStatistics(""), false)

		case args.Commands[1].SubCommands[0].Name:
			datahelper.ShowStatistics(api.GetStatistics(""), true)

		default:
			args.StatisticsUsage()
		}

	case args.Commands[2].Name:
		myaccount := api.GetStatistics("")
		friends := api.GetFriends()
		switch arguments.SubCommand {
		case "":
			datahelper.ShowFriends(friends.AddSelf(myaccount).AllTime())
		case args.Commands[2].SubCommands[0].Name:
			datahelper.ShowFriends(friends.AddSelf(myaccount).LastMonth())
		case args.Commands[2].SubCommands[1].Name:
			datahelper.ShowFriends(friends.AddSelf(myaccount).LastWeek())
		default:
			args.FriendsUsage()
		}

	default:
		args.Usage()
		return
	}
}
