package main

import (
	"errors"
	"fmt"
	"time"

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
	logger.ColorsEnabled = !arguments.DisableColors
	apiengine.MeasureTime = arguments.MeasureRequests

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
				logger.Error(errors.New("You're already signed in on given account."))
				return
			}

			result, status := api.Login(username, *datahelper.AskPassword(""))
			if status.Err != "" {
				printErr(31, "Login failed", status.Err)
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

			password := datahelper.AskPassword("")
			result, status := api.Register(username, *password)
			zeroizePassowrds(password)
			if status.Err != "" {
				printErr(31, "Registration failed", status.Err)
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

			api.ChangePassword(*oldPassword, *newPassword)
			zeroizePassowrds(oldPassword, newPassword)

			utils.ColoredPrint(32, "Password was changed!\n")

		// User has massive a skill issue
		default:
			args.CommandUsage(args.AccountCommand)
		}

	case args.StatisticsCommand.Name:
		switch arguments.SubCommand {

		// User wants to see their current statistics
		case "":
			datahelper.ShowStatistics(api.GetStatistics("", false, time.Time{}), false)

		// User wants to also see their top projects and languages
		case args.StatisticsCommand.SubCommands["top"].Name:
			filterTime := time.Time{}
			switch utils.NthElement(arguments.OtherCommands, 2) {
			case "":
				break

			case "pastWeek":
				filterTime = time.Now().AddDate(0, 0, -7)

			case "pastMonth":
				filterTime = time.Now().AddDate(0, -1, 0)

			default:
				args.SubCommandUsage(
					args.StatisticsCommand,
					args.StatisticsCommand.SubCommands["top"].SubCommands,
				)
				return
			}
			datahelper.ShowStatistics(api.GetStatistics("", true, filterTime), true)

		default:
			args.CommandUsage(args.StatisticsCommand)
			return
		}

	case args.FriendsCommand.Name:
		switch arguments.SubCommand {
		case "":
			myaccount := api.GetStatistics("", false, time.Time{})
			friends := api.GetFriends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).AllTime())

		case args.FriendsCommand.SubCommands["pastWeek"].Name:
			myaccount := api.GetStatistics("", false, time.Now().AddDate(0, 0, -7))
			friends := api.GetFriends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).PastWeek())

		case args.FriendsCommand.SubCommands["pastMonth"].Name:
			myaccount := api.GetStatistics("", false, time.Now().AddDate(0, -1, 0))
			friends := api.GetFriends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).PastMonth())

		case args.FriendsCommand.SubCommands["add"].Name:
			var friendcode string
			if len(arguments.OtherCommands) < 3 {
				friendcode = datahelper.AskInput("Friend code")
			} else {
				friendcode = arguments.OtherCommands[2]
			}

			friend, err := api.AddFriend(friendcode)
			if err.Err != "" {
				printErr(31, "Friend left unadded", err.Err)
				return
			}
			utils.ColoredPrint(32, "Friend added!\n")
			datahelper.ShowFriend(friend)

		case args.FriendsCommand.SubCommands["remove"].Name:
			var friendcode string
			if len(arguments.OtherCommands) < 3 {
				friendcode = datahelper.AskInput("Friend name")
			} else {
				friendcode = arguments.OtherCommands[2]
			}

			api.RemoveFriend(friendcode)
			utils.ColoredPrint(32, "Friend removed!\n")

		default:
			args.CommandUsage(args.FriendsCommand)
		}

	case args.UserCommand.Name:
		if arguments.SubCommand == "" {
			args.CommandUsage(args.UserCommand)
			return
		}
		topCommand := args.UserCommand.SubCommands["<user>"].SubCommands["top"]

		switch utils.NthElement(arguments.OtherCommands, 2) {
		case "":
			datahelper.ShowStatistics(api.GetStatistics(arguments.SubCommand, false, time.Time{}), false)

		case topCommand.Name:
			filterTime := time.Time{}
			switch utils.NthElement(arguments.OtherCommands, 3) {
			case "":

			case topCommand.SubCommands["pastWeek"].Name:
				filterTime = time.Now().AddDate(0, 0, -7)

			case topCommand.SubCommands["pastMonth"].Name:
				filterTime = time.Now().AddDate(0, -1, 0)

			default:
				args.SubCommandUsage(
					args.UserCommand,
					topCommand.SubCommands,
				)
				return
			}
			datahelper.ShowStatistics(api.GetStatistics(arguments.SubCommand, true, filterTime), true)
		default:
			args.SubCommandUsage(
				args.UserCommand,
				args.UserCommand.SubCommands["<user>"].SubCommands,
			)
			return
		}

	default:
		args.Usage()
		return
	}
}

func printErr(color int, errtype, errmsg string) {
	utils.ColoredPrint(color, errtype)
	fmt.Println(":", errmsg)

}

func zeroizePassowrds[T *string](x ...T) {
	for _, p := range x {
		for i := 0; i < 128; i++ {
			*p += "\x00"
		}
	}
}
