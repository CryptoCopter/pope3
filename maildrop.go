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
	"crypto/md5"
	"encoding/json"
	"fmt"
	"hash"
	"io/ioutil"
)

type Maildrop struct {
	Count    int
	Size     int
	Messages []Message
}

type Message struct {
	Size    int
	Display string
	UID     []byte
	Content MessageContent
}

type MessageContent struct {
	Subject string
	Body    string
}

func LoadMaildrop(filename string) (Maildrop, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Maildrop{}, err
	}

	var messagesContent []MessageContent
	err = json.Unmarshal(data, &messagesContent)
	if err != nil {
		return Maildrop{}, err
	}

	var size = 0
	var messages []Message
	var messageSize int
	var messageDisplay string
	var hasher hash.Hash
	var hash []byte

	for _, message := range messagesContent {
		messageDisplay = fmt.Sprintf("Subject: %v\n\n%v", message.Subject, message.Body)
		messageSize = len([]byte(messageDisplay))
		size += messageSize

		hasher = md5.New()
		hasher.Write([]byte(messageDisplay))
		hash = hasher.Sum(nil)

		messages = append(messages, Message{Size: messageSize, Display: messageDisplay, UID: hash, Content: message})
	}

	count := len(messages)

	return Maildrop{Count: count, Size: size, Messages: messages}, nil
}
