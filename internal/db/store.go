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
package db

import (
	"github.com/superhero-match/consumer-superhero-chat/internal/db/model"
)

//'test_id',
//	'sender_id',
//	'receiver_id',
//	'message',
//	'2019-09-19'

// StoreMatch saves new match.
func(db *DB) StoreMessage (m model.Message) error {
	_, err := db.stmtInsertNewMessage.Exec(
		m.ID,
		m.SenderID,
		m.ReceiverID,
		m.Message,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}


