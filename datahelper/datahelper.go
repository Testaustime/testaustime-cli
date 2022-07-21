package datahelper

import (
	"fmt"
	"time"

	"github.com/romeq/testaustime-cli/apiEngine"
	"github.com/romeq/testaustime-cli/logger"
)

func ShowProfile(user apiEngine.User) {
	printField("Username", user.Username, 32)
	printField("Registration date", user.RegTime.Format(time.RFC1123), 37)
	printField("Friend code", "ttfc_"+user.FriendCode, 37)
}

func ShowStatistics(stats apiEngine.Statistics) {
	printBold("Coding statistics")
	printField("Today", rawTimeToHumanReadable(stats.Today), 32)
	printField("Last week", rawTimeToHumanReadable(stats.LastWeek), 37)
	printField("Last month", rawTimeToHumanReadable(stats.LastMonth), 37)
	printField("All time", rawTimeToHumanReadable(stats.AllTime), 37)

	printBold("\nTop 10 languages")
	showList(stats.TopLanguages)

	printBold("\nTop 10 projects")
	showList(stats.TopProjects)
}

func showList(list apiEngine.TopStatsList) {
	for i, item := range list {
		if i >= 10 {
			break
		}

		if item.Name == "" {
			item.Name = "<none>"
		}
		color := 37
		if i == 0 {
			color = 32
		}

		printField(item.Name, rawTimeToHumanReadable(item.Time), color)
	}
}
func rawTimeToHumanReadable(minutesCoded float32) string {
	var daysCoded, hoursCoded int
	remainderMinutes := int(minutesCoded)

	for remainderMinutes >= 60 {
		if remainderMinutes >= 60*24 {
			remainderMinutes -= 60 * 24
			daysCoded += 1
			continue
		}
		remainderMinutes -= 60
		hoursCoded += 1
	}

	if daysCoded == 0 && hoursCoded > 0 {
		return fmt.Sprintf("%dh, %dm", hoursCoded, remainderMinutes)
	} else if daysCoded == 0 && hoursCoded == 0 {
		return fmt.Sprintf("%dm", remainderMinutes)
	}
	return fmt.Sprintf("%dd, %dh, %dm", daysCoded, hoursCoded, remainderMinutes)
}

func printField(key string, value any, color int) {
	if logger.ColorsEnabled {
		fmt.Println(fmt.Sprintf("\033[%dm%s\033[0m: %s", color, key, value))
	} else {
		fmt.Println(fmt.Sprintf("%s: %s", key, value))
	}
}

func printBold(a ...any) {
	printColored(1, a...)
}

func printColored(color int, a ...any) {
	if logger.ColorsEnabled {
		fmt.Printf("\033[%dm", color)
		fmt.Print(a...)
		fmt.Printf("\033[0m\n")
		return
	}
	fmt.Print(a...)
	fmt.Printf("\n")
}
