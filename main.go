package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/romeq/testaustime-cli/apiengine"
	"github.com/romeq/testaustime-cli/arguments"
	"github.com/romeq/testaustime-cli/config"
	"github.com/romeq/testaustime-cli/datahelper"
	"github.com/romeq/testaustime-cli/logger"
	"github.com/romeq/testaustime-cli/utils"
)

func main() {
	// parse args
	args := arguments.Parse()
	logger.ColorsEnabled = !args.DisableColors
	apiengine.MeasureTime = args.MeasureRequests

	tokenLocation := utils.ResolveWantedPath(utils.EnvOrString(
		"$XDG_DATA_HOME",
		"$HOME/.local/share",
	), "token")

	token := utils.ReadFile(tokenLocation, true)
	cfg := config.GetConfiguration(args.AlternateConfigFile)
	api := apiengine.New(token, cfg.ApiUrl, cfg.CaseInsensitiveFields)

	switch args.Command {
	case arguments.AccountCommand.Name:
		switch args.SubCommand {

		// User wants to see his account's information
		case "":
			userProfile := api.Profile()
			datahelper.ShowAccount(userProfile)

			// User wants to login
		case arguments.AccountCommand.SubCommands["login"].Name:
			username := nthElOrInput("Username", args.OtherCommands, 3)

			password := datahelper.AskPassword("")
			result, status := api.Login(username, *password)
			zeroizePasswords(password)
			if status.Err != "" {
				printErr(31, "Login failed", status.Err)
				break
			}

			utils.WriteFile(tokenLocation, result.Token)
			utils.ColoredPrint(32, "Login succeeded and credinteals were saved!\n")

		// User wants to register a new account
		case arguments.AccountCommand.SubCommands["register"].Name:
			username := nthElOrInput("New username", args.OtherCommands, 3)

			password := datahelper.AskPassword("")
			result, status := api.Register(username, *password)
			zeroizePasswords(password)
			if status.Err != "" {
				printErr(31, "Registration failed", status.Err)
				break
			}

			utils.WriteFile(tokenLocation, result.Token)
			utils.ColoredPrint(32, "Registration succeeded and credinteals were saved!\n")

		// User queries their current authentication token
		case arguments.AccountCommand.SubCommands["token"].Name:
			utils.ColoredPrint(35, fmt.Sprintf("%s\n", api.GetAuthtoken()))

		// User wants to generate a new authentication token
		case arguments.AccountCommand.SubCommands["newToken"].Name:
			token := api.NewAuthtoken()
			utils.ColoredPrint(35, fmt.Sprintf("%s\n", token))
			utils.WriteFile(tokenLocation, token)

		// User wants to generate a new friend code
		case arguments.AccountCommand.SubCommands["newFriendcode"].Name:
			token := api.NewFriendcode()
			utils.ColoredPrint(35, fmt.Sprintf("ttfc_%s\n", token))

		// User wants to change password
		case arguments.AccountCommand.SubCommands["changePassword"].Name:
			oldPassword := datahelper.AskPassword("Old password")
			newPassword := datahelper.AskPassword("New password")

			api.ChangePassword(*oldPassword, *newPassword)
			zeroizePasswords(oldPassword, newPassword)

			utils.ColoredPrint(32, "Password was changed!\n")

		// User has massive a skill issue
		default:
			arguments.CommandUsage(arguments.AccountCommand)
		}

	case arguments.LeaderboardCommand.Name:
		switch args.SubCommand {
		case "":
			leaderboards := api.Leaderboards()
			datahelper.ShowLeaderboards(leaderboards)

		case arguments.LeaderboardCommand.SubCommands["join"].Name:
			code := nthElOrInput("Leaderboard code", args.OtherCommands, 3)
			api.JoinLeaderboard(code)

		case arguments.LeaderboardCommand.SubCommands["create"].Name:
			name := nthElOrInput("Leaderboard name", args.OtherCommands, 3)
			invite := api.NewLeaderboard(name)
			utils.ColoredPrint(35, fmt.Sprintf("ttfic_%s\n", invite.Code))

		case arguments.LeaderboardCommand.SubCommands["delete"].Name:
			utils.ColoredPrint(31, "Hey you! ")
			fmt.Println("This process cannot be reversed. Make sure there is nothing to lose before deleting.")
			name := nthElOrInput("Leaderboard name", args.OtherCommands, 3)
			nameConfirm := datahelper.AskInput("Confirm leaderboard name to delete")
			if name != nameConfirm {
				logger.Error(errors.New("The names don't match!"))
			}

			api.DeleteLeaderboard(name)
			utils.ColoredPrint(32, "Leaderboard deleted!\n")

		case arguments.LeaderboardCommand.SubCommands["leave"].Name:
			api.LeaveLeaderboard(nthElOrInput("Leaderboard name", args.OtherCommands, 3))

		case arguments.LeaderboardCommand.SubCommands["regenerate"].Name:
			invitecode := api.RegenerateLeaderboardInvite(nthElOrInput(
				"Leaderboard name",
				args.OtherCommands,
				3,
			))
			utils.ColoredPrint(35, fmt.Sprintf("ttfic_%s\n", invitecode.Code))

		case arguments.LeaderboardCommand.SubCommands["kick"].Name:
			ldname := nthElOrInput("Leaderboard name", args.OtherCommands, 3)
			uname := nthElOrInput("Username to kick", args.OtherCommands, 4)
			api.KickMember(ldname, uname)
			utils.ColoredPrint(32, "The member has been kicked!\n")

		default:
			leaderboard := api.Leaderboard(nthElOrInput("Leaderboard name", args.OtherCommands, 2))
			userAccount := api.Profile()
			datahelper.ShowLeaderboard(leaderboard, userAccount.Username)
		}

	case arguments.StatisticsCommand.Name:
		switch args.SubCommand {

		// User wants to see their current statistics
		case "":
			datahelper.ShowStatistics(api.Statistics("", false, time.Time{}), false, 1)

		// User wants to also see their top projects and languages
		case arguments.StatisticsCommand.SubCommands["top"].Name:
			filterTime := time.Time{}
			activeField := 1
			switch utils.NthElement(args.OtherCommands, 2) {
			case "":
				break

			case "pastWeek":
				activeField = 3
				filterTime = datahelper.Dates.PastWeek

			case "pastMonth":
				activeField = 4
				filterTime = datahelper.Dates.PastMonth

			default:
				arguments.SubCommandUsage(
					arguments.StatisticsCommand,
					arguments.StatisticsCommand.SubCommands["top"],
				)
				return
			}
			datahelper.ShowStatistics(api.Statistics(
				"",
				true,
				filterTime,
			), true, activeField)

		default:
			arguments.CommandUsage(arguments.StatisticsCommand)
			return
		}

	case arguments.FriendsCommand.Name:
		switch args.SubCommand {
		case "":
			myaccount := api.Statistics("", false, time.Time{})
			friends := api.Friends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).AllTime())

		case arguments.FriendsCommand.SubCommands["pastWeek"].Name:
			myaccount := api.Statistics("", false, time.Now().AddDate(0, 0, -7))
			friends := api.Friends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).PastWeek())

		case arguments.FriendsCommand.SubCommands["pastMonth"].Name:
			myaccount := api.Statistics("", false, time.Now().AddDate(0, -1, 0))
			friends := api.Friends()
			datahelper.ShowFriends(friends.AddSelf(myaccount).PastMonth())

		case arguments.FriendsCommand.SubCommands["add"].Name:
			friendcode := nthElOrInput("Friend code", args.OtherCommands, 2)
			friend, err := api.AddFriend(friendcode)
			if err.Err != "" {
				printErr(31, "Friend left unadded", err.Err)
				return
			}
			utils.ColoredPrint(32, "Friend added!\n")
			datahelper.ShowFriend(friend)

		case arguments.FriendsCommand.SubCommands["remove"].Name:
			friendName := nthElOrInput("Friend's username", args.OtherCommands, 2)
			api.RemoveFriend(friendName)
			utils.ColoredPrint(33, "Friend removed!\n")

		default:
			arguments.CommandUsage(arguments.FriendsCommand)
		}

	case arguments.UserCommand.Name:
		if args.SubCommand == "" {
			arguments.CommandUsage(arguments.UserCommand)
			return
		}
		topCommand := arguments.UserCommand.SubCommands["<user>"].SubCommands["top"]

		switch utils.NthElement(args.OtherCommands, 2) {
		case "":
			datahelper.ShowStatistics(api.Statistics(
				args.SubCommand,
				false,
				time.Time{},
			), false, 0)

		case topCommand.Name:
			var filterTime time.Time
			activeField := 1
			switch utils.NthElement(args.OtherCommands, 3) {
			case "":
				filterTime = time.Time{}

			case topCommand.SubCommands["pastWeek"].Name:
				activeField = 3
				filterTime = datahelper.Dates.PastWeek

			case topCommand.SubCommands["pastMonth"].Name:
				activeField = 4
				filterTime = datahelper.Dates.PastMonth

			default:
				arguments.SubCommandUsage(
					arguments.UserCommand,
					topCommand,
				)
				return
			}
			datahelper.ShowStatistics(api.Statistics(
				args.SubCommand,
				true,
				filterTime,
			), true, activeField)

		default:
			arguments.SubCommandUsage(
				arguments.UserCommand,
				arguments.UserCommand.SubCommands["<user>"],
			)
			return
		}

	default:
		arguments.Usage()
		return
	}
}

func printErr(color int, errtype, errmsg string) {
	utils.ColoredPrint(color, errtype)
	fmt.Println(":", errmsg)
}

func zeroizePasswords[T *string](passwords ...T) {
	blk := make([]byte, 128)
	for _, password := range passwords {
		_, err := rand.Read(blk)
		utils.Check(err)
		*password = string(blk)
	}
}

func nthElOrInput(prompt string, bla []string, n int) string {
	if arg := utils.NthElement(bla, n-1); arg != "" {
		return arg
	}
	return datahelper.AskInput(prompt)
}
