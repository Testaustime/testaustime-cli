package datahelper

import "github.com/romeq/testaustime-cli/apiengine"

func ShowSummary(summary apiengine.SummaryResponse) {
	printField("All time", rawTimeToHumanReadable(float32(summary.AllTime.Total)), 32)
	printField("Last month", rawTimeToHumanReadable(float32(summary.LastMonth.Total)), 37)
	printField("Last week", rawTimeToHumanReadable(float32(summary.LastWeek.Total)), 37)

	showLanguages("All time", summary.AllTime.Languages)
	showLanguages("Last month", summary.LastMonth.Languages)
	showLanguages("Last week", summary.LastWeek.Languages)
}

func showLanguages(blaa string, languages map[string]int) {
	if len(languages) == 0 {
		return
	}

	printBold("\n", blaa)
	for lang, time := range languages {
		printField(lang, time, 37)
	}
}
