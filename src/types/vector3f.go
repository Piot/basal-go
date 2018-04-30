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

// Vector3f : Vector type
type Vector3f struct {
	X float32
	Y float32
	Z float32
}

// NewVector3f : Creates a new vector
func NewVector3f(x float32, y float32, z float32) Vector3f {
	return Vector3f{X: x, Y: y, Z: z}
}

func (v Vector3f) String() string {
	return fmt.Sprintf("[vector3f %0.2f, %0.2f, %0.2f]", v.X, v.Y, v.Z)
}

// DistanceTo calculates the Euclidean distance between two points
func (v Vector3f) DistanceTo(o Vector3f) float32 {
	return float32(math.Sqrt(float64(v.SquaredDistance(o))))
}

// SquaredDistance calculates the squared Euclidean distance between two points
func (v Vector3f) SquaredDistance(o Vector3f) float32 {
	xd := o.X - v.X
	yd := o.Y - v.Y
	zd := o.Z - v.Z

	return xd*xd + yd*yd + zd*zd
}
