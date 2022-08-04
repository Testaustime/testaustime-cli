package apiengine

import (
	"encoding/json"

	"github.com/romeq/testaustime-cli/utils"
)

type UserResponse struct {
	Id         int
	Username   string `json:"username"`
	Token      string `json:"auth_token"`
	FriendCode string `json:"friend_code"`
	RegTime    string `json:"registration_time"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

type Settings struct {
	PublicProfile any
}

// Login returns UserResponse which includes data what
// was given in API response. If login fails, errResponse will
// contain the reason for the failure.
func (a *Api) Login(
	username,
	password string,
) (response UserResponse, errResponse ErrorResponse) {
	requestJson, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: username,
		Password: password,
	})
	utils.Check(err)

	res := a.postRequest("auth/login", requestJson)
	defer res.Body.Close()
	verifyResponse(res, 200)

	utils.Check(json.NewDecoder(res.Body).Decode(&response))
	return response, errResponse
}

// Register returns UserResponse which includes data what
// was given in API response. If registration fails, errResponse will
// contain the reason for the failure.
func (a *Api) Register(
	username,
	password string,
) (response UserResponse, errResponse ErrorResponse) {
	requestJson, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: username,
		Password: password,
	})
	utils.Check(err)

	res := a.postRequest("auth/register", requestJson)
	defer res.Body.Close()
	verifyResponse(res, 200)

	utils.Check(json.NewDecoder(res.Body).Decode(&response))

	return response, errResponse
}

// ChangePassword changes account's password using the old password and a new password.
// Requests are made using authentication token.
func (a *Api) ChangePassword(
	oldPassword,
	newPassword string,
) {
	requestJson, err := json.Marshal(struct {
		Old string `json:"old"`
		New string `json:"new"`
	}{
		Old: oldPassword,
		New: newPassword,
	})
	utils.Check(err)

	res := a.postRequest("auth/changepassword", requestJson)
	defer res.Body.Close()
	verifyResponse(res, 200)
}

// GetAuthtoken returns current authentication token
func (a *Api) GetAuthtoken() string {
	return a.token
}

// NewAuthtoken generates a new authentication token and
// returns it. Request is made using old authentication token.
func (a *Api) NewAuthtoken() string {
	res := a.postRequest("auth/regenerate", nil)
	defer res.Body.Close()
	verifyResponse(res, 200)

	response := struct {
		Token string
	}{}
	utils.Check(json.NewDecoder(res.Body).Decode(&response))

	return response.Token
}

// NewFriendcode makes a request to testaustime API generating
// a new friend code and returns it. Request is made using authentication
// token.
func (a *Api) NewFriendcode() string {
	res := a.postRequest("friends/regenerate", nil)
	defer res.Body.Close()
	verifyResponse(res, 200)

	response := struct {
		FriendCode string `json:"friend_code"`
	}{}
	utils.Check(json.NewDecoder(res.Body).Decode(&response))

	return response.FriendCode
}

func (a *Api) UpdateSettings(settings Settings) {
	reqJson, err := json.Marshal(settings)
	utils.Check(err)

	res := a.postRequest("account/settings", reqJson)
	verifyResponse(res, 200)
	defer res.Body.Close()
}
