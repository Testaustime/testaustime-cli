package datahelper

import (
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
