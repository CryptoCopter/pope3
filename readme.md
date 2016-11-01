# Pope3 - Dummy POP3 server

The most pontifical way to serve email...

## Capabilities
This software is sort-of compliant with the POP3 standard specified in RFC 1939. It does support all non-optional commands, however while several users can be specified in the userinfo-file there is only a single shared maildrop for all. Also, message deletion is not supported, since this was built for an educational context and we didn't want students to troll each other by deleting everything. Thus the DELE command will always result in an error.

At this point, mails only consist of the subject and message body. Headers might be included in the future...

## Usage
User information must be JSON encoded in the form of an array (even if there is only a single user) containing objects with the fileds `Username` and `Password` (both capitalised). Example:

`[{"Username": "foo","Password": "bar"},{"Username": "fizz","Password": "buzz"}]`

By default, pope3 expects a file called `userinfo.json` in the working directory, this behaviour can be changed via the `-u` flag.

Similarly, messages are JSON encoded as an array of objects with the fields `Subject` and `Body`. The default maildrop file is called `maildrop.json` and can be overwritten using the `-m` flag.

For the full list of flags, use `-h`.

## Copyright-stuff
Copyright 2016 Markus Sommer
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

### Includes Software by Jesse van den Kieboom (go-flags)
Copyright (c) 2012 Jesse van den Kieboom. All rights reserved.
Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
     notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
     copyright notice, this list of conditions and the following disclaimer
     in the documentation and/or other materials provided with the
     distribution.
   * Neither the name of Google Inc. nor the names of its
     contributors may be used to endorse or promote products derived from
     this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

## Disclaimer
May contain traces of dank memes.
