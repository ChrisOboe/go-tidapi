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

package main

import (
	"fmt"
	"github.com/ChrisOboe/tidapi"
	"net/url"
	"os"
)

func main() {
	// parse args
	if len(os.Args) < 3 {
		fmt.Println("Usage: request <username> <password> <path>")
		os.Exit(0)
	}

	user := os.Args[1]
	pass := os.Args[2]
	path := os.Args[3]

	api := tidapi.New()
	login, err := api.LoginUsername(user, pass)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	api.SetSessionId(login.SessionId)
	api.SetCountryCode(login.CountryCode)

	get := url.Values{}
	post := url.Values{}

	result, err := api.Request(path, get, post)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
