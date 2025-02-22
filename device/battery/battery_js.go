// Code generated by webidl-bind. DO NOT EDIT.

package battery

import "syscall/js"

import (
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// javascript.PromiseFinally

// source idl files:
// battery-status.idl
// promises.idl

// transform files:
// battery-status.go.md
// promises.go.md

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

// callback: PromiseTemplateOnFulfilled
type PromiseBatteryManagerOnFulfilledFunc func(value *BatteryManager)

// PromiseBatteryManagerOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseBatteryManagerOnFulfilled js.Func

func PromiseBatteryManagerOnFulfilledToJS(callback PromiseBatteryManagerOnFulfilledFunc) *PromiseBatteryManagerOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseBatteryManagerOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *BatteryManager // javascript: BatteryManager value
		)
		_p0 = BatteryManagerFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseBatteryManagerOnFulfilledFromJS(_value js.Value) PromiseBatteryManagerOnFulfilledFunc {
	return func(value *BatteryManager) {
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
type PromiseBatteryManagerOnRejectedFunc func(reason js.Value)

// PromiseBatteryManagerOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseBatteryManagerOnRejected js.Func

func PromiseBatteryManagerOnRejectedToJS(callback PromiseBatteryManagerOnRejectedFunc) *PromiseBatteryManagerOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseBatteryManagerOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromiseBatteryManagerOnRejectedFromJS(_value js.Value) PromiseBatteryManagerOnRejectedFunc {
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

// class: BatteryManager
type BatteryManager struct {
	domcore.EventTarget
}

// BatteryManagerFromJS is casting a js.Value into BatteryManager.
func BatteryManagerFromJS(value js.Value) *BatteryManager {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &BatteryManager{}
	ret.Value_JS = value
	return ret
}

// BatteryManagerFromJS is casting from something that holds a js.Value into BatteryManager.
func BatteryManagerFromWrapper(input core.Wrapper) *BatteryManager {
	return BatteryManagerFromJS(input.JSValue())
}

// Charging returning attribute 'charging' with
// type bool (idl: boolean).
func (_this *BatteryManager) Charging() bool {
	var ret bool
	value := _this.Value_JS.Get("charging")
	ret = (value).Bool()
	return ret
}

// ChargingTime returning attribute 'chargingTime' with
// type float64 (idl: unrestricted double).
func (_this *BatteryManager) ChargingTime() float64 {
	var ret float64
	value := _this.Value_JS.Get("chargingTime")
	ret = (value).Float()
	return ret
}

// DischargingTime returning attribute 'dischargingTime' with
// type float64 (idl: unrestricted double).
func (_this *BatteryManager) DischargingTime() float64 {
	var ret float64
	value := _this.Value_JS.Get("dischargingTime")
	ret = (value).Float()
	return ret
}

// Level returning attribute 'level' with
// type float64 (idl: double).
func (_this *BatteryManager) Level() float64 {
	var ret float64
	value := _this.Value_JS.Get("level")
	ret = (value).Float()
	return ret
}

// OnChargingChange returning attribute 'onchargingchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *BatteryManager) OnChargingChange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onchargingchange")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnChargingTimeChange returning attribute 'onchargingtimechange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *BatteryManager) OnChargingTimeChange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onchargingtimechange")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnDischargingTimeChange returning attribute 'ondischargingtimechange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *BatteryManager) OnDischargingTimeChange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("ondischargingtimechange")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnLevelChange returning attribute 'onlevelchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *BatteryManager) OnLevelChange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onlevelchange")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// event attribute: domcore.Event
func eventFuncBatteryManager_domcore_Event(listener func(event *domcore.Event, target *BatteryManager)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *domcore.Event
		value := args[0]
		incoming := value.Get("target")
		ret = domcore.EventFromJS(value)
		src := BatteryManagerFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddChargingChange is adding doing AddEventListener for 'ChargingChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *BatteryManager) AddEventChargingChange(listener func(event *domcore.Event, currentTarget *BatteryManager)) js.Func {
	cb := eventFuncBatteryManager_domcore_Event(listener)
	_this.Value_JS.Call("addEventListener", "chargingchange", cb)
	return cb
}

// SetOnChargingChange is assigning a function to 'onchargingchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *BatteryManager) SetOnChargingChange(listener func(event *domcore.Event, currentTarget *BatteryManager)) js.Func {
	cb := eventFuncBatteryManager_domcore_Event(listener)
	_this.Value_JS.Set("onchargingchange", cb)
	return cb
}

// AddChargingTimeChange is adding doing AddEventListener for 'ChargingTimeChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *BatteryManager) AddEventChargingTimeChange(listener func(event *domcore.Event, currentTarget *BatteryManager)) js.Func {
	cb := eventFuncBatteryManager_domcore_Event(listener)
	_this.Value_JS.Call("addEventListener", "chargingtimechange", cb)
	return cb
}

// SetOnChargingTimeChange is assigning a function to 'onchargingtimechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *BatteryManager) SetOnChargingTimeChange(listener func(event *domcore.Event, currentTarget *BatteryManager)) js.Func {
	cb := eventFuncBatteryManager_domcore_Event(listener)
	_this.Value_JS.Set("onchargingtimechange", cb)
	return cb
}

