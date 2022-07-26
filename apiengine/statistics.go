package apiengine

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/romeq/testaustime-cli/utils"
)

// response from testaustime API is parsed with this struct
type topStats struct {
	Name string
	Time float32
}

type TopStatsList []topStats

type heartbeatStruct struct {
	StartTime   string `json:"start_time"`
	EditorName  string `json:"editor_name"`
	ProjectName string `json:"project_name"`
	Duration    int
	Language    string
	Hostname    string
}
type apiresponse []heartbeatStruct

type Statistics struct {
	Today     float32
	LastWeek  float32
	LastMonth float32
	AllTime   float32

	TopLanguages TopStatsList
	TopProjects  TopStatsList
}

func (s *TopStatsList) SortByTime() TopStatsList {
	sortedArr := *s
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i].Time > sortedArr[j].Time
	})

	return sortedArr
}

func (a *Api) GetStatistics(username string) Statistics {
	res := a.getRequest(fmt.Sprintf("users/%s/activity/data", utils.StringOr(username, "@me")))
	verifyRequest(res.StatusCode, 200)
	defer res.Body.Close()

	var responseJson apiresponse
	utils.Check(json.NewDecoder(res.Body).Decode(&responseJson))

	return calculateCodingStatistics(responseJson)
}

func calculateCodingStatistics(rawdata apiresponse) (codestats Statistics) {
	timenow := time.Now()
	for _, heartbeat := range rawdata {
		if heartbeat.Duration == 0 {
			continue
		}

		getLatestLanguages(heartbeat, &codestats)
		getLatestProjects(heartbeat, &codestats)

		parsedTime, err := time.Parse(ctLayout, heartbeat.StartTime)
		utils.Check(err)

		elapsed := timenow.Sub(parsedTime)
		elapsedHours := elapsed.Hours()
		switch {
		case elapsedHours <= 24:
			codestats.Today += float32(heartbeat.Duration) / 60.0
			fallthrough
		case elapsedHours <= 24*7:
			codestats.LastWeek += float32(heartbeat.Duration) / 60.0
			fallthrough
		case elapsedHours <= 24*30:
			codestats.LastMonth += float32(heartbeat.Duration) / 60.0
		}
		codestats.AllTime += float32(heartbeat.Duration) / 60.0
	}

	return codestats
}

func getLatestLanguages(heartbeat heartbeatStruct, codestats *Statistics) {
	getLatest(heartbeat, &heartbeat.Language, codestats, &codestats.TopLanguages)
}

func getLatestProjects(heartbeat heartbeatStruct, codestats *Statistics) {
	getLatest(heartbeat, &heartbeat.ProjectName, codestats, &codestats.TopProjects)
}

func getLatest(
	heartbeat heartbeatStruct,
	elName *string,
	codestats *Statistics,
	itemsPointer *TopStatsList,
) {
	items := *itemsPointer
	found := false
	for i, itemStats := range items {
		if *elName == itemStats.Name {
			items[i].Time += float32(heartbeat.Duration / 60)
			found = true
			break
		}
	}
	if !found {
		items = append(items, topStats{
			*elName,
			float32(heartbeat.Duration / 60),
		})
	}
	items = items.SortByTime()
	*itemsPointer = items
}
