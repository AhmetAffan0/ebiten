// Copyright 2021 The Ebiten Authors
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

package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/internal/driver"
)

// UpdateGamepads updates the gamepad states.
// UpdateGamepads is called when the gamepad states are given from outside like Android.
func (u *UserInterface) UpdateGamepads(gamepads []Gamepad) {
	u.input.updateGamepads(gamepads)
	if u.fpsMode == driver.FPSModeVsyncOffMinimum {
		u.renderRequester.RequestRenderIfNeeded()
	}
}

func (i *Input) updateGamepads(gamepads []Gamepad) {
	i.gamepads = i.gamepads[:0]
	i.gamepads = append(i.gamepads, gamepads...)
}
