package apiEngine

import (
	"encoding/json"

	"github.com/romeq/testaustime-cli/utils"
)

type User struct {
	Id         int
	RegTime    DateFormat `json:"registration_time"`
	Username   string     `json:"username"`
	FriendCode string     `json:"friend_code"`
}

func (a *Api) GetProfile() (user User) {
	res := a.getRequest("users/@me")
	verifyRequest(res.StatusCode, 200)
	defer res.Body.Close()

	utils.Check(json.NewDecoder(res.Body).Decode(&user))
	return user
}

func (a *Api) GetAuthtoken() string {
	return a.token
}

func (a *Api) NewAuthtoken() string {
	res := a.postRequest("auth/regenerate")
	verifyRequest(res.StatusCode, 200)
	defer res.Body.Close()

	response := struct {
		Token string
	}{}
	utils.Check(json.NewDecoder(res.Body).Decode(&response))

	return response.Token
}

func (a *Api) NewFriendcode() string {
	res := a.postRequest("friends/regenerate")
	verifyRequest(res.StatusCode, 200)
	defer res.Body.Close()

	response := struct {
		Friend_code string
	}{}
	utils.Check(json.NewDecoder(res.Body).Decode(&response))

	return response.Friend_code
}
