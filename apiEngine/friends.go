package apiEngine

import (
	"encoding/json"
	"sort"

	"github.com/romeq/testaustime-cli/utils"
)

type FriendActivity struct {
	AllTime   float32 `json:"all_time"`
	LastMonth float32 `json:"past_month"`
	LastWeek  float32 `json:"past_week"`
}

type Friend struct {
	Username   string         `json:"username"`
	CodingTime FriendActivity `json:"coding_time"`
}

type FriendsCodingTime struct {
	Username   string
	CodingTime float32
}

type Friends []Friend

func (a *Api) GetFriends() (friends Friends) {
	res := a.getRequest("friends/list")
	verifyRequest(res.StatusCode, 200)
	defer res.Body.Close()

	utils.Check(json.NewDecoder(res.Body).Decode(&friends))
	return friends
}

func (f Friends) AddSelf(statistics Statistics) *Friends {
	f = append(f, Friend{
		"@me",
		FriendActivity{
			statistics.AllTime * 60,
			statistics.LastMonth * 60,
			statistics.LastWeek * 60,
		},
	})
	return &f
}

// LastMonth sorts friends' data with their past month's coding statistics
func (f *Friends) AllTime() (result []FriendsCodingTime) {
	friends := *f
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].CodingTime.AllTime > friends[j].CodingTime.AllTime
	})
	for _, x := range *f {
		result = append(result, FriendsCodingTime{
			x.Username,
			x.CodingTime.AllTime,
		})
	}

	return result
}

// LastMonth sorts friends' data with their past month's coding statistics
func (f *Friends) LastMonth() (result []FriendsCodingTime) {
	friends := *f
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].CodingTime.LastMonth > friends[j].CodingTime.LastMonth
	})
	for _, x := range *f {
		result = append(result, FriendsCodingTime{
			x.Username,
			x.CodingTime.LastMonth,
		})
	}

	return result
}

func (f *Friends) LastWeek() (result []FriendsCodingTime) {
	friends := *f
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].CodingTime.LastWeek > friends[j].CodingTime.LastWeek
	})
	for _, x := range *f {
		result = append(result, FriendsCodingTime{
			x.Username,
			x.CodingTime.LastWeek,
		})
	}

	return result
}
