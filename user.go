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
	"strconv"
)

type UsersFavorites struct {
	Playlist []string
	Track    []string
	Artist   []string
	Video    []string
	Album    []string
}

func (t Tidal) UserFavorites(userId int) (UsersFavorites, error) {
	path := "/users/" + strconv.Itoa(userId) + "/favorites/ids"
	data := UsersFavorites{}

	// needed params
	get := url.Values{}
	post := url.Values{}

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Getting favorites failed.")
	}

	return data, nil
}
