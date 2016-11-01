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
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type Server struct {
	Address  string
	Port     int
	Timeout  int
	Users    map[string]string
	Maildrop Maildrop
}

func (server *Server) Listen() error {
	incoming, err := net.Listen("tcp", fmt.Sprintf("%v:%v", server.Address, server.Port))
	if err != nil {
		return err
	}

	for {
		connection, err := incoming.Accept()
		if err != nil {
			fmt.Printf("Connection error: %v\n", err)
			continue
		}
		fmt.Printf("Incoming connection from %v\n", connection.RemoteAddr())
		go server.HandleClient(connection)
	}
}

func (server *Server) HandleClient(connection net.Conn) {
	authstate := &AuthState{Username: "", Authenticated: false, Users: server.Users}

	connection.SetDeadline(time.Now().Add(time.Duration(server.Timeout) * time.Second))
	_, err := connection.Write([]byte("+OK O Hai!\n"))
	if err != nil {
		connection.Close()
		fmt.Printf("Connection to %v terminated: %v\n", connection.RemoteAddr(), err)
		return
	}

	reader := bufio.NewReader(connection)

	for {
		command, err := ReadCommand(*reader)
		if err != nil {
			connection.Close()
			fmt.Printf("Connection to %v terminated: %v\n", connection.RemoteAddr(), err)
			return
		}
		connection.SetDeadline(time.Now().Add(time.Duration(server.Timeout) * time.Second))
		fmt.Printf("Received from %v: %#v\n", connection.RemoteAddr(), command)

		reply, closeConnection := HandleCommand(command, authstate, server.Maildrop)
		connection.Write([]byte(reply + "\n"))
		if closeConnection {
			connection.Close()
			fmt.Printf("%v quit", connection.RemoteAddr())
			return
		}
	}
}

func ReadCommand(reader bufio.Reader) (command []string, err error) {
	data, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimRight(data, "\r\n"), " "), nil
}
