package apiengine

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/romeq/testaustime-cli/utils"
)

type topStats struct {
	Name string
	Time float32
}

type heartbeatStruct struct {
	StartTime   string `json:"start_time"`
	EditorName  string `json:"editor_name"`
	ProjectName string `json:"project_name"`
	Duration    int
	Language    string
	Hostname    string
}

type Statistics struct {
	Today     float32
	PastWeek  float32
	PastMonth float32
	AllTime   float32

	TopLanguages TopStatsList
	TopProjects  TopStatsList
}

type TopStatsList []topStats
type apiresponse []heartbeatStruct

func (a *Api) GetStatistics(username string, latest bool, since time.Time) Statistics {
	res := a.getRequest(fmt.Sprintf("users/%s/activity/data", utils.StringOr(username, "@me")))
	verifyRequest(res.StatusCode, 200)
	defer res.Body.Close()

	var responseJson apiresponse
	utils.Check(json.NewDecoder(res.Body).Decode(&responseJson))

	return calculateCodingStatistics(responseJson, latest, since)
}

func calculateCodingStatistics(
	rawdata apiresponse,
	latest bool,
	since time.Time,
) (codestats Statistics) {
	timenow := time.Now()
	for _, heartbeat := range rawdata {
		if heartbeat.Duration == 0 {
			continue
		}

		parsedTime, err := time.Parse(ctLayout, heartbeat.StartTime)
		utils.Check(err)
		if since.Sub(parsedTime) > 0 {
			continue
		}

		if latest {
			getTopLanguages(heartbeat, &codestats, since)
			getTopProjects(heartbeat, &codestats, since)
		}

		elapsed := timenow.Sub(parsedTime)
		elapsedHours := elapsed.Hours()
		switch {
		case elapsedHours <= 24:
			codestats.Today += float32(heartbeat.Duration) / 60.0
			fallthrough
		case elapsedHours <= 24*7:
			codestats.PastWeek += float32(heartbeat.Duration) / 60.0
			fallthrough
		case elapsedHours <= 24*30:
			codestats.PastMonth += float32(heartbeat.Duration) / 60.0
		}
		codestats.AllTime += float32(heartbeat.Duration) / 60.0
	}

	return codestats
}

func (s *TopStatsList) SortByTime() TopStatsList {
	sortedArr := *s
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i].Time > sortedArr[j].Time
	})

	return sortedArr
}

func getTopLanguages(heartbeat heartbeatStruct, codestats *Statistics, since time.Time) {
	getTop(heartbeat, &heartbeat.Language, codestats, &codestats.TopLanguages, since)
}

func getTopProjects(heartbeat heartbeatStruct, codestats *Statistics, since time.Time) {
	getTop(heartbeat, &heartbeat.ProjectName, codestats, &codestats.TopProjects, since)
}

func getTop(
	heartbeat heartbeatStruct,
	elName *string,
	codestats *Statistics,
	itemsPointer *TopStatsList,
	since time.Time,
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
