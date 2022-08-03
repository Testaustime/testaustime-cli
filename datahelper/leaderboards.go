package datahelper

import (
	"fmt"

	"github.com/romeq/testaustime-cli/apiengine"
	"github.com/romeq/testaustime-cli/utils"
)

func ShowLeaderboards(leaderboards apiengine.LeaderboardList) {
	if len(leaderboards) == 0 {
		utils.ColoredPrint(33, "You haven't joined any leaderboards!\n")
		return
	}

	for _, leaderboard := range leaderboards {
		printField(leaderboard.Name, leaderboard.MemberCount, 37)
	}
}

func ShowLeaderboard(leaderboard apiengine.Leaderboard, username string) {
	printField("Leaderboard name", leaderboard.Name, 32)
	printField("Invite code", fmt.Sprintf("ttlic_%s", leaderboard.Invite), 37)
	printField("Creation time", leaderboard.CreationTime, 37)

	if len(leaderboard.Members) == 0 {
		utils.ColoredPrint(33, "There is no members on this leaderboard.\n")
		return
	}

    rank := 0
    sortedmembers := leaderboard.SortMembersByTime()
    for i, x := range sortedmembers {
        if x.Username == username {
            rank = i+1
            break
        }
    }

	printField("Your rank", prettyPrintRank(rank), 37)
	utils.ColoredPrint(32, "\nMembers on this leaderboard\n")
	for i, user := range sortedmembers {
        if i == 50 {
            break
        }

		color := 37
		if user.Admin && user.Username == username {
			color = 33
		} else if user.Username == username {
			color = 32
		} else if user.Admin {
			color = 31
		}
		printField(user.Username, rawTimeToHumanReadable(float32(user.TimeCoded)/60), color)
	}
}

func prettyPrintRank(rank int) string {
    if rank == 1 {
        return "1st! 🥇"
    } else if rank == 2 {
        return "2nd 🥈"
    } else if rank == 3 {
        return "3rd 🥉"
    }
    return fmt.Sprintf("%dth", rank)
}

