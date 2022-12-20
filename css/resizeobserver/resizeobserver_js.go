// Code generated by webidl-bind. DO NOT EDIT.

package resizeobserver

import "syscall/js"

import (
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/dom/geometry"
)

// using following types:
// dom.Element
// geometry.DOMRectReadOnly

// source idl files:
// ResizeObserver.idl

// transform files:
// ResizeObserver.go.md

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

// callback: ResizeObserverCallback
type ResizeObserverCallbackFunc func(entries []*ResizeObserverEntry, observer *ResizeObserver)

// ResizeObserverCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type ResizeObserverCallback js.Func

func ResizeObserverCallbackToJS(callback ResizeObserverCallbackFunc) *ResizeObserverCallback {
	if callback == nil {
		return nil
	}
	ret := ResizeObserverCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 []*ResizeObserverEntry // javascript: sequence<ResizeObserverEntry> entries
			_p1 *ResizeObserver        // javascript: ResizeObserver observer
		)
		__length0 := args[0].Length()
		__array0 := make([]*ResizeObserverEntry, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 *ResizeObserverEntry
			__seq_in0 := args[0].Index(__idx0)
			__seq_out0 = ResizeObserverEntryFromJS(__seq_in0)
			__array0[__idx0] = __seq_out0
		}
		_p0 = __array0
		_p1 = ResizeObserverFromJS(args[1])
		callback(_p0, _p1)

		// returning no return value
		return nil
	}))
	return &ret
}

func ResizeObserverCallbackFromJS(_value js.Value) ResizeObserverCallbackFunc {
	return func(entries []*ResizeObserverEntry, observer *ResizeObserver) {
		var (
			_args [2]interface{}
			_end  int
		)
		_p0 := js.Global().Get("Array").New(len(entries))
		for __idx0, __seq_in0 := range entries {
			__seq_out0 := __seq_in0.JSValue()
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
		_p1 := observer.JSValue()
		_args[1] = _p1
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// class: ResizeObservation
type ResizeObservation struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *ResizeObservation) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// ResizeObservationFromJS is casting a js.Value into ResizeObservation.
func ResizeObservationFromJS(value js.Value) *ResizeObservation {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &ResizeObservation{}
	ret.Value_JS = value
	return ret
}

// ResizeObservationFromJS is casting from something that holds a js.Value into ResizeObservation.
func ResizeObservationFromWrapper(input core.Wrapper) *ResizeObservation {
	return ResizeObservationFromJS(input.JSValue())
}

func NewResizeObservation(target *dom.Element) (_result *ResizeObservation) {
	_klass := js.Global().Get("ResizeObservation")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := target.JSValue()
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ResizeObservation // javascript: ResizeObservation _what_return_name
	)
	_converted = ResizeObservationFromJS(_returned)
	_result = _converted
	return
}

// Target returning attribute 'target' with
// type dom.Element (idl: Element).
func (_this *ResizeObservation) Target() *dom.Element {
	var ret *dom.Element
	value := _this.Value_JS.Get("target")
	ret = dom.ElementFromJS(value)
	return ret
}

// BroadcastWidth returning attribute 'broadcastWidth' with
// type float32 (idl: float).
func (_this *ResizeObservation) BroadcastWidth() float32 {
	var ret float32
	value := _this.Value_JS.Get("broadcastWidth")
	ret = (float32)((value).Float())
	return ret
}

// BroadcastHeight returning attribute 'broadcastHeight' with
// type float32 (idl: float).
func (_this *ResizeObservation) BroadcastHeight() float32 {
	var ret float32
	value := _this.Value_JS.Get("broadcastHeight")
	ret = (float32)((value).Float())
	return ret
}

func (_this *ResizeObservation) IsActive() (_result bool) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("isActive", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

// class: ResizeObserver
type ResizeObserver struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *ResizeObserver) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// ResizeObserverFromJS is casting a js.Value into ResizeObserver.
func ResizeObserverFromJS(value js.Value) *ResizeObserver {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &ResizeObserver{}
	ret.Value_JS = value
	return ret
}

// ResizeObserverFromJS is casting from something that holds a js.Value into ResizeObserver.
func ResizeObserverFromWrapper(input core.Wrapper) *ResizeObserver {
	return ResizeObserverFromJS(input.JSValue())
}

func NewResizeObserver(callback *ResizeObserverCallback) (_result *ResizeObserver) {
	_klass := js.Global().Get("ResizeObserver")
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if callback != nil {
		__callback0 = (*callback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ResizeObserver // javascript: ResizeObserver _what_return_name
	)
	_converted = ResizeObserverFromJS(_returned)
	_result = _converted
	return
}

func (_this *ResizeObserver) Observe(target *dom.Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := target.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("observe", _args[0:_end]...)
	return
}

func (_this *ResizeObserver) Unobserve(target *dom.Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := target.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("unobserve", _args[0:_end]...)
	return
}

func (_this *ResizeObserver) Disconnect() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("disconnect", _args[0:_end]...)
	return
}

// class: ResizeObserverEntry
type ResizeObserverEntry struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *ResizeObserverEntry) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// ResizeObserverEntryFromJS is casting a js.Value into ResizeObserverEntry.
func ResizeObserverEntryFromJS(value js.Value) *ResizeObserverEntry {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &ResizeObserverEntry{}
	ret.Value_JS = value
	return ret
}

// ResizeObserverEntryFromJS is casting from something that holds a js.Value into ResizeObserverEntry.
func ResizeObserverEntryFromWrapper(input core.Wrapper) *ResizeObserverEntry {
	return ResizeObserverEntryFromJS(input.JSValue())
}

func NewResizeObserverEntry(target *dom.Element) (_result *ResizeObserverEntry) {
	_klass := js.Global().Get("ResizeObserverEntry")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := target.JSValue()
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ResizeObserverEntry // javascript: ResizeObserverEntry _what_return_name
	)
	_converted = ResizeObserverEntryFromJS(_returned)
	_result = _converted
	return
}

// Target returning attribute 'target' with
// type dom.Element (idl: Element).
func (_this *ResizeObserverEntry) Target() *dom.Element {
	var ret *dom.Element
	value := _this.Value_JS.Get("target")
	ret = dom.ElementFromJS(value)
	return ret
}

// ContentRect returning attribute 'contentRect' with
// type geometry.DOMRectReadOnly (idl: DOMRectReadOnly).
func (_this *ResizeObserverEntry) ContentRect() *geometry.DOMRectReadOnly {
	var ret *geometry.DOMRectReadOnly
	value := _this.Value_JS.Get("contentRect")
	ret = geometry.DOMRectReadOnlyFromJS(value)
	return ret
}
