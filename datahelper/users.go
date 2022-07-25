package datahelper

import (
	"time"

	"github.com/romeq/testaustime-cli/apiEngine"
)

func Show(key string, timeCoded float32, color int) {
	printField(key, rawTimeToHumanReadable(timeCoded), color)
}

func ShowStatistics(stats apiEngine.Statistics, showLatest bool) {
	printBold("Coding statistics")
	printField("Last 24h", rawTimeToHumanReadable(stats.Today), 32)
	printField("Last week", rawTimeToHumanReadable(stats.LastWeek), 37)
	printField("Last month", rawTimeToHumanReadable(stats.LastMonth), 37)
	printField("All time", rawTimeToHumanReadable(stats.AllTime), 37)

	if showLatest {
		printBold("\nLatest languages")
		showList(stats.TopLanguages)

		printBold("\nLatest projects")
		showList(stats.TopProjects)
	}
}

func ShowAccount(user apiEngine.User) {
	printField("Username", user.Username, 32)
	printField("Registration date", user.RegTime.Format(time.RFC1123), 37)
	printField("Friend code", "ttfc_"+user.FriendCode, 37)
}
