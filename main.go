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
	case args.AccountCommand.Name:
		switch arguments.SubCommand {
		case "":
			userProfile := api.GetProfile()
			datahelper.ShowAccount(userProfile)

		case args.AccountCommand.SubCommands[0].Name:
			utils.ColoredPrint(35, api.GetAuthtoken())

		case args.AccountCommand.SubCommands[1].Name:
			token := api.NewAuthtoken()
			utils.ColoredPrint(35, token)
			cfg.UpdateField(&cfg.Token, token)

		case args.AccountCommand.SubCommands[2].Name:
			token := api.NewFriendcode()
			utils.ColoredPrint(35, fmt.Sprintf("ttfc_%s", token))

		default:
			args.CommandUsage(args.AccountCommand)
		}

	// statistics
	case args.StatisticsCommand.Name:
		switch arguments.SubCommand {
		case "":
			datahelper.ShowStatistics(api.GetStatistics(""), false)

		case args.StatisticsCommand.SubCommands[0].Name:
			datahelper.ShowStatistics(api.GetStatistics(""), true)

		default:
			args.CommandUsage(args.StatisticsCommand)
		}

	case args.FriendsCommand.Name:
		switch arguments.SubCommand {
		case "":
			myaccount := api.GetStatistics("")
			friends := api.GetFriends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).AllTime())
		case args.FriendsCommand.SubCommands[0].Name:
			myaccount := api.GetStatistics("")
			friends := api.GetFriends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).LastMonth())
		case args.FriendsCommand.SubCommands[1].Name:
			myaccount := api.GetStatistics("")
			friends := api.GetFriends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).LastWeek())
		default:
			args.CommandUsage(args.FriendsCommand)
		}

	case args.UserCommand.Name:
		if arguments.SubCommand == "" {
			args.CommandUsage(args.UserCommand)
			return
		}

		datahelper.ShowStatistics(api.GetStatistics(arguments.SubCommand), false)

	default:
		args.Usage()
		return
	}
}
