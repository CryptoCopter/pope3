/*
Copyright 2016 Markus Sommer

This file is part of Pope3.

Pope3 is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Pope3 is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Pope3.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/json"
	"io/ioutil"
)

type User struct {
	Username string
	Password string
}

func LoadUserArray(filename string) ([]User, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func UserArrayToMap(userArray []User) map[string]string {
	userMap := make(map[string]string)

	for _, element := range userArray {
		userMap[element.Username] = element.Password
	}

	return userMap
}

func LoadUsers(filename string) (map[string]string, error) {
	userArray, err := LoadUserArray(filename)
	if err != nil {
		return nil, err
	}
	return UserArrayToMap(userArray), nil
}
