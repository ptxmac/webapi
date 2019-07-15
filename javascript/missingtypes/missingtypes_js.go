// Code generated by webidl-bind. DO NOT EDIT.

package missingtypes

import "syscall/js"

// using following types:

// source idl files:
// missingtypes.idl

// transform files:
// missingtypes.go.md

// ReleasableApiResource is used to release underlaying
// allocated resources.
type ReleasableApiResource interface {
	Release()
}

type releasableApiResourceList []ReleasableApiResource

func (a releasableApiResourceList) Release() {
	for _, v := range a {
		v.Release()
	}
}

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) *Union {
	return &Union{Value: value}
}

// interface: Coordinates
type Coordinates struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Coordinates) JSValue() js.Value {
	return _this.Value_JS
}

// CoordinatesFromJS is casting a js.Wrapper into Coordinates.
func CoordinatesFromJS(value js.Wrapper) *Coordinates {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Coordinates{}
	ret.Value_JS = input
	return ret
}

// interface: Date
type Date struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Date) JSValue() js.Value {
	return _this.Value_JS
}

// DateFromJS is casting a js.Wrapper into Date.
func DateFromJS(value js.Wrapper) *Date {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Date{}
	ret.Value_JS = input
	return ret
}

// interface: Dictionary
type Dictionary struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Dictionary) JSValue() js.Value {
	return _this.Value_JS
}

// DictionaryFromJS is casting a js.Wrapper into Dictionary.
func DictionaryFromJS(value js.Wrapper) *Dictionary {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Dictionary{}
	ret.Value_JS = input
	return ret
}

// interface: WritableStream
type WritableStream struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *WritableStream) JSValue() js.Value {
	return _this.Value_JS
}

// WritableStreamFromJS is casting a js.Wrapper into WritableStream.
func WritableStreamFromJS(value js.Wrapper) *WritableStream {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &WritableStream{}
	ret.Value_JS = input
	return ret
}
