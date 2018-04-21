// Copyright (c) 2018 ChrisOboe
//
// This file is part of tidapi
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package tidapi

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const baseUrl string = "https://api.tidal.com/v1"
const token string = "4zx46pyr9o8qZNRw"
const clientVersion string = "2.4.5--3"

type Tidal struct {
	http        *http.Client
	sessionId   string
	countryCode string
}

func New() Tidal {
	return Tidal{
		&http.Client{Timeout: time.Second * 10},
		"",
		"WW",
	}
}

func (t *Tidal) SetCountryCode(countryCode string) {
	t.countryCode = countryCode
}

func (t *Tidal) SetSessionId(sessionId string) {
	t.sessionId = sessionId
}

func (t Tidal) CountryCode() string {
	return t.countryCode
}

func (t Tidal) SessionId() string {
	return t.sessionId
}

func (t Tidal) AppToken() string {
	return token
}

func (t Tidal) HttpClient() *http.Client {
	return t.http
}

func (t Tidal) Request(path string, get url.Values, post url.Values) (*http.Response, error) {
	return t.request(path, get, post)
}

// request does an request to the tidal api
// When no post values are specified it does an GET, else it
// does an POST.
// When sessionId is set it appends the sessionID header
// When countryCode is set it appends the countryCode parameter
// it returns an response
func (t Tidal) request(path string, get url.Values, post url.Values) (*http.Response, error) {
	// check if post or get request
	method := "POST"
	if len(post) == 0 {
		method = "GET"
	}

	// append needed parameters
	get.Add("countryCode", t.countryCode)
	if method == "POST" {
		post.Add("clientVersion", clientVersion)
	}

	// build body
	body := ""
	if method == "POST" {
		body = post.Encode()
	}

	// build request
	request, err := http.NewRequest(method,
		baseUrl+path+"?"+get.Encode(),
		bytes.NewBufferString(body))
	if err != nil {
		return nil, errors.Wrap(err, "Creating request failed.")
	}

	// create headers
	// Standard stuff
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) TIDAL/2.1.14 Chrome/58.0.3029.110 Electron/1.7.10 Safari/537.36")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	// Tidal specific stuff
	request.Header.Add("X-Tidal-Token", token)
	if t.sessionId != "" {
		request.Header.Add("X-Tidal-SessionId", t.sessionId)
	}

	// fire request
	response, err := t.http.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "Getting response failed.")
	}

	// we should get an OK here
	if response.StatusCode != http.StatusOK {
		dump, _ := httputil.DumpResponse(response, true)
		return nil, errors.New("Tidal request didn't receive an OK status. Request: " +
			request.URL.String() + " Response: " + string(dump))
	}

	return response, nil
}

// jsonRequest is like request, but instead of returning the response
// it writes the json data to the struct specified by data
func (t Tidal) jsonRequest(path string, get url.Values, post url.Values, data interface{}) error {
	response, err := t.request(path, get, post)
	if err != nil {
		return errors.Wrap(err, "Request from jsonRequest failed.")
	}

	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		return errors.Wrap(err, "jsonRequest couldn't decode the response body.")
	}
	return nil
}
