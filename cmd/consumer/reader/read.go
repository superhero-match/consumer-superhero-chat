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
package reader

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"strings"

	cm "github.com/superhero-match/consumer-superhero-chat/internal/cache/model"
	"github.com/superhero-match/consumer-superhero-chat/internal/consumer/model"
	dbm "github.com/superhero-match/consumer-superhero-chat/internal/db/model"
	fm "github.com/superhero-match/consumer-superhero-chat/internal/firebase/model"
)

// Read consumes the Kafka topic and stores the match to DB.
func (r *Reader) Read() error {
	ctx := context.Background()

	for {
		fmt.Println("before FetchMessage")
		m, err := r.Consumer.Consumer.FetchMessage(ctx)
		fmt.Print("after FetchMessage")
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		fmt.Printf(
			"message at topic/partition/offset \n%v/\n%v/\n%v: \n%s = \n%s\n",
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)
		fmt.Print()

		var message model.Message

		if err := json.Unmarshal(m.Value, &message); err != nil {
			fmt.Println("json.Unmarshal")
			fmt.Println(err)
			err = r.Consumer.Consumer.Close()
			if err != nil {
				fmt.Println("r.Consumer.Consumer.Close()")
				fmt.Println(err)
				return err
			}

			return err
		}

		err = r.DB.StoreMessage(dbm.Message{
			ID:         strings.ReplaceAll(uuid.New().String(), "-", ""),
			SenderID:   message.SenderID,
			ReceiverID: message.ReceiverID,
			Message:    message.Message,
			CreatedAt:  message.CreatedAt,
		}, )
		if err != nil {
			fmt.Println("r.DB.StoreMatch")
			fmt.Println(err)

			err = r.Consumer.Consumer.Close()
			if err != nil {
				fmt.Println("r.Consumer.Consumer.Close()")
				fmt.Println(err)

				return err
			}

			return err
		}

		err = r.Cache.StoreMessage(cm.Message{
			SenderID:   message.SenderID,
			ReceiverID: message.ReceiverID,
			Message:    message.Message,
			CreatedAt:  message.CreatedAt,
		})
		if err != nil {
			fmt.Println("r.Cache.StoreMessage")
			fmt.Println(err)

			err = r.Consumer.Consumer.Close()
			if err != nil {
				fmt.Println("r.Consumer.Consumer.Close()")
				fmt.Println(err)

				return err
			}

			return err
		}

		token, err := r.Cache.GetFirebaseMessagingToken(fmt.Sprintf(r.Cache.TokenKeyFormat, message.ReceiverID))
		if err != nil || token == nil {
			fmt.Println("r.Cache.GetFirebaseMessagingToken")
			fmt.Println(err)
			fmt.Println(token)

			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		err = r.Firebase.PushNewMessageNotification(fm.Request{
			Token:       token.Token,
			SuperheroID: message.ReceiverID,
		})
		if err != nil {
			fmt.Println("r.Firebase.PushNewMatchNotification")
			fmt.Println(err)

			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		err = r.Consumer.Consumer.CommitMessages(ctx, m)
		if err != nil {
			fmt.Println("r.Consumer.Consumer.CommitMessages")
			fmt.Println(err)

			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}
	}
}
