package apiengine

import (
	"encoding/json"

	"github.com/romeq/testaustime-cli/utils"
)

type LeaderboardList []struct{
    Name string
    MemberCount int `json:"member_count"`
}

type InviteCode struct {
    Code string `json:"invite_code"`
}


func (a *Api) Leaderboards() (r LeaderboardList) {
    res := a.getRequest("users/@me/leaderboards")
    verifyResponse(res, 200)
    defer res.Body.Close()
    
    utils.Check(json.NewDecoder(res.Body).Decode(&r))
    return r
}

func (a *Api) JoinLeaderboard(code string) {
    reqJson, err := json.Marshal(map[string]string {
        "invite": code,
    })
    utils.Check(err)

    res := a.postRequest("leaderboards/join", reqJson)
    verifyResponse(res, 200)
    defer res.Body.Close()
}

func (a *Api) NewLeaderboard(name string) (r InviteCode) {
    reqJson, err := json.Marshal(map[string]string {
        "name": name,
    })
    utils.Check(err)

    res := a.postRequest("leaderboards/create", reqJson)
    verifyResponse(res, 200)
    defer res.Body.Close()

    utils.Check(json.NewDecoder(res.Body).Decode(&r))
    return r
}

