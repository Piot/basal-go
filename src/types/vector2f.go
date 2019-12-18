/*

MIT License

Copyright (c) 2017 Peter Bjorklund

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

// Package types ...
package types

import (
	"fmt"
	"math"
)

// Vector2f : Vector type
type Vector2f struct {
	X int32
	Y int32
}

func (v Vector2f) ToFloats() (float32, float32) {
	return FixedToFloat(v.X), FixedToFloat(v.Y)
}

// NewVector2f : Creates a new vector
func NewVector2f(x int32, y int32) Vector2f {
	return Vector2f{X: x, Y: y}
}

func MakeVector2fFromFloats(x float32, y float32) Vector2f {
	return NewVector2f(FloatToFixed(x), FloatToFixed(y))
}

func (v Vector2f) String() string {
	return fmt.Sprintf("[Vector2f %9.2f, %9.2f]", FixedToFloat(v.X), FixedToFloat(v.Y))
}

func (v Vector2f) DebugString() string {
	return fmt.Sprintf("[Vector2f fixed %d, %d]", v.X, v.Y)
}

func (v Vector2f) Delta(o Vector2f) Vector2f {
	return NewVector2f(v.X-o.X, v.Y-o.Y)
}

// DistanceTo calculates the Euclidean distance between two points
func (v Vector2f) DistanceTo(o Vector2f) int32 {
	return int32(math.Sqrt(float64(v.SquaredDistance(o))))
}

// IsEqual checks if both vectors are considered equal
func (v Vector2f) IsEqual(o Vector2f) bool {
	return v.X == o.X && v.Y == o.Y
}

// SquaredDistance calculates the squared Euclidean distance between two points
func (v Vector2f) SquaredDistance(o Vector2f) int64 {
	xd := int64(o.X - v.X)
	yd := int64(o.Y - v.Y)

	return xd*xd + yd*yd
}
