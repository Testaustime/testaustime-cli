package apiEngine

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/romeq/testaustime-cli/logger"
	"github.com/romeq/testaustime-cli/utils"
)

type DateFormat struct {
	time.Time
}

var ctLayout string = "2006-01-02T15:04:05"
var nilTime = (time.Time{}).UnixNano()

func (ct *DateFormat) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *DateFormat) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}


type Api struct {
	token string
	url   string
}

func New(token, url string) Api {
	return Api{
		token,
		utils.StringOr(url, "https://api.testaustime.fi"),
	}
}

func verifyRequest(statusCode, wantedStatusCode int) {
	if statusCode != wantedStatusCode {
		logger.Error(fmt.Errorf("Request failed with status code %d.", statusCode))
	}
}

func (a *Api) getRequest(path string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", a.url, path), nil)
	utils.Check(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))

	res, err := client.Do(req)
	utils.Check(err)

	return res
}

func (a *Api) postRequest(path string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", a.url, path), nil)
	utils.Check(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))

	res, err := client.Do(req)
	utils.Check(err)

	return res
}
