package apiengine

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/romeq/testaustime-cli/utils"
)

type LeaderboardUser struct {
	Username  string
	Admin     bool
	TimeCoded int `json:"time_coded"`
}

type Leaderboard struct {
	Name         string
	Invite       string
	CreationTime DateFormat
	Members      []LeaderboardUser
}

type LeaderboardList []struct {
	Name        string
	MemberCount int `json:"member_count"`
}

type InviteCode struct {
	Code string `json:"invite_code"`
}

// Leaderboards function returns a list of your leaderboards
// including their names and member counts
func (a *Api) Leaderboards() (r LeaderboardList) {
	res := a.getRequest("users/@me/leaderboards")
	verifyResponse(res, 200)
	defer res.Body.Close()

	utils.Check(json.NewDecoder(res.Body).Decode(&r))
	return r
}

// NewLeaderboard creates a new leaderboard with given name
func (a *Api) NewLeaderboard(name string) (r InviteCode) {
	reqJson, err := json.Marshal(map[string]string{
		"name": name,
	})
	utils.Check(err)

	res := a.postRequest("leaderboards/create", reqJson)
	verifyResponse(res, 200)
	defer res.Body.Close()

	utils.Check(json.NewDecoder(res.Body).Decode(&r))
	return r
}

// Leaderboard returns a struct including all possible information
// from a leaderboard
func (a *Api) Leaderboard(name string) (r Leaderboard) {
	res := a.getRequest(fmt.Sprintf("leaderboards/%s", name))
	verifyResponse(res, 200)
	defer res.Body.Close()

	utils.Check(json.NewDecoder(res.Body).Decode(&r))
	return r
}

// JoinLeaderboard joins a new leaderboard with given code
func (a *Api) JoinLeaderboard(code string) {
	reqJson, err := json.Marshal(map[string]string{
		"invite": code,
	})
	utils.Check(err)

	res := a.postRequest("leaderboards/join", reqJson)
	verifyResponse(res, 200)
	defer res.Body.Close()
}

// KickMember kicks given user from given leaderboard
func (a *Api) KickMember(leaderboardName, username string) {
	reqJson, err := json.Marshal(map[string]string{
		"user": username,
	})
	utils.Check(err)

	res := a.postRequest(fmt.Sprintf("leaderboards/%s/kick", leaderboardName), reqJson)
	verifyResponse(res, 200)
	defer res.Body.Close()
}

// LeaveLeaderboard leaves from given leaderboard
func (a *Api) LeaveLeaderboard(name string) {
	res := a.postRequest(fmt.Sprintf("leaderboards/%s/leave", name), []byte{})
	verifyResponse(res, 200)
	defer res.Body.Close()
}

// DeleteLeaderboard deletes a leaderboard. Admin permissions are required
func (a *Api) DeleteLeaderboard(name string) {
	res := a.deleteRequest(fmt.Sprintf("leaderboards/%s", name), []byte{})
	verifyResponse(res, 200)
	defer res.Body.Close()
}

// RegenerateLeaderboardInvite regenerates a leaderboard invite code. Admin
// permissions are required
func (a *Api) RegenerateLeaderboardInvite(name string) (r InviteCode) {
	res := a.postRequest(fmt.Sprintf("leaderboards/%s/regenerate", name), []byte{})
	verifyResponse(res, 200)
	defer res.Body.Close()

	utils.Check(json.NewDecoder(res.Body).Decode(&r))
	return r
}

// SortMembersByTime returns users sorted by tiem coded in it's leaderboard
func (l *Leaderboard) SortMembersByTime() []LeaderboardUser {
	sortedArr := l.Members
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i].TimeCoded > sortedArr[j].TimeCoded
	})

	return sortedArr
}
