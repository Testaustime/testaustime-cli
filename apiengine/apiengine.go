package apiengine

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	token                 string
	url                   string
	caseInsensitiveFields []string
}

var MeasureTime bool = false
var ctLayout string = "2006-01-02T15:04:05"

// New creates a new Api struct with given parameters.
// if url is empty string, it will be replaced with
// current production api server.
func New(token, url string, caseInsensitiveFields []string) Api {
	return Api{
		token,
		utils.StringOr(url, "https://api.testaustime.fi"),
		caseInsensitiveFields,
	}
}

// verifyResponse will make sure res.StatusCode matches wantedStatusCode.
// If they don't match, function will result in error defined in response
// which is then handled by logger.Error().
func verifyResponse(res *http.Response, wantedStatusCode int) {
	if res.StatusCode == wantedStatusCode {
		return
	}

	var errResponse ErrorResponse
	err := json.NewDecoder(res.Body).Decode(&errResponse)
	if err == io.EOF {
		logger.Error(fmt.Errorf("Request failed (%d)", res.StatusCode))
	}
	utils.Check(err)

	switch res.StatusCode {
	case http.StatusUnauthorized:
		logger.Error(errors.New(
			"You are not authorized",
		))

	default:
		logger.Error(fmt.Errorf("Request failed: \"%s\" (%d)", errResponse.Err,
			res.StatusCode))
	}
}

// sendRequest sends request with given client and measures time if
// wanted
func sendRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	if !MeasureTime {
		return client.Do(req)
	}

	timeBefore := time.Now()
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf(
		"Time taken in request: %.5fs", time.Since(timeBefore).Seconds()))
	logger.Info(fmt.Sprintf(
		"Response size: %s", responseSize(float32(res.ContentLength))))
	logger.Info("")

	return res, err
}

// getRequest sends a GET request to path with authentication
// and returns it's response
func (a *Api) getRequest(path string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", a.url, path), nil)
	utils.Check(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))

	res, err := sendRequest(client, req)
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

	res, err := sendRequest(client, req)
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

	res, err := sendRequest(client, req)
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

// responseSize returns string formatted with response size in
// either kilobytes or megabytes, depending on it's size
func responseSize(responseLength float32) string {
	responseSizeInKBs := responseLength / 1024
	responseSizeInMBs := responseSizeInKBs / 1024

	switch {
	case responseSizeInKBs > 0:
		return fmt.Sprintf("%.0f K", responseSizeInKBs)
	case responseSizeInMBs > 0:
		return fmt.Sprintf("%.3f M", responseSizeInMBs)
	default:
		return fmt.Sprint(responseLength)
	}
}
