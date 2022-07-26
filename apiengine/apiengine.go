package apiengine

import (
	"bytes"
	"errors"
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

type Api struct {
	token string
	url   string
}

var ctLayout string = "2006-01-02T15:04:05"

// New creates a new Api struct with given parameters.
// if url is empty string, it will be replaced with
// current production api server.
func New(token, url string) Api {
	return Api{
		token,
		utils.StringOr(url, "https://api.testaustime.fi"),
	}
}

// verifyRequest will make sure statusCode matches wantedStatusCode.
// If they don't match, it will result in an error which is handled by
// logger.Error().
func verifyRequest(statusCode, wantedStatusCode int) {
	switch statusCode {
	case http.StatusUnauthorized:
		logger.Error(errors.New("Request failed. You are not authorized."))

	case wantedStatusCode:
		return

	default:
		logger.Error(fmt.Errorf("Request failed with status code %d.", statusCode))
	}
}

// getRequest sends a GET request to path with authentication
// and returns it's response
func (a *Api) getRequest(path string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", a.url, path), nil)
	utils.Check(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))

	res, err := client.Do(req)
	utils.Check(err)

	return res
}

// postRequest sends a POST request with wanted data and
// authentication
func (a *Api) postRequest(path string, requestData []byte) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", a.url, path),
		bytes.NewBuffer(requestData),
	)

	utils.Check(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))
	req.Header.Set("Content-type", "application/json")

	res, err := client.Do(req)
	utils.Check(err)

	return res
}

// deleteRequest sends a DELETE request with wanted data and
// authentication
func (a *Api) deleteRequest(path string, requestData []byte) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/%s", a.url, path),
		bytes.NewBuffer(requestData),
	)

	utils.Check(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))
	req.Header.Set("Content-type", "application/json")

	res, err := client.Do(req)
	utils.Check(err)

	return res
}

// Implement UnmarshalJSON for DateFormat
// so that date can be parsed from API responses
func (ct *DateFormat) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}
