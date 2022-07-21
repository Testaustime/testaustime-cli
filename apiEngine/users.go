package apiEngine

import (
	"encoding/json"
	"time"

	"github.com/romeq/testaustime-cli/utils"
)

type User struct {
	id         int
	Username   string
	RegTime    time.Time
	FriendCode string
}

func (a *Api) GetProfile() User {
	res := a.makeRequest("users/@me")
	verifyRequest(res.StatusCode, 200)
	defer res.Body.Close()

	responseJson := struct {
		Id                int
		Friend_code       string
		Username          string
		Registration_time string
	}{}
	jsonDecoder := json.NewDecoder(res.Body)
	jsonDecoder.Decode(&responseJson)

    userRegistrationTime, err := time.Parse(api_dateformat, responseJson.Registration_time)
    utils.Check(err)

	return User{
		responseJson.Id,
		responseJson.Username,
		userRegistrationTime,
		responseJson.Friend_code,
	}
}
