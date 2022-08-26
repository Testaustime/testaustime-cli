package apiengine

import (
	"encoding/json"
	"sort"

	"github.com/romeq/testaustime-cli/utils"
)

type FriendActivity struct {
	AllTime   float32 `json:"all_time"`
	PastMonth float32 `json:"past_month"`
	PastWeek  float32 `json:"past_week"`
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

// Get list of friends
func (a *Api) Friends() (friends Friends) {
	res := a.getRequest("friends/list")
	verifyResponse(res, 200)
	defer res.Body.Close()

	utils.Check(json.NewDecoder(res.Body).Decode(&friends))
	return friends
}

// AddFriend adds a new friend
func (a *Api) AddFriend(friendCode string) (friend Friend) {
	res := a.postRequest("friends/add", []byte(friendCode))
    verifyResponse(res, 200)
	defer res.Body.Close()

	utils.Check(json.NewDecoder(res.Body).Decode(&friend))
	return friend
}

// RemoveFriend removes a friend
func (a *Api) RemoveFriend(friendName string) (errResponse ErrorResponse) {
	res := a.deleteRequest("friends/remove", []byte(friendName))
	verifyResponse(res, 200)
	defer res.Body.Close()

	return errResponse
}

// AddSelf returns list of friends user's account has been appended
func (f Friends) AddSelf(statistics Statistics) *Friends {
	f = append(f, Friend{
		"@me",
		FriendActivity{
			statistics.AllTime * 60,
			statistics.PastMonth * 60,
			statistics.PastWeek * 60,
		},
	})
	return &f
}

// AllTime sorts friends' data with their past month's coding statistics
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

// PastMonth sorts friends' data with their past month's coding statistics
func (f *Friends) PastMonth() (result []FriendsCodingTime) {
	friends := *f
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].CodingTime.PastMonth > friends[j].CodingTime.PastMonth
	})
	for _, x := range *f {
		result = append(result, FriendsCodingTime{
			x.Username,
			x.CodingTime.PastMonth,
		})
	}

	return result
}

// PastWeek sorts friends' data with their past week's coding statistics
func (f *Friends) PastWeek() (result []FriendsCodingTime) {
	friends := *f
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].CodingTime.PastWeek > friends[j].CodingTime.PastWeek
	})
	for _, x := range *f {
		result = append(result, FriendsCodingTime{
			x.Username,
			x.CodingTime.PastWeek,
		})
	}

	return result
}
