// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package gamepad

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.Event
// javascript.FrozenArray

// source idl files:
// gamepad.idl

// transform files:
// gamepad.go.md

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

// enum: GamepadMappingType
type GamepadMappingType int

const (
	EmptyString0GamepadMappingType GamepadMappingType = iota
	StandardGamepadMappingType
)

var gamepadMappingTypeToWasmTable = []string{
	"", "standard",
}

var gamepadMappingTypeFromWasmTable = map[string]GamepadMappingType{
	"": EmptyString0GamepadMappingType, "standard": StandardGamepadMappingType,
}

// JSValue is converting this enum into a javascript object
func (this *GamepadMappingType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this GamepadMappingType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(gamepadMappingTypeToWasmTable) {
		return gamepadMappingTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// GamepadMappingTypeFromJS is converting a javascript value into
// a GamepadMappingType enum value.
func GamepadMappingTypeFromJS(value js.Value) GamepadMappingType {
	key := value.String()
	conv, ok := gamepadMappingTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: GamepadEventInit
type GamepadEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Gamepad    *Gamepad
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *GamepadEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Gamepad.JSValue()
	out.Set("gamepad", value3)
	return out
}

// GamepadEventInitFromJS is allocating a new
// GamepadEventInit object and copy all values in the value javascript object.
func GamepadEventInitFromJS(value js.Value) *GamepadEventInit {
	var out GamepadEventInit
	var (
		value0 bool     // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool     // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool     // javascript: boolean {composed Composed composed}
		value3 *Gamepad // javascript: Gamepad {gamepad Gamepad gamepad}
	)
	value0 = (value.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (value.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (value.Get("composed")).Bool()
	out.Composed = value2
	value3 = GamepadFromJS(value.Get("gamepad"))
	out.Gamepad = value3
	return &out
}

// class: Gamepad
type Gamepad struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Gamepad) JSValue() js.Value {
	return _this.Value_JS
}

// GamepadFromJS is casting a js.Value into Gamepad.
func GamepadFromJS(value js.Value) *Gamepad {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Gamepad{}
	ret.Value_JS = value
	return ret
}

// GamepadFromJS is casting from something that holds a js.Value into Gamepad.
func GamepadFromWrapper(input core.Wrapper) *Gamepad {
	return GamepadFromJS(input.JSValue())
}

// Id returning attribute 'id' with
// type string (idl: DOMString).
func (_this *Gamepad) Id() string {
	var ret string
	value := _this.Value_JS.Get("id")
	ret = (value).String()
	return ret
}

// Index returning attribute 'index' with
// type int (idl: long).
func (_this *Gamepad) Index() int {
	var ret int
	value := _this.Value_JS.Get("index")
	ret = (value).Int()
	return ret
}

// Connected returning attribute 'connected' with
// type bool (idl: boolean).
func (_this *Gamepad) Connected() bool {
	var ret bool
	value := _this.Value_JS.Get("connected")
	ret = (value).Bool()
	return ret
}

// Timestamp returning attribute 'timestamp' with
// type float64 (idl: double).
func (_this *Gamepad) Timestamp() float64 {
	var ret float64
	value := _this.Value_JS.Get("timestamp")
	ret = (value).Float()
	return ret
}

// Mapping returning attribute 'mapping' with
// type GamepadMappingType (idl: GamepadMappingType).
func (_this *Gamepad) Mapping() GamepadMappingType {
	var ret GamepadMappingType
	value := _this.Value_JS.Get("mapping")
	ret = GamepadMappingTypeFromJS(value)
	return ret
}

// Axes returning attribute 'axes' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *Gamepad) Axes() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("axes")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// Buttons returning attribute 'buttons' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *Gamepad) Buttons() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("buttons")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// DisplayId returning attribute 'displayId' with
// type uint (idl: unsigned long).
func (_this *Gamepad) DisplayId() uint {
	var ret uint
	value := _this.Value_JS.Get("displayId")
	ret = (uint)((value).Int())
	return ret
}

// class: GamepadButton
type GamepadButton struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *GamepadButton) JSValue() js.Value {
	return _this.Value_JS
}

// GamepadButtonFromJS is casting a js.Value into GamepadButton.
func GamepadButtonFromJS(value js.Value) *GamepadButton {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &GamepadButton{}
	ret.Value_JS = value
	return ret
}

// GamepadButtonFromJS is casting from something that holds a js.Value into GamepadButton.
func GamepadButtonFromWrapper(input core.Wrapper) *GamepadButton {
	return GamepadButtonFromJS(input.JSValue())
}

// Pressed returning attribute 'pressed' with
// type bool (idl: boolean).
func (_this *GamepadButton) Pressed() bool {
	var ret bool
	value := _this.Value_JS.Get("pressed")
	ret = (value).Bool()
	return ret
}

// Touched returning attribute 'touched' with
// type bool (idl: boolean).
func (_this *GamepadButton) Touched() bool {
	var ret bool
	value := _this.Value_JS.Get("touched")
	ret = (value).Bool()
	return ret
}

// Value returning attribute 'value' with
// type float64 (idl: double).
func (_this *GamepadButton) Value() float64 {
	var ret float64
	value := _this.Value_JS.Get("value")
	ret = (value).Float()
	return ret
}

// class: GamepadEvent
type GamepadEvent struct {
	domcore.Event
}

// GamepadEventFromJS is casting a js.Value into GamepadEvent.
func GamepadEventFromJS(value js.Value) *GamepadEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &GamepadEvent{}
	ret.Value_JS = value
	return ret
}

// GamepadEventFromJS is casting from something that holds a js.Value into GamepadEvent.
func GamepadEventFromWrapper(input core.Wrapper) *GamepadEvent {
	return GamepadEventFromJS(input.JSValue())
}

func NewGamepadEvent(_type string, eventInitDict *GamepadEventInit) (_result *GamepadEvent) {
	_klass := js.Global().Get("GamepadEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := eventInitDict.JSValue()
	_args[1] = _p1
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *GamepadEvent // javascript: GamepadEvent _what_return_name
	)
	_converted = GamepadEventFromJS(_returned)
	_result = _converted
	return
}

// Gamepad returning attribute 'gamepad' with
// type Gamepad (idl: Gamepad).
func (_this *GamepadEvent) Gamepad() *Gamepad {
	var ret *Gamepad
	value := _this.Value_JS.Get("gamepad")
	ret = GamepadFromJS(value)
	return ret
}
