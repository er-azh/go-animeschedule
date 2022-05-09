package animeschedule

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const APIRoot = "https://animeschedule.net/api/v2"

// getJSON is a simple helper function that does a HTTP request and returns the parsed response as T
// It will throw an error if the http response code is not 200
func getJSON[T any](url string) (*T, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// make a client and do the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseBytes, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("invalid http response: %d,response: %s", response.StatusCode, string(responseBytes))
	}

	var resp T
	return &resp, json.NewDecoder(response.Body).Decode(&resp)
}

type TimetableTime struct {
	Year, Week int
}

// GetCurrentTimetable gets the current timetable or the timetable in a specific week
// when timetableTime is specified it tries to get that week's timetable.
// timezone can be set to return times in a different time zone.
func GetCurrentTimetable(timetableTime *TimetableTime, timezone *time.Location) (*Timetable, error) {
	values := url.Values{}
	if timetableTime != nil {
		values.Add("year", strconv.Itoa(timetableTime.Year))
		values.Add("week", strconv.Itoa(timetableTime.Week))
	}
	if timezone != nil {
		values.Add("tz", timezone.String())
	}
	return getJSON[Timetable](APIRoot + "/timetables?" + values.Encode())
}

// GetShow gets a show by its slug
func GetShow(slug string) (*ShowDetail, error) {
	return getJSON[ShowDetail](APIRoot + "/anime/" + url.QueryEscape(slug))
}
