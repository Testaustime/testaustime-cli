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

func ShowLeaderboard(leaderboard apiengine.Leaderboard, highlightedUsername string) {
	printField("Name", leaderboard.Name, 32)
	printField("Invite", fmt.Sprintf("ttlic_%s", leaderboard.Invite), 37)
	printField("Creation time", leaderboard.CreationTime, 37)

	if len(leaderboard.Members) == 0 {
		utils.ColoredPrint(33, "There is no members on this leaderboard.\n")
		return
	}

	utils.ColoredPrint(32, "\nMembers on this leaderboard\n")
	for _, user := range leaderboard.SortMembersByTime() {
		color := 37
		if user.Admin && user.Username == highlightedUsername {
			color = 33
        } else if user.Username == highlightedUsername {
			color = 32
		} else if user.Admin {
			color = 31
		}
		printField(user.Username, rawTimeToHumanReadable(float32(user.TimeCoded)/60), color)
	}
}
