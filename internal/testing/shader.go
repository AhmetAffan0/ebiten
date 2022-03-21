// Copyright 2020 The Ebiten Authors
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

package testing

import (
	"fmt"
	"strings"
)

// ShaderProgramFill returns a shader source to fill the frambuffer.
func ShaderProgramFill(r, g, b, a byte) []byte {
	return []byte(fmt.Sprintf(`package main

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	return vec4(%0.9f, %0.9f, %0.9f, %0.9f)
}
`, float64(r)/0xff, float64(g)/0xff, float64(b)/0xff, float64(a)/0xff))
}

// ShaderProgramImages returns a shader source to render the frambuffer with the given images.
func ShaderProgramImages(numImages int) []byte {
	if numImages <= 0 {
		panic("testing: numImages must be >= 1")
	}

	var exprs []string
	for i := 0; i < numImages; i++ {
		exprs = append(exprs, fmt.Sprintf("imageSrc%dUnsafeAt(texCoord)", i))
	}

	return []byte(fmt.Sprintf(`package main

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	return %s
}
`, strings.Join(exprs, " + ")))
}