// AddDischargingTimeChange is adding doing AddEventListener for 'DischargingTimeChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *BatteryManager) AddEventDischargingTimeChange(listener func(event *domcore.Event, currentTarget *BatteryManager)) js.Func {
	cb := eventFuncBatteryManager_domcore_Event(listener)
	_this.Value_JS.Call("addEventListener", "dischargingtimechange", cb)
	return cb
}

// SetOnDischargingTimeChange is assigning a function to 'ondischargingtimechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *BatteryManager) SetOnDischargingTimeChange(listener func(event *domcore.Event, currentTarget *BatteryManager)) js.Func {
	cb := eventFuncBatteryManager_domcore_Event(listener)
	_this.Value_JS.Set("ondischargingtimechange", cb)
	return cb
}

// AddLevelChange is adding doing AddEventListener for 'LevelChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *BatteryManager) AddEventLevelChange(listener func(event *domcore.Event, currentTarget *BatteryManager)) js.Func {
	cb := eventFuncBatteryManager_domcore_Event(listener)
	_this.Value_JS.Call("addEventListener", "levelchange", cb)
	return cb
}

// SetOnLevelChange is assigning a function to 'onlevelchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *BatteryManager) SetOnLevelChange(listener func(event *domcore.Event, currentTarget *BatteryManager)) js.Func {
	cb := eventFuncBatteryManager_domcore_Event(listener)
	_this.Value_JS.Set("onlevelchange", cb)
	return cb
}

// class: Promise
type PromiseBatteryManager struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseBatteryManager) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseBatteryManagerFromJS is casting a js.Value into PromiseBatteryManager.
func PromiseBatteryManagerFromJS(value js.Value) *PromiseBatteryManager {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromiseBatteryManager{}
	ret.Value_JS = value
	return ret
}

// PromiseBatteryManagerFromJS is casting from something that holds a js.Value into PromiseBatteryManager.
func PromiseBatteryManagerFromWrapper(input core.Wrapper) *PromiseBatteryManager {
	return PromiseBatteryManagerFromJS(input.JSValue())
}

func (_this *PromiseBatteryManager) Then(onFulfilled *PromiseBatteryManagerOnFulfilled, onRejected *PromiseBatteryManagerOnRejected) (_result *PromiseBatteryManager) {
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
		_converted *PromiseBatteryManager // javascript: Promise _what_return_name
	)
	_converted = PromiseBatteryManagerFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseBatteryManager) Catch(onRejected *PromiseBatteryManagerOnRejected) (_result *PromiseBatteryManager) {
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
		_converted *PromiseBatteryManager // javascript: Promise _what_return_name
	)
	_converted = PromiseBatteryManagerFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseBatteryManager) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseBatteryManager) {
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
		_converted *PromiseBatteryManager // javascript: Promise _what_return_name
	)
	_converted = PromiseBatteryManagerFromJS(_returned)
	_result = _converted
	return
}
