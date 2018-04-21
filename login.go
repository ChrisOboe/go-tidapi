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
	"github.com/pkg/errors"
	"net/url"
)

type LoginUsername struct {
	UserId      int
	SessionId   string
	CountryCode string
}

func (t Tidal) LoginUsername(username string, password string) (LoginUsername, error) {
	// declare local constants
	const path string = "/login/username"

	// create output struct
	data := LoginUsername{}

	// needed params
	get := url.Values{}
	post := url.Values{}
	post.Add("username", username)
	post.Add("password", password)

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Couldn't log in")
	}

	return data, nil
}
