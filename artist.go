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

type Artist struct {
	ID         int
	Name       string
	Picture    string
	Popularity int
	Url        string
}

type ArtistReleases struct {
	Limit              int `json:"limit"`
	Offset             int `json:"offset"`
	TotalNumberOfItems int `json:"totalNumberOfItems"`
	Items              []struct {
		ID                   int         `json:"id"`
		Title                string      `json:"title"`
		Duration             int         `json:"duration"`
		StreamReady          bool        `json:"streamReady"`
		StreamStartDate      string      `json:"streamStartDate"`
		AllowStreaming       bool        `json:"allowStreaming"`
		PremiumStreamingOnly bool        `json:"premiumStreamingOnly"`
		NumberOfTracks       int         `json:"numberOfTracks"`
		NumberOfVideos       int         `json:"numberOfVideos"`
		NumberOfVolumes      int         `json:"numberOfVolumes"`
		ReleaseDate          string      `json:"releaseDate"`
		Copyright            string      `json:"copyright"`
		Type                 string      `json:"type"`
		Version              interface{} `json:"version"`
		URL                  string      `json:"url"`
		Cover                string      `json:"cover"`
		VideoCover           interface{} `json:"videoCover"`
		Explicit             bool        `json:"explicit"`
		Upc                  string      `json:"upc"`
		Popularity           int         `json:"popularity"`
		AudioQuality         string      `json:"audioQuality"`
		Artist               struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"artist"`
		Artists []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"artists"`
	} `json:"items"`
}

func (t Tidal) Artist(artistId int) (Artist, error) {
	path := "/artists/" + strconv.Itoa(artistId)
	data := Artist{}

	// needed params
	get := url.Values{}
	post := url.Values{}

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Getting artist data failed.")
	}

	return data, nil
}

func (t Tidal) ArtistAlbums(artistId int) (ArtistReleases, error) {
	path := "/artists/" + strconv.Itoa(artistId) + "/albums"
	data := ArtistReleases{}

	// needed params
	get := url.Values{}
	get.Set("limit", "999")
	post := url.Values{}

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Getting artist release data failed.")
	}

	return data, nil
}

func (t Tidal) ArtistEpsAndSingles(artistId int) (ArtistReleases, error) {
	path := "/artists/" + strconv.Itoa(artistId) + "/albums"
	data := ArtistReleases{}

	// needed params
	get := url.Values{}
	get.Set("limit", "999")
	get.Add("filter", "EPSANDSINGLES")
	post := url.Values{}

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Getting artist releases data failed.")
	}

	return data, nil
}

func (t Tidal) ArtistCompilations(artistId int) (ArtistReleases, error) {
	path := "/artists/" + strconv.Itoa(artistId) + "/albums"
	data := ArtistReleases{}

	// needed params
	get := url.Values{}
	get.Set("limit", "999")
	get.Add("filter", "COMPILATIONS")
	post := url.Values{}

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Getting artist releases data failed.")
	}

	return data, nil
}
