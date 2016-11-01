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
	"fmt"
	"strconv"
	"strings"
)

func HandleCommand(command []string, authstate *AuthState, maildrop Maildrop) (string, bool) {
	if command[0] == "" {
		return "-ERR Y U NO GIVE COMMAND?", false
	}

	switch strings.ToLower(command[0]) {
	case "quit":
		return "+OK thxbye", true
	case "user":
		return CaseUser(command, authstate), false
	case "pass":
		return CasePass(command, authstate), false
	case "stat":
		if !authstate.Authenticated {
			return "-ERR You are not authenticated", false
		}
		return CaseStat(maildrop), false
	case "list":
		if !authstate.Authenticated {
			return "-ERR You are not authenticated", false
		}
		return CaseList(command, maildrop), false
	case "retr":
		if !authstate.Authenticated {
			return "-ERR You are not authenticated", false
		}
		return CaseRetr(command, maildrop), false
	case "dele":
		if !authstate.Authenticated {
			return "-ERR You are not authenticated", false
		}
		return "-ERR We do not support message deletion", false
	case "noop":
		return "+OK", false
	case "rset":
		if !authstate.Authenticated {
			return "-ERR You are not authenticated", false
		}
		return "+OK We didn't want to delete anything anyways...", false
	case "uidl":
		if !authstate.Authenticated {
			return "-ERR You are not authenticated", false
		}
		return CaseUidl(command, maildrop), false
	default:
		return "-ERR No such command", false
	}
}

func CaseUser(command []string, authstate *AuthState) string {
	if len(command) < 2 {
		return "-ERR No username given"
	}

	err := authstate.SetUser(command[1])
	if err != nil {
		return "-ERR Who dis?"
	} else {
		return "+OK o shit waddup"
	}
}

func CasePass(command []string, authstate *AuthState) string {
	if len(command) < 2 {
		return "-ERR No password specified"
	}

	err := authstate.Authenticate(command[1])
	if err != nil {
		return fmt.Sprintf("-ERR %v", err)
	} else {
		return "+OK Access granted"
	}
}

func CaseStat(maildrop Maildrop) string {
	return fmt.Sprintf("+OK %v %v", maildrop.Count, maildrop.Size)
}

func CaseList(command []string, maildrop Maildrop) string {
	if len(command) > 1 {
		index, err := strconv.Atoi(command[1])
		if err != nil {
			return fmt.Sprintf("-ERR %v is not a valid index", command[1])
		}
		index -= 1
		if maildrop.Count <= index || index < 0 {
			return "-ERR Index out of bounds"
		}
		return fmt.Sprintf("+OK %v %v", index+1, maildrop.Messages[index].Size)
	} else {
		var reply = []string{fmt.Sprintf("+OK %v messages (%v octets)", maildrop.Count, maildrop.Size)}
		var index = 1

		for _, message := range maildrop.Messages {
			reply = append(reply, fmt.Sprintf("%v %v", index, message.Size))
			index += 1
		}
		return strings.Join(reply, "\n")
	}
}

func CaseRetr(command []string, maildrop Maildrop) string {
	if len(command) < 2 {
		return "-ERR Specify index"
	}
	index, err := strconv.Atoi(command[1])
	if err != nil {
		return fmt.Sprintf("-ERR %v is not a valid index", command[1])
	}
	index -= 1
	if maildrop.Count <= index || index < 0 {
		return "-ERR Index out of bounds"
	}

	return fmt.Sprintf("+OK %v octets\n%v", maildrop.Messages[index].Size, maildrop.Messages[index].Display)
}

func CaseUidl(command []string, maildrop Maildrop) string {
	if len(command) > 1 {
		index, err := strconv.Atoi(command[1])
		if err != nil {
			return fmt.Sprintf("-ERR %v is not a valid index", command[1])
		}
		index -= 1
		if maildrop.Count <= index || index < 0 {
			return "-ERR Index out of bounds"
		}
		return fmt.Sprintf("+OK %v %x", index+1, maildrop.Messages[index].UID)
	} else {
		var reply = []string{"+OK"}
		var index = 1

		for _, message := range maildrop.Messages {
			reply = append(reply, fmt.Sprintf("%v %x", index, message.UID))
			index += 1
		}
		return strings.Join(reply, "\n")
	}
}
