// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package client

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.VisibilityState
// javascript.FrozenArray
// javascript.Object
// javascript.PromiseFinally

// source idl files:
// promises.idl
// service-workers.idl

// transform files:
// promises.go.md
// service-workers.go.md

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

// enum: ClientType
type ClientType int

const (
	WindowClientType ClientType = iota
	WorkerClientType
	SharedworkerClientType
	AllClientType
)

var clientTypeToWasmTable = []string{
	"window", "worker", "sharedworker", "all",
}

var clientTypeFromWasmTable = map[string]ClientType{
	"window": WindowClientType, "worker": WorkerClientType, "sharedworker": SharedworkerClientType, "all": AllClientType,
}

// JSValue is converting this enum into a javascript object
func (this *ClientType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ClientType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(clientTypeToWasmTable) {
		return clientTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// ClientTypeFromJS is converting a javascript value into
// a ClientType enum value.
func ClientTypeFromJS(value js.Value) ClientType {
	key := value.String()
	conv, ok := clientTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: FrameType
type FrameType int

const (
	AuxiliaryFrameType FrameType = iota
	TopLevelFrameType
	NestedFrameType
	NoneFrameType
)

var frameTypeToWasmTable = []string{
	"auxiliary", "top-level", "nested", "none",
}

var frameTypeFromWasmTable = map[string]FrameType{
	"auxiliary": AuxiliaryFrameType, "top-level": TopLevelFrameType, "nested": NestedFrameType, "none": NoneFrameType,
}

// JSValue is converting this enum into a javascript object
func (this *FrameType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this FrameType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(frameTypeToWasmTable) {
		return frameTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// FrameTypeFromJS is converting a javascript value into
// a FrameType enum value.
func FrameTypeFromJS(value js.Value) FrameType {
	key := value.String()
	conv, ok := frameTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: PromiseTemplateOnFulfilled
type PromiseNilWindowClientOnFulfilledFunc func(value *WindowClient)

// PromiseNilWindowClientOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseNilWindowClientOnFulfilled js.Func

func PromiseNilWindowClientOnFulfilledToJS(callback PromiseNilWindowClientOnFulfilledFunc) *PromiseNilWindowClientOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseNilWindowClientOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *WindowClient // javascript: WindowClient value
		)
		_p0 = WindowClientFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseNilWindowClientOnFulfilledFromJS(_value js.Value) PromiseNilWindowClientOnFulfilledFunc {
	return func(value *WindowClient) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := value.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromiseNilWindowClientOnRejectedFunc func(reason js.Value)

// PromiseNilWindowClientOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseNilWindowClientOnRejected js.Func

func PromiseNilWindowClientOnRejectedToJS(callback PromiseNilWindowClientOnRejectedFunc) *PromiseNilWindowClientOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseNilWindowClientOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 js.Value // javascript: any reason
		)
		_p0 = args[0]
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseNilWindowClientOnRejectedFromJS(_value js.Value) PromiseNilWindowClientOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnFulfilled
type PromiseWindowClientOnFulfilledFunc func(value *WindowClient)

// PromiseWindowClientOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseWindowClientOnFulfilled js.Func

func PromiseWindowClientOnFulfilledToJS(callback PromiseWindowClientOnFulfilledFunc) *PromiseWindowClientOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseWindowClientOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *WindowClient // javascript: WindowClient value
		)
		_p0 = WindowClientFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseWindowClientOnFulfilledFromJS(_value js.Value) PromiseWindowClientOnFulfilledFunc {
	return func(value *WindowClient) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := value.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromiseWindowClientOnRejectedFunc func(reason js.Value)

// PromiseWindowClientOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseWindowClientOnRejected js.Func

func PromiseWindowClientOnRejectedToJS(callback PromiseWindowClientOnRejectedFunc) *PromiseWindowClientOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseWindowClientOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 js.Value // javascript: any reason
		)
		_p0 = args[0]
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseWindowClientOnRejectedFromJS(_value js.Value) PromiseWindowClientOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// class: Client
type Client struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Client) JSValue() js.Value {
	return _this.Value_JS
}

// ClientFromJS is casting a js.Wrapper into Client.
func ClientFromJS(value js.Wrapper) *Client {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Client{}
	ret.Value_JS = input
	return ret
}

// Url returning attribute 'url' with
// type string (idl: USVString).
func (_this *Client) Url() string {
	var ret string
	value := _this.Value_JS.Get("url")
	ret = (value).String()
	return ret
}

// FrameType returning attribute 'frameType' with
// type FrameType (idl: FrameType).
func (_this *Client) FrameType() FrameType {
	var ret FrameType
	value := _this.Value_JS.Get("frameType")
	ret = FrameTypeFromJS(value)
	return ret
}

// Id returning attribute 'id' with
// type string (idl: DOMString).
func (_this *Client) Id() string {
	var ret string
	value := _this.Value_JS.Get("id")
	ret = (value).String()
	return ret
}

// Type returning attribute 'type' with
// type ClientType (idl: ClientType).
func (_this *Client) Type() ClientType {
	var ret ClientType
	value := _this.Value_JS.Get("type")
	ret = ClientTypeFromJS(value)
	return ret
}

func (_this *Client) PostMessage(message interface{}, transfer []*javascript.Object) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := message
	_args[0] = _p0
	_end++
	if transfer != nil {
		_p1 := js.Global().Get("Array").New(len(transfer))
		for __idx1, __seq_in1 := range transfer {
			__seq_out1 := __seq_in1.JSValue()
			_p1.SetIndex(__idx1, __seq_out1)
		}
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("postMessage", _args[0:_end]...)
	return
}

// class: Promise
type PromiseNilWindowClient struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseNilWindowClient) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseNilWindowClientFromJS is casting a js.Wrapper into PromiseNilWindowClient.
func PromiseNilWindowClientFromJS(value js.Wrapper) *PromiseNilWindowClient {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromiseNilWindowClient{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseNilWindowClient) Then(onFulfilled *PromiseNilWindowClientOnFulfilled, onRejected *PromiseNilWindowClientOnRejected) (_result *PromiseNilWindowClient) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("then", _args[0:_end]...)
	var (
		_converted *PromiseNilWindowClient // javascript: Promise _what_return_name
	)
	_converted = PromiseNilWindowClientFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseNilWindowClient) Catch(onRejected *PromiseNilWindowClientOnRejected) (_result *PromiseNilWindowClient) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("catch", _args[0:_end]...)
	var (
		_converted *PromiseNilWindowClient // javascript: Promise _what_return_name
	)
	_converted = PromiseNilWindowClientFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseNilWindowClient) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseNilWindowClient) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("finally", _args[0:_end]...)
	var (
		_converted *PromiseNilWindowClient // javascript: Promise _what_return_name
	)
	_converted = PromiseNilWindowClientFromJS(_returned)
	_result = _converted
	return
}

