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

// FirebaseMessagingToken holds data related to Firebase high priority messaging token.
// This token is used when pushing notifications to clients.
type FirebaseMessagingToken struct {
	Token       string `json:"token"`
	SuperheroID string `json:"superheroID"`
	UpdatedAt   string `json:"updatedAt"`
}

// MarshalBinary ...
func (f FirebaseMessagingToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(f)
}

// UnmarshalBinary ...
func (f *FirebaseMessagingToken) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &f)
}