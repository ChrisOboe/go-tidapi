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

type UrlUsageMode string

const (
	UrlUsageModeStream  UrlUsageMode = "STREAM"
	UrlUsageModeOffline UrlUsageMode = "OFFLINE"
)

type Audioquality string

const (
	AudioqualityLow      Audioquality = "LOW"
	AudioqualityHigh     Audioquality = "HIGH"
	AudioqualityLossless Audioquality = "LOSSLESS"
	AudioqualityHiRes    Audioquality = "HI_RES"
)

type TracksUrlpostpaywall struct {
	AssetPresentation string
	AudioQuality      string
	Codec             string
	SecurityToken     string
	SecurityType      string
	TrackID           int
	Urls              []string
}

type TracksUrl struct {
	URL                   string
	TrackID               int
	PlayTimeLeftInMinutes int
	SoundQuality          string
	EncryptionKey         string
	Codec                 string
}

func (t Tidal) Tracks(trackId int) (TracksUrlpostpaywall, error) {

	// declare local constants
	path := "/tracks/" + strconv.Itoa(trackId)

	// create output struct
	data := TracksUrlpostpaywall{}

	// needed params
	get := url.Values{}
	post := url.Values{}

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Getting track failed.")
	}

	return data, nil
}

func (t Tidal) TracksUrlpostpaywall(trackId int,
	audioquality Audioquality, urlusagemode UrlUsageMode) (TracksUrlpostpaywall, error) {

	// declare local constants
	path := "/tracks/" + strconv.Itoa(trackId) + "/urlpostpaywall"

	// create output struct
	data := TracksUrlpostpaywall{}

	// needed params
	get := url.Values{}
	get.Add("audioquality", string(audioquality))
	get.Add("urlusagemode", string(urlusagemode))
	get.Add("assetpresentation", "FULL")
	post := url.Values{}

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Getting track failed.")
	}

	return data, nil
}

// the following API somehow works. But the desktop client doesn't use it.
// There are a lot of API functions the desktop client doesn't make use off.
/**
func (t Tidal) tracksUrl(trackId int, audioquality Audioquality,
	urlType int) (TracksUrl, error) {

	path := "/tracks/" + strconv.Itoa(trackId)
	if urlType == typeStream {
		path += "/streamUrl"
	} else {
		path += "/offlineUrl"
	}

	data := TracksUrl{}

	// needed params
	get := url.Values{}
	get.Add("soundQuality", string(audioquality))
	post := url.Values{}

	// do request
	err := t.jsonRequest(path, get, post, &data)
	if err != nil {
		return data, errors.Wrap(err, "Getting track failed.")
	}

	return data, nil
}

func (t Tidal) TracksOfflineUrl(trackId int,
	audioquality Audioquality) (TracksUrl, error) {

	return t.tracksUrl(trackId, audioquality, typeOffline)
}

func (t Tidal) TracksStreamUrl(trackId int,
	audioquality Audioquality) (TracksUrl, error) {

	return t.tracksUrl(trackId, audioquality, typeStream)
}
**/