// class: Promise
type PromiseWindowClient struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseWindowClient) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseWindowClientFromJS is casting a js.Wrapper into PromiseWindowClient.
func PromiseWindowClientFromJS(value js.Wrapper) *PromiseWindowClient {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromiseWindowClient{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseWindowClient) Then(onFulfilled *PromiseWindowClientOnFulfilled, onRejected *PromiseWindowClientOnRejected) (_result *PromiseWindowClient) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("then", _args[0:_end]...)
	var (
		_converted *PromiseWindowClient // javascript: Promise _what_return_name
	)
	_converted = PromiseWindowClientFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseWindowClient) Catch(onRejected *PromiseWindowClientOnRejected) (_result *PromiseWindowClient) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("catch", _args[0:_end]...)
	var (
		_converted *PromiseWindowClient // javascript: Promise _what_return_name
	)
	_converted = PromiseWindowClientFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseWindowClient) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseWindowClient) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("finally", _args[0:_end]...)
	var (
		_converted *PromiseWindowClient // javascript: Promise _what_return_name
	)
	_converted = PromiseWindowClientFromJS(_returned)
	_result = _converted
	return
}

// class: WindowClient
type WindowClient struct {
	Client
}

// WindowClientFromJS is casting a js.Wrapper into WindowClient.
func WindowClientFromJS(value js.Wrapper) *WindowClient {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &WindowClient{}
	ret.Value_JS = input
	return ret
}

// VisibilityState returning attribute 'visibilityState' with
// type domcore.VisibilityState (idl: VisibilityState).
func (_this *WindowClient) VisibilityState() domcore.VisibilityState {
	var ret domcore.VisibilityState
	value := _this.Value_JS.Get("visibilityState")
	ret = domcore.VisibilityStateFromJS(value)
	return ret
}

// Focused returning attribute 'focused' with
// type bool (idl: boolean).
func (_this *WindowClient) Focused() bool {
	var ret bool
	value := _this.Value_JS.Get("focused")
	ret = (value).Bool()
	return ret
}

// AncestorOrigins returning attribute 'ancestorOrigins' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *WindowClient) AncestorOrigins() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("ancestorOrigins")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

func (_this *WindowClient) Focus() (_result *PromiseWindowClient) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("focus", _args[0:_end]...)
	var (
		_converted *PromiseWindowClient // javascript: Promise _what_return_name
	)
	_converted = PromiseWindowClientFromJS(_returned)
	_result = _converted
	return
}

func (_this *WindowClient) Navigate(url string) (_result *PromiseNilWindowClient) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := url
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("navigate", _args[0:_end]...)
	var (
		_converted *PromiseNilWindowClient // javascript: Promise _what_return_name
	)
	_converted = PromiseNilWindowClientFromJS(_returned)
	_result = _converted
	return
}
