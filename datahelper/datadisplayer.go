package datahelper

import (
	"time"

	"github.com/romeq/testaustime-cli/apiengine"
)

func Show(key string, timeCoded float32, color int) {
	printField(key, rawTimeToHumanReadable(timeCoded), color)
}

func ShowStatistics(stats apiengine.Statistics, showTop bool) {
	printBold("Coding statistics")
	printField("Past 24h", rawTimeToHumanReadable(stats.Today), 32)
	printField("Past week", rawTimeToHumanReadable(stats.PastWeek), 37)
	printField("Past month", rawTimeToHumanReadable(stats.PastMonth), 37)
	printField("All time", rawTimeToHumanReadable(stats.AllTime), 37)

	if showTop {
		printBold("\nTop languages")
		showList(stats.TopLanguages)

		printBold("\nTop projects")
		showList(stats.TopProjects)

		printBold("\nTop hosts")
		showList(stats.TopHosts)

		printBold("\nTop editors")
		showList(stats.TopEditors)
	}
}

func ShowAccount(user apiengine.User) {
	printField("Username", user.Username, 32)
	printField("Registration date", user.RegTime.Format(time.RFC1123), 37)
	printField("Friend code", "ttfc_"+user.FriendCode, 37)
}
