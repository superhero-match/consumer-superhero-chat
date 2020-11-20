/*
  Copyright (C) 2019 - 2021 MWSOFT
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
package config

// Firebase holds all the Firebase related data.
type Firebase struct {
	FunctionAddress string `env:"FIREBASE_NEW_MESSAGE_FUNCTION_ADDRESS" yaml:"function_address" default:"https://us-central1-superheromatch.cloudfunctions.net/newMessage"`
	ContentType     string `env:"FIREBASE_NEW_MESSAGE_CONTENT_TYPE" yaml:"content_type" default:"application/json"`
}