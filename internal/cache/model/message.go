/*
  Copyright (C) 2019 - 2020 MWSOFT
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
*/
package model

import "encoding/json"

// Message holds the data received from client.
type Message struct {
	SenderID   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
	Message    string `json:"message"`
	CreatedAt  string `json:"createdAt"`
}

// MarshalBinary ...
func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

// UnmarshalBinary ...
func (m *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &m)
}
