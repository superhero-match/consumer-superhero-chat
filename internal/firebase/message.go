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
package firebase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/consumer-superhero-chat/internal/firebase/model"
)

// PushNewMessageNotification pushes new message notification to Firebase.
func (f *Firebase) PushNewMessageNotification(req model.Request) error {
	requestBody, err := json.Marshal(map[string]string{
		"token":       req.Token,
		"superheroId": req.SuperheroID,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(
		f.FunctionAddress,
		f.ContentType,
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("firebase request returned status: %d", resp.StatusCode)
	}

	return nil
}
