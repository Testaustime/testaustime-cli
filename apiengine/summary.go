package apiengine

import (
	"encoding/json"
	"fmt"

	"github.com/romeq/testaustime-cli/utils"
)

type Stats struct {
	Languages map[string]int
	Total     int
}

type SummaryResponse struct {
	AllTime   Stats `json:"all_time"`
	LastMonth Stats `json:"last_month"`
	LastWeek  Stats `json:"last_week"`
}

func (a *Api) Summary(username string) (r SummaryResponse) {
	summaryPath := fmt.Sprintf("users/%s/activity/summary", utils.StringOr(username, "@me"))
	res := a.getRequest(summaryPath)
	defer res.Body.Close()
	verifyResponse(res, 200)

	utils.Check(json.NewDecoder(res.Body).Decode(&r))
	return r
}
