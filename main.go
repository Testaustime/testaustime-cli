package main

import (
	"errors"
	"fmt"

	"github.com/romeq/testaustime-cli/apiengine"
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
	api := apiengine.New(cfg.Token, cfg.ApiUrl)

	switch arguments.Command {
	// account
	case args.AccountCommand.Name:
		switch arguments.SubCommand {

		// User wants to see his account information
		case "":
			userProfile := api.GetProfile()
			datahelper.ShowAccount(userProfile)

			// User wants to login
		case args.AccountCommand.SubCommands["login"].Name:
			var username string
			if len(arguments.OtherCommands) < 3 {
				username = datahelper.AskInput("Username")
			} else {
				username = arguments.OtherCommands[2]
			}
			if username == api.GetProfile().Username {
				logger.Error(errors.New(fmt.Sprintf("You're already signed in on given account.")))
				return
			}

			result, status := api.Login(username, datahelper.AskPassword(""))
			if status.Err != "" {
				utils.ColoredPrint(31, "Login failed")
				fmt.Println(":", status.Err)
				break
			}

			utils.ColoredPrint(32, "Login succeeded and credinteals were saved!\n")
			cfg.UpdateField(&cfg.Token, result.Token)

			// User wants to register a new account
		case args.AccountCommand.SubCommands["register"].Name:
			var username string
			if len(arguments.OtherCommands) < 3 {
				username = datahelper.AskInput("New username")
			} else {
				username = arguments.OtherCommands[2]
			}

			result, status := api.Register(username, datahelper.AskPassword(""))
			if status.Err != "" {
				utils.ColoredPrint(31, "Registration failed")
				fmt.Println(":", status.Err)
				break
			}

			utils.ColoredPrint(32, "Registration succeeded and credinteals were saved!\n")
			cfg.UpdateField(&cfg.Token, result.Token)

			// User queries their current authentication token
		case args.AccountCommand.SubCommands["token"].Name:
			utils.ColoredPrint(35, fmt.Sprintf("%s\n", api.GetAuthtoken()))

			// User wants to generate a new authentication token
		case args.AccountCommand.SubCommands["newToken"].Name:
			token := api.NewAuthtoken()
			utils.ColoredPrint(35, fmt.Sprintf("%s\n", token))
			cfg.UpdateField(&cfg.Token, token)

			// User wants to generate a new friend code
		case args.AccountCommand.SubCommands["newFriendcode"].Name:
			token := api.NewFriendcode()
			utils.ColoredPrint(35, fmt.Sprintf("ttfc_%s\n", token))

			// User wants to change password
		case args.AccountCommand.SubCommands["changePassword"].Name:
			oldPassword := datahelper.AskPassword("Old password")
			newPassword := datahelper.AskPassword("New password")
			err := api.ChangePassword(oldPassword, newPassword)
			if err.Err != "" {
				utils.ColoredPrint(31, "Password was not changed")
				fmt.Println(":", err.Err)
				return
			}
			utils.ColoredPrint(32, "Password was changed!\n")

			// User has massive a skill issue
		default:
			args.CommandUsage(args.AccountCommand)
		}

	case args.StatisticsCommand.Name:
		switch arguments.SubCommand {

		// User wants to see their current statistics
		case "":
			datahelper.ShowStatistics(api.GetStatistics(""), false)

			// User wants to also see their latest projects and languages
		case args.StatisticsCommand.SubCommands["latest"].Name:
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

		case args.FriendsCommand.SubCommands["pastWeek"].Name:
			myaccount := api.GetStatistics("")
			friends := api.GetFriends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).PastWeek())

		case args.FriendsCommand.SubCommands["pastMonth"].Name:
			myaccount := api.GetStatistics("")
			friends := api.GetFriends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).PastMonth())

		case args.FriendsCommand.SubCommands["add"].Name:
			var friendcode string
			if len(arguments.OtherCommands) < 3 {
				friendcode = datahelper.AskInput("Friend code")
			} else {
				friendcode = arguments.OtherCommands[2]
			}

			err := api.AddFriend(friendcode)
			if err.Err != "" {
				utils.ColoredPrint(31, "Friend left unadded")
				fmt.Println(":", err.Err)
				return
			}
			utils.ColoredPrint(32, "Friend added!\n")

		case args.FriendsCommand.SubCommands["remove"].Name:
			var friendcode string
			if len(arguments.OtherCommands) < 3 {
				friendcode = datahelper.AskInput("Friend name")
			} else {
				friendcode = arguments.OtherCommands[2]
			}

			err := api.RemoveFriend(friendcode)
			if err.Err != "" {
				utils.ColoredPrint(31, "Friend could not be removed")
				fmt.Println(":", err.Err)
				return
			}
			utils.ColoredPrint(32, "Friend removed!\n")

		default:
			args.CommandUsage(args.FriendsCommand)
		}

	case args.UserCommand.Name:
		if arguments.SubCommand == "" {
			args.UserUsage()
			return
		}

		datahelper.ShowStatistics(api.GetStatistics(arguments.SubCommand), false)

	default:
		args.Usage()
		return
	}
}
