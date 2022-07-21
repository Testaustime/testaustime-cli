package apiEngine

import (
	"encoding/json"
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

func (s *TopStatsList) SortByTime() {
	sortedArr := *s
	compareTime := func(i, j int) bool {
		return sortedArr[i].Time > sortedArr[j].Time
	}

	sort.Slice(*s, compareTime)
}

type heartbeatStruct struct {
	Start_time   string
	Duration     int
	Project_name string
	Language     string
	Editor_name  string
	Hostname     string
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

func (a *Api) GetStatistics() Statistics {
	res := a.makeRequest("users/@me/activity/data")
	verifyRequest(res.StatusCode, 200)
	defer res.Body.Close()

	var responseJson apiresponse
	jsonDecoder := json.NewDecoder(res.Body)
	jsonDecoder.Decode(&responseJson)
	return calculateCodingStatistics(responseJson)
}

func calculateCodingStatistics(rawdata apiresponse) (codestats Statistics) {
	timenow := time.Now()
	for _, heartbeat := range rawdata {
		if heartbeat.Duration == 0 {
			continue
		}

		getTopLanguages(heartbeat, &codestats)
		getTopProjects(heartbeat, &codestats)

		parsedTime, err := time.Parse(api_dateformat, heartbeat.Start_time)
		utils.Check(err)

		elapsed := timenow.Sub(parsedTime)
		elapsedHours := elapsed.Hours()
		switch {
		case elapsedHours < 24:
			codestats.Today += float32(heartbeat.Duration) / 60.0
		case elapsedHours < 24*7:
			codestats.LastWeek += float32(heartbeat.Duration) / 60.0
		case elapsedHours < 24*30:
			codestats.LastMonth += float32(heartbeat.Duration) / 60.0
		}
		codestats.AllTime += float32(heartbeat.Duration) / 60.0
	}

	return codestats
}

func getTopLanguages(heartbeat heartbeatStruct, codestats *Statistics) {
	getTop(heartbeat, &heartbeat.Language, codestats, &codestats.TopLanguages)
}

func getTopProjects(heartbeat heartbeatStruct, codestats *Statistics) {
	getTop(heartbeat, &heartbeat.Project_name, codestats, &codestats.TopProjects)
}

func getTop(
	heartbeat heartbeatStruct,
	y *string,
	codestats *Statistics,
	x *TopStatsList,
) {
	z := *x
	found := false
	for i, itemStats := range z {
		if *y == itemStats.Name {
			z[i].Time += float32(heartbeat.Duration / 60)
			found = true
			break
		}
	}
	if !found {
		z = append(z, topStats{
			*y,
			float32(heartbeat.Duration / 60),
		})
	}
	z.SortByTime()
	*x = z
}
