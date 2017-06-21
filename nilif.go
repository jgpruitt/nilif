/*
MIT License

Copyright (c) *2017 John Pruitt

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

package nilif

func Bool(arg bool, val bool) *bool {
	if arg == val {
		return nil
	}
	return &arg
}

func Rune(arg rune, val rune) *rune {
	if arg == val {
		return nil
	}
	return &arg
}

func String(arg string, val string) *string {
	if arg == val {
		return nil
	}
	return &arg
}

func Int(arg int, val int) *int {
	if arg == val {
		return nil
	}
	return &arg
}

func Int8(arg int8, val int8) *int8 {
	if arg == val {
		return nil
	}
	return &arg
}

func Int16(arg int16, val int16) *int16 {
	if arg == val {
		return nil
	}
	return &arg
}

func Int32(arg int32, val int32) *int32 {
	if arg == val {
		return nil
	}
	return &arg
}

func Int64(arg int64, val int64) *int64 {
	if arg == val {
		return nil
	}
	return &arg
}

func Uint(arg uint, val uint) *uint {
	if arg == val {
		return nil
	}
	return &arg
}

func Uint8(arg uint8, val uint8) *uint8 {
	if arg == val {
		return nil
	}
	return &arg
}

func Uint16(arg uint16, val uint16) *uint16 {
	if arg == val {
		return nil
	}
	return &arg
}

func Uint32(arg uint32, val uint32) *uint32 {
	if arg == val {
		return nil
	}
	return &arg
}

func Uint64(arg uint64, val uint64) *uint64 {
	if arg == val {
		return nil
	}
	return &arg
}

func Float32(arg float32, val float32) *float32 {
	if arg == val {
		return nil
	}
	return &arg
}

func Float64(arg float64, val float64) *float64 {
	if arg == val {
		return nil
	}
	return &arg
}

