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

type AlbumItems struct {
	Limit              int
	Offset             int
	TotalNumberOfItems int
	Items              []struct {
		Item struct {
			ID                   int
			Title                string
			Duration             int
			ReplayGain           float64
			Peak                 float64
			AllowStreaming       bool
			StreamReady          bool
			StreamStartDate      string
			PremiumStreamingOnly bool
			TrackNumber          int
			VolumeNumber         int
			Version              interface{}
			Popularity           int
			Copyright            string
			URL                  string
			Isrc                 string
			Editable             bool
			Explicit             bool
			AudioQuality         string
			Artist               struct {
				ID   int
				Name string
				Type string
			}
			Artists []struct {
				ID   int
				Name string
				Type string
			}
			Album struct {
				ID    int
				Title string
				Cover string
			}
		}
		Type string
	}
}

type Album struct {
	ID                   int
	Title                string
	Duration             int
	StreamReady          bool
	StreamStartDate      string
	AllowStreaming       bool
	PremiumStreamingOnly bool
	NumberOfTracks       int
	NumberOfVideos       int
	NumberOfVolumes      int
	ReleaseDate          string
	Copyright            string
	Type                 string
	Version              interface{}
	URL                  string
	Cover                string
	VideoCover           interface{}
	Explicit             bool
	Upc                  string
	Popularity           int
	AudioQuality         string
	Artist               struct {
		ID   int
		Name string
		Type string
	}
	Artists []struct {
		ID   int
		Name string
		Type string
	}
}

func (t Tidal) AlbumItems(albumId int) (AlbumItems, error) {
	// declare local constants
	path := "/albums/" + strconv.Itoa(albumId) + "/items"

	// create output struct
	data := AlbumItems{}

	// needed params
	get := url.Values{}
	get.Add("offset", "0")
	get.Add("limit", "100")
	post := url.Values{}

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Getting album items failed.")
	}

	return data, nil
}

func (t Tidal) Album(albumId int) (Album, error) {
	// declare local constants
	path := "/albums/" + strconv.Itoa(albumId)

	// create output struct
	data := Album{}

	// needed params
	get := url.Values{}
	post := url.Values{}

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Getting album failed.")
	}

	return data, nil
}
