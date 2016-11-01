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
	"os"

	"github.com/jessevdk/go-flags"
)

func ErrorHandler(err error) {
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	var opts struct {
		UserInfo string `short:"u" long:"user-info" default:"userinfo.json" description:"Path of the JSON file containing the user information."`
		MailDrop string `short:"m" long:"mail-drop" default:"maildrop.json" description:"Path of the JSON file containing the mail."`
		Address  string `short:"a" long:"address"   default:"0.0.0.0"       description:"Address on which the server will listen. (default:0.0.0.0."`
		Port     int    `short:"p" long:"port"      default:"110"           description:"Port on which the server will listen. (default 110)"`
		Timeout  int    `short:"t" long:"timeout"   default:"60"            description:"Time (in seconds) before an inactive connection times out"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	users, err := LoadUsers(opts.UserInfo)
	ErrorHandler(err)
	fmt.Printf("Users: %#v\n", users)

	maildrop, err := LoadMaildrop(opts.MailDrop)
	ErrorHandler(err)
	fmt.Printf("Messages: %#v\n", maildrop)

	server := Server{Address: opts.Address, Port: opts.Port, Timeout: opts.Timeout, Users: users, Maildrop: maildrop}
	err = server.Listen()
	ErrorHandler(err)
}
