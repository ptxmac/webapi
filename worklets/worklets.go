// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package worklets

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/fetch"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// fetch.RequestCredentials
// javascript.PromiseVoid

// source idl files:
// worklets.idl

// transform files:
// worklets.go.md

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

// dictionary: WorkletOptions
type WorkletOptions struct {
	Credentials fetch.RequestCredentials
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *WorkletOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Credentials.JSValue()
	out.Set("credentials", value0)
	return out
}

// WorkletOptionsFromJS is allocating a new
// WorkletOptions object and copy all values in the value javascript object.
func WorkletOptionsFromJS(value js.Value) *WorkletOptions {
	var out WorkletOptions
	var (
		value0 fetch.RequestCredentials // javascript: RequestCredentials {credentials Credentials credentials}
	)
	value0 = fetch.RequestCredentialsFromJS(value.Get("credentials"))
	out.Credentials = value0
	return &out
}

// class: Worklet
type Worklet struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Worklet) JSValue() js.Value {
	return _this.Value_JS
}

// WorkletFromJS is casting a js.Value into Worklet.
func WorkletFromJS(value js.Value) *Worklet {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Worklet{}
	ret.Value_JS = value
	return ret
}

// WorkletFromJS is casting from something that holds a js.Value into Worklet.
func WorkletFromWrapper(input core.Wrapper) *Worklet {
	return WorkletFromJS(input.JSValue())
}

func (_this *Worklet) AddModule(moduleURL string, options *WorkletOptions) (_result *javascript.PromiseVoid) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := moduleURL
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("addModule", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}

// class: WorkletGlobalScope
type WorkletGlobalScope struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *WorkletGlobalScope) JSValue() js.Value {
	return _this.Value_JS
}

// WorkletGlobalScopeFromJS is casting a js.Value into WorkletGlobalScope.
func WorkletGlobalScopeFromJS(value js.Value) *WorkletGlobalScope {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &WorkletGlobalScope{}
	ret.Value_JS = value
	return ret
}

// WorkletGlobalScopeFromJS is casting from something that holds a js.Value into WorkletGlobalScope.
func WorkletGlobalScopeFromWrapper(input core.Wrapper) *WorkletGlobalScope {
	return WorkletGlobalScopeFromJS(input.JSValue())
}
