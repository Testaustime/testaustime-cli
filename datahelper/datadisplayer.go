package datahelper

import (
	"time"

	"github.com/romeq/testaustime-cli/apiengine"
)

func Show(key string, timeCoded float32, color int) {
	printField(key, rawTimeToHumanReadable(timeCoded), color)
}

func ShowStatistics(stats apiengine.Statistics, showTop bool, highlighted int) {
    notActiveColor := 37
    color := notActiveColor
    var i int
    checkColor := func(currentIndex int) {
        currentIndex++
        if currentIndex == highlighted {
            color = 32
        } else if currentIndex > highlighted {
            color = notActiveColor
        }
        i = currentIndex
    }

	printBold("Coding statistics")

    checkColor(i)
	printField("All time", rawTimeToHumanReadable(stats.AllTime), color)

    checkColor(i)
	printField("Past 24h", rawTimeToHumanReadable(stats.Today), color)

    checkColor(i)
	printField("Past week", rawTimeToHumanReadable(stats.PastWeek), color)

    checkColor(i)
	printField("Past month", rawTimeToHumanReadable(stats.PastMonth), color)


	if showTop {
		showList("Top languages", stats.TopLanguages)
		showList("Top projects", stats.TopProjects)
		showList("Top hosts", stats.TopHosts)
		showList("Top editors", stats.TopEditors)
	}
}

func ShowAccount(user apiengine.User) {
	printField("Username", user.Username, 32)
	printField("Registration date", user.RegTime.Format(time.RFC1123), 37)
	printField("Friend code", "ttfc_"+user.FriendCode, 37)
}
