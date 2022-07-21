package apiEngine

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/romeq/testaustime-cli/logger"
	"github.com/romeq/testaustime-cli/utils"
)

var api_dateformat string = "2006-01-02T15:04:05"

type Api struct {
	token string
	url   string
}

func New(token string, url string) Api {
	return Api{
		token,
		utils.StringOr(url, "https://api.testaustime.fi"),
	}
}

func verifyRequest(statusCode int, wantedStatusCode int) {
	if statusCode != wantedStatusCode {
		logger.Error(errors.New(fmt.Sprintf(
			"Request failed with status code %d.", statusCode)))
	}
}

func (a *Api) makeRequest(path string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", a.url, path), nil)
	utils.Check(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))

	res, err := client.Do(req)
	utils.Check(err)
	return res
}
