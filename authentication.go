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

import "errors"

type AuthState struct {
	Username      string
	Authenticated bool
	Users         map[string]string
}

func (authstate *AuthState) SetUser(username string) error {
	_, prs := authstate.Users[username]
	if prs {
		authstate.Username = username
		return nil
	} else {
		return errors.New("User does not exist")
	}
}

func (authstate *AuthState) Authenticate(password string) error {
	if authstate.Username == "" {
		return errors.New("No username given")
	} else if password == "" {
		return errors.New("Empty passwords are not allowed")
	}

	if password == authstate.Users[authstate.Username] {
		authstate.Authenticated = true
		return nil
	} else {
		return errors.New("Password missmatch")
	}
}
