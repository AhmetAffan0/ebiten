// Copyright 2016 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build android || ios
// +build android ios

package ebitenmobileview

import (
	"github.com/hajimehoshi/ebiten/v2/internal/driver"
	"github.com/hajimehoshi/ebiten/v2/internal/uidriver/mobile"
)

type position struct {
	x int
	y int
}

var (
	keys     = map[driver.Key]struct{}{}
	runes    []rune
	touches  = map[driver.TouchID]position{}
	gamepads = map[driver.GamepadID]mobile.Gamepad{}
)

var (
	touchSlice   []mobile.Touch
	gamepadSlice []mobile.Gamepad
)

func updateInput() {
	touchSlice = touchSlice[:0]
	for id, position := range touches {
		touchSlice = append(touchSlice, mobile.Touch{
			ID: id,
			X:  position.x,
			Y:  position.y,
		})
	}

	gamepadSlice = gamepadSlice[:0]
	for _, g := range gamepads {
		gamepadSlice = append(gamepadSlice, g)
	}

	mobile.Get().UpdateInput(keys, runes, touchSlice, gamepadSlice)
}
