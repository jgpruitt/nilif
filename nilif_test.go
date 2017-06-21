/*
MIT License

Copyright (c) 2017 John Pruitt

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

package nilif_test

import (
	"nilif"
	"testing"
)

func TestBool1(t *testing.T) {
	if nilif.Bool(true, true) != nil {
		t.Fail()
	}
}

func TestBool2(t *testing.T) {
	if nilif.Bool(true, false) == nil {
		t.Fail()
	}
}

func TestRune1(t *testing.T) {
	if nilif.Rune('a', 'a') != nil {
		t.Fail()
	}
}

func TestRune2(t *testing.T) {
	if nilif.Rune('a', 'b') == nil {
		t.Fail()
	}
}

func TestString1(t *testing.T) {
	if nilif.String("a", "a") != nil {
		t.Fail()
	}
}

func TestString2(t *testing.T) {
	if nilif.String("a", "b") == nil {
		t.Fail()
	}
}

func TestFloat321(t *testing.T) {
	if nilif.Float32(1.0, 1.0) != nil {
		t.Fail()
	}
}

func TestFloat322(t *testing.T) {
	if nilif.Float32(1.0, 2.0) == nil {
		t.Fail()
	}
}

func TestFloat641(t *testing.T) {
	if nilif.Float64(1.0, 1.0) != nil {
		t.Fail()
	}
}

func TestFloat642(t *testing.T) {
	if nilif.Float64(1.0, 2.0) == nil {
		t.Fail()
	}
}

func TestInt1(t *testing.T) {
	if nilif.Int(1, 1) != nil {
		t.Fail()
	}
}

func TestInt2(t *testing.T) {
	if nilif.Int(1, 2) == nil {
		t.Fail()
	}
}

func TestInt81(t *testing.T) {
	if nilif.Int8(1, 1) != nil {
		t.Fail()
	}
}

func TestInt82(t *testing.T) {
	if nilif.Int8(1, 2) == nil {
		t.Fail()
	}
}

func TestInt161(t *testing.T) {
	if nilif.Int16(1, 1) != nil {
		t.Fail()
	}
}

func TestInt162(t *testing.T) {
	if nilif.Int16(1, 2) == nil {
		t.Fail()
	}
}

func TestInt321(t *testing.T) {
	if nilif.Int32(1, 1) != nil {
		t.Fail()
	}
}

func TestInt322(t *testing.T) {
	if nilif.Int32(1, 2) == nil {
		t.Fail()
	}
}

func TestInt641(t *testing.T) {
	if nilif.Int64(1, 1) != nil {
		t.Fail()
	}
}

func TestInt642(t *testing.T) {
	if nilif.Int64(1, 2) == nil {
		t.Fail()
	}
}

func TestUint1(t *testing.T) {
	if nilif.Uint(1, 1) != nil {
		t.Fail()
	}
}

func TestUint2(t *testing.T) {
	if nilif.Uint(1, 2) == nil {
		t.Fail()
	}
}

func TestUint81(t *testing.T) {
	if nilif.Uint8(1, 1) != nil {
		t.Fail()
	}
}

func TestUint82(t *testing.T) {
	if nilif.Uint8(1, 2) == nil {
		t.Fail()
	}
}

func TestUint161(t *testing.T) {
	if nilif.Uint16(1, 1) != nil {
		t.Fail()
	}
}

func TestUint162(t *testing.T) {
	if nilif.Uint16(1, 2) == nil {
		t.Fail()
	}
}

func TestUint321(t *testing.T) {
	if nilif.Uint32(1, 1) != nil {
		t.Fail()
	}
}

func TestUint322(t *testing.T) {
	if nilif.Uint32(1, 2) == nil {
		t.Fail()
	}
}

func TestUint641(t *testing.T) {
	if nilif.Uint64(1, 1) != nil {
		t.Fail()
	}
}

func TestUint642(t *testing.T) {
	if nilif.Uint64(1, 2) == nil {
		t.Fail()
	}
}
