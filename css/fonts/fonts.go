// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package fonts

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/css/cssom"
)

// using following types:
// cssom.CSSRule
// cssom.CSSStyleDeclaration

// source idl files:
// css-fonts.idl

// transform files:
// css-fonts.go.md

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

// callback: CSSFontFeatureValuesMapForEach
type CSSFontFeatureValuesMapForEachFunc func(currentValue []uint, currentKey string, listObj *CSSFontFeatureValuesMap)

// CSSFontFeatureValuesMapForEach is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type CSSFontFeatureValuesMapForEach js.Func

func CSSFontFeatureValuesMapForEachToJS(callback CSSFontFeatureValuesMapForEachFunc) *CSSFontFeatureValuesMapForEach {
	if callback == nil {
		return nil
	}
	ret := CSSFontFeatureValuesMapForEach(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 []uint                   // javascript: sequence<unsigned long> currentValue
			_p1 string                   // javascript: DOMString currentKey
			_p2 *CSSFontFeatureValuesMap // javascript: CSSFontFeatureValuesMap listObj
		)
		__length0 := args[0].Length()
		__array0 := make([]uint, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 uint
			__seq_in0 := args[0].Index(__idx0)
			__seq_out0 = (uint)((__seq_in0).Int())
			__array0[__idx0] = __seq_out0
		}
		_p0 = __array0
		_p1 = (args[1]).String()
		_p2 = CSSFontFeatureValuesMapFromJS(args[2])
		callback(_p0, _p1, _p2)

		// returning no return value
		return nil
	}))
	return &ret
}

func CSSFontFeatureValuesMapForEachFromJS(_value js.Value) CSSFontFeatureValuesMapForEachFunc {
	return func(currentValue []uint, currentKey string, listObj *CSSFontFeatureValuesMap) {
		var (
			_args [3]interface{}
			_end  int
		)
		_p0 := js.Global().Get("Array").New(len(currentValue))
		for __idx0, __seq_in0 := range currentValue {
			__seq_out0 := __seq_in0
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
		_p1 := currentKey
		_args[1] = _p1
		_end++
		_p2 := listObj.JSValue()
		_args[2] = _p2
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: CSSFontPaletteValuesRuleForEach
type CSSFontPaletteValuesRuleForEachFunc func(currentValue string, currentKey uint, listObj *CSSFontPaletteValuesRule)

// CSSFontPaletteValuesRuleForEach is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type CSSFontPaletteValuesRuleForEach js.Func

func CSSFontPaletteValuesRuleForEachToJS(callback CSSFontPaletteValuesRuleForEachFunc) *CSSFontPaletteValuesRuleForEach {
	if callback == nil {
		return nil
	}
	ret := CSSFontPaletteValuesRuleForEach(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 string                    // javascript: DOMString currentValue
			_p1 uint                      // javascript: unsigned long currentKey
			_p2 *CSSFontPaletteValuesRule // javascript: CSSFontPaletteValuesRule listObj
		)
		_p0 = (args[0]).String()
		_p1 = (uint)((args[1]).Int())
		_p2 = CSSFontPaletteValuesRuleFromJS(args[2])
		callback(_p0, _p1, _p2)

		// returning no return value
		return nil
	}))
	return &ret
}

func CSSFontPaletteValuesRuleForEachFromJS(_value js.Value) CSSFontPaletteValuesRuleForEachFunc {
	return func(currentValue string, currentKey uint, listObj *CSSFontPaletteValuesRule) {
		var (
			_args [3]interface{}
			_end  int
		)
		_p0 := currentValue
		_args[0] = _p0
		_end++
		_p1 := currentKey
		_args[1] = _p1
		_end++
		_p2 := listObj.JSValue()
		_args[2] = _p2
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: CSSFontFeatureValuesMapEntryIteratorValue
type CSSFontFeatureValuesMapEntryIteratorValue struct {
	Value []js.Value
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CSSFontFeatureValuesMapEntryIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Value))
	for __idx0, __seq_in0 := range _this.Value {
		__seq_out0 := __seq_in0
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// CSSFontFeatureValuesMapEntryIteratorValueFromJS is allocating a new
// CSSFontFeatureValuesMapEntryIteratorValue object and copy all values from
// input javascript object
func CSSFontFeatureValuesMapEntryIteratorValueFromJS(value js.Wrapper) *CSSFontFeatureValuesMapEntryIteratorValue {
	input := value.JSValue()
	var out CSSFontFeatureValuesMapEntryIteratorValue
	var (
		value0 []js.Value // javascript: sequence<any> {value Value value}
		value1 bool       // javascript: boolean {done Done done}
	)
	__length0 := input.Get("value").Length()
	__array0 := make([]js.Value, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 js.Value
		__seq_in0 := input.Get("value").Index(__idx0)
		__seq_out0 = __seq_in0
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// dictionary: CSSFontFeatureValuesMapKeyIteratorValue
type CSSFontFeatureValuesMapKeyIteratorValue struct {
	Value string
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CSSFontFeatureValuesMapKeyIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Value
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// CSSFontFeatureValuesMapKeyIteratorValueFromJS is allocating a new
// CSSFontFeatureValuesMapKeyIteratorValue object and copy all values from
// input javascript object
func CSSFontFeatureValuesMapKeyIteratorValueFromJS(value js.Wrapper) *CSSFontFeatureValuesMapKeyIteratorValue {
	input := value.JSValue()
	var out CSSFontFeatureValuesMapKeyIteratorValue
	var (
		value0 string // javascript: DOMString {value Value value}
		value1 bool   // javascript: boolean {done Done done}
	)
	value0 = (input.Get("value")).String()
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// dictionary: CSSFontFeatureValuesMapValueIteratorValue
type CSSFontFeatureValuesMapValueIteratorValue struct {
	Value []uint
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CSSFontFeatureValuesMapValueIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Value))
	for __idx0, __seq_in0 := range _this.Value {
		__seq_out0 := __seq_in0
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// CSSFontFeatureValuesMapValueIteratorValueFromJS is allocating a new
// CSSFontFeatureValuesMapValueIteratorValue object and copy all values from
// input javascript object
func CSSFontFeatureValuesMapValueIteratorValueFromJS(value js.Wrapper) *CSSFontFeatureValuesMapValueIteratorValue {
	input := value.JSValue()
	var out CSSFontFeatureValuesMapValueIteratorValue
	var (
		value0 []uint // javascript: sequence<unsigned long> {value Value value}
		value1 bool   // javascript: boolean {done Done done}
	)
	__length0 := input.Get("value").Length()
	__array0 := make([]uint, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 uint
		__seq_in0 := input.Get("value").Index(__idx0)
		__seq_out0 = (uint)((__seq_in0).Int())
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// dictionary: CSSFontPaletteValuesRuleEntryIteratorValue
type CSSFontPaletteValuesRuleEntryIteratorValue struct {
	Value []js.Value
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CSSFontPaletteValuesRuleEntryIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Value))
	for __idx0, __seq_in0 := range _this.Value {
		__seq_out0 := __seq_in0
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// CSSFontPaletteValuesRuleEntryIteratorValueFromJS is allocating a new
// CSSFontPaletteValuesRuleEntryIteratorValue object and copy all values from
// input javascript object
func CSSFontPaletteValuesRuleEntryIteratorValueFromJS(value js.Wrapper) *CSSFontPaletteValuesRuleEntryIteratorValue {
	input := value.JSValue()
	var out CSSFontPaletteValuesRuleEntryIteratorValue
	var (
		value0 []js.Value // javascript: sequence<any> {value Value value}
		value1 bool       // javascript: boolean {done Done done}
	)
	__length0 := input.Get("value").Length()
	__array0 := make([]js.Value, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 js.Value
		__seq_in0 := input.Get("value").Index(__idx0)
		__seq_out0 = __seq_in0
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// dictionary: CSSFontPaletteValuesRuleKeyIteratorValue
type CSSFontPaletteValuesRuleKeyIteratorValue struct {
	Value uint
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CSSFontPaletteValuesRuleKeyIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Value
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// CSSFontPaletteValuesRuleKeyIteratorValueFromJS is allocating a new
// CSSFontPaletteValuesRuleKeyIteratorValue object and copy all values from
// input javascript object
func CSSFontPaletteValuesRuleKeyIteratorValueFromJS(value js.Wrapper) *CSSFontPaletteValuesRuleKeyIteratorValue {
	input := value.JSValue()
	var out CSSFontPaletteValuesRuleKeyIteratorValue
	var (
		value0 uint // javascript: unsigned long {value Value value}
		value1 bool // javascript: boolean {done Done done}
	)
	value0 = (uint)((input.Get("value")).Int())
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// dictionary: CSSFontPaletteValuesRuleValueIteratorValue
type CSSFontPaletteValuesRuleValueIteratorValue struct {
	Value string
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CSSFontPaletteValuesRuleValueIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Value
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// CSSFontPaletteValuesRuleValueIteratorValueFromJS is allocating a new
// CSSFontPaletteValuesRuleValueIteratorValue object and copy all values from
// input javascript object
func CSSFontPaletteValuesRuleValueIteratorValueFromJS(value js.Wrapper) *CSSFontPaletteValuesRuleValueIteratorValue {
	input := value.JSValue()
	var out CSSFontPaletteValuesRuleValueIteratorValue
	var (
		value0 string // javascript: DOMString {value Value value}
		value1 bool   // javascript: boolean {done Done done}
	)
	value0 = (input.Get("value")).String()
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// class: CSSFontFaceRule
type CSSFontFaceRule struct {
	cssom.CSSRule
}

// CSSFontFaceRuleFromJS is casting a js.Wrapper into CSSFontFaceRule.
func CSSFontFaceRuleFromJS(value js.Wrapper) *CSSFontFaceRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontFaceRule{}
	ret.Value_JS = input
	return ret
}

// Style returning attribute 'style' with
// type cssom.CSSStyleDeclaration (idl: CSSStyleDeclaration).
func (_this *CSSFontFaceRule) Style() *cssom.CSSStyleDeclaration {
	var ret *cssom.CSSStyleDeclaration
	value := _this.Value_JS.Get("style")
	ret = cssom.CSSStyleDeclarationFromJS(value)
	return ret
}

// class: CSSFontFeatureValuesMap
type CSSFontFeatureValuesMap struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSFontFeatureValuesMap) JSValue() js.Value {
	return _this.Value_JS
}

// CSSFontFeatureValuesMapFromJS is casting a js.Wrapper into CSSFontFeatureValuesMap.
func CSSFontFeatureValuesMapFromJS(value js.Wrapper) *CSSFontFeatureValuesMap {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontFeatureValuesMap{}
	ret.Value_JS = input
	return ret
}

// Size returning attribute 'size' with
// type int (idl: long).
func (_this *CSSFontFeatureValuesMap) Size() int {
	var ret int
	value := _this.Value_JS.Get("size")
	ret = (value).Int()
	return ret
}

func (_this *CSSFontFeatureValuesMap) Set(featureValueName string, values *Union) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := featureValueName
	_args[0] = _p0
	_end++
	_p1 := values.JSValue()
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("set", _args[0:_end]...)
	return
}

func (_this *CSSFontFeatureValuesMap) Entries() (_result *CSSFontFeatureValuesMapEntryIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("entries", _args[0:_end]...)
	var (
		_converted *CSSFontFeatureValuesMapEntryIterator // javascript: CSSFontFeatureValuesMapEntryIterator _what_return_name
	)
	_converted = CSSFontFeatureValuesMapEntryIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSFontFeatureValuesMap) ForEach(callback *CSSFontFeatureValuesMapForEach, optionalThisForCallbackArgument interface{}) {
	var (
		_args [2]interface{}
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
	if optionalThisForCallbackArgument != nil {
		_p1 := optionalThisForCallbackArgument
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("forEach", _args[0:_end]...)
	return
}

func (_this *CSSFontFeatureValuesMap) Keys() (_result *CSSFontFeatureValuesMapKeyIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("keys", _args[0:_end]...)
	var (
		_converted *CSSFontFeatureValuesMapKeyIterator // javascript: CSSFontFeatureValuesMapKeyIterator _what_return_name
	)
	_converted = CSSFontFeatureValuesMapKeyIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSFontFeatureValuesMap) Values() (_result *CSSFontFeatureValuesMapValueIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("values", _args[0:_end]...)
	var (
		_converted *CSSFontFeatureValuesMapValueIterator // javascript: CSSFontFeatureValuesMapValueIterator _what_return_name
	)
	_converted = CSSFontFeatureValuesMapValueIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSFontFeatureValuesMap) Get(key string) (_result []uint) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := key
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted []uint // javascript: sequence<unsigned long> _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		__length0 := _returned.Length()
		__array0 := make([]uint, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 uint
			__seq_in0 := _returned.Index(__idx0)
			__seq_out0 = (uint)((__seq_in0).Int())
			__array0[__idx0] = __seq_out0
		}
		_converted = __array0
	}
	_result = _converted
	return
}

func (_this *CSSFontFeatureValuesMap) Has(key string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := key
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("has", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *CSSFontFeatureValuesMap) Clear() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("clear", _args[0:_end]...)
	return
}

func (_this *CSSFontFeatureValuesMap) Delete(key string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := key
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("delete", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

// class: CSSFontFeatureValuesMapEntryIterator
type CSSFontFeatureValuesMapEntryIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSFontFeatureValuesMapEntryIterator) JSValue() js.Value {
	return _this.Value_JS
}

// CSSFontFeatureValuesMapEntryIteratorFromJS is casting a js.Wrapper into CSSFontFeatureValuesMapEntryIterator.
func CSSFontFeatureValuesMapEntryIteratorFromJS(value js.Wrapper) *CSSFontFeatureValuesMapEntryIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontFeatureValuesMapEntryIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *CSSFontFeatureValuesMapEntryIterator) Next() (_result *CSSFontFeatureValuesMapEntryIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *CSSFontFeatureValuesMapEntryIteratorValue // javascript: CSSFontFeatureValuesMapEntryIteratorValue _what_return_name
	)
	_converted = CSSFontFeatureValuesMapEntryIteratorValueFromJS(_returned)
	_result = _converted
	return
}

// class: CSSFontFeatureValuesMapKeyIterator
type CSSFontFeatureValuesMapKeyIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSFontFeatureValuesMapKeyIterator) JSValue() js.Value {
	return _this.Value_JS
}

// CSSFontFeatureValuesMapKeyIteratorFromJS is casting a js.Wrapper into CSSFontFeatureValuesMapKeyIterator.
func CSSFontFeatureValuesMapKeyIteratorFromJS(value js.Wrapper) *CSSFontFeatureValuesMapKeyIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontFeatureValuesMapKeyIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *CSSFontFeatureValuesMapKeyIterator) Next() (_result *CSSFontFeatureValuesMapKeyIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *CSSFontFeatureValuesMapKeyIteratorValue // javascript: CSSFontFeatureValuesMapKeyIteratorValue _what_return_name
	)
	_converted = CSSFontFeatureValuesMapKeyIteratorValueFromJS(_returned)
	_result = _converted
	return
}

// class: CSSFontFeatureValuesMapValueIterator
type CSSFontFeatureValuesMapValueIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSFontFeatureValuesMapValueIterator) JSValue() js.Value {
	return _this.Value_JS
}

// CSSFontFeatureValuesMapValueIteratorFromJS is casting a js.Wrapper into CSSFontFeatureValuesMapValueIterator.
func CSSFontFeatureValuesMapValueIteratorFromJS(value js.Wrapper) *CSSFontFeatureValuesMapValueIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontFeatureValuesMapValueIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *CSSFontFeatureValuesMapValueIterator) Next() (_result *CSSFontFeatureValuesMapValueIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *CSSFontFeatureValuesMapValueIteratorValue // javascript: CSSFontFeatureValuesMapValueIteratorValue _what_return_name
	)
	_converted = CSSFontFeatureValuesMapValueIteratorValueFromJS(_returned)
	_result = _converted
	return
}

// class: CSSFontFeatureValuesRule
type CSSFontFeatureValuesRule struct {
	cssom.CSSRule
}

// CSSFontFeatureValuesRuleFromJS is casting a js.Wrapper into CSSFontFeatureValuesRule.
func CSSFontFeatureValuesRuleFromJS(value js.Wrapper) *CSSFontFeatureValuesRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontFeatureValuesRule{}
	ret.Value_JS = input
	return ret
}

// FontFamily returning attribute 'fontFamily' with
// type string (idl: DOMString).
func (_this *CSSFontFeatureValuesRule) FontFamily() string {
	var ret string
	value := _this.Value_JS.Get("fontFamily")
	ret = (value).String()
	return ret
}

// SetFontFamily setting attribute 'fontFamily' with
// type string (idl: DOMString).
func (_this *CSSFontFeatureValuesRule) SetFontFamily(value string) {
	input := value
	_this.Value_JS.Set("fontFamily", input)
}

// Annotation returning attribute 'annotation' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) Annotation() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("annotation")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// Ornaments returning attribute 'ornaments' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) Ornaments() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("ornaments")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// Stylistic returning attribute 'stylistic' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) Stylistic() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("stylistic")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// Swash returning attribute 'swash' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) Swash() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("swash")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// CharacterVariant returning attribute 'characterVariant' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) CharacterVariant() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("characterVariant")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// Styleset returning attribute 'styleset' with
// type CSSFontFeatureValuesMap (idl: CSSFontFeatureValuesMap).
func (_this *CSSFontFeatureValuesRule) Styleset() *CSSFontFeatureValuesMap {
	var ret *CSSFontFeatureValuesMap
	value := _this.Value_JS.Get("styleset")
	ret = CSSFontFeatureValuesMapFromJS(value)
	return ret
}

// class: CSSFontPaletteValuesRule
type CSSFontPaletteValuesRule struct {
	cssom.CSSRule
}

// CSSFontPaletteValuesRuleFromJS is casting a js.Wrapper into CSSFontPaletteValuesRule.
func CSSFontPaletteValuesRuleFromJS(value js.Wrapper) *CSSFontPaletteValuesRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontPaletteValuesRule{}
	ret.Value_JS = input
	return ret
}

// FontFamily returning attribute 'fontFamily' with
// type string (idl: DOMString).
func (_this *CSSFontPaletteValuesRule) FontFamily() string {
	var ret string
	value := _this.Value_JS.Get("fontFamily")
	ret = (value).String()
	return ret
}

// SetFontFamily setting attribute 'fontFamily' with
// type string (idl: DOMString).
func (_this *CSSFontPaletteValuesRule) SetFontFamily(value string) {
	input := value
	_this.Value_JS.Set("fontFamily", input)
}

// BasePalette returning attribute 'basePalette' with
// type string (idl: DOMString).
func (_this *CSSFontPaletteValuesRule) BasePalette() string {
	var ret string
	value := _this.Value_JS.Get("basePalette")
	ret = (value).String()
	return ret
}

// SetBasePalette setting attribute 'basePalette' with
// type string (idl: DOMString).
func (_this *CSSFontPaletteValuesRule) SetBasePalette(value string) {
	input := value
	_this.Value_JS.Set("basePalette", input)
}

// Size returning attribute 'size' with
// type int (idl: long).
func (_this *CSSFontPaletteValuesRule) Size() int {
	var ret int
	value := _this.Value_JS.Get("size")
	ret = (value).Int()
	return ret
}

func (_this *CSSFontPaletteValuesRule) Entries() (_result *CSSFontPaletteValuesRuleEntryIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("entries", _args[0:_end]...)
	var (
		_converted *CSSFontPaletteValuesRuleEntryIterator // javascript: CSSFontPaletteValuesRuleEntryIterator _what_return_name
	)
	_converted = CSSFontPaletteValuesRuleEntryIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSFontPaletteValuesRule) ForEach(callback *CSSFontPaletteValuesRuleForEach, optionalThisForCallbackArgument interface{}) {
	var (
		_args [2]interface{}
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
	if optionalThisForCallbackArgument != nil {
		_p1 := optionalThisForCallbackArgument
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("forEach", _args[0:_end]...)
	return
}

func (_this *CSSFontPaletteValuesRule) Keys() (_result *CSSFontPaletteValuesRuleKeyIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("keys", _args[0:_end]...)
	var (
		_converted *CSSFontPaletteValuesRuleKeyIterator // javascript: CSSFontPaletteValuesRuleKeyIterator _what_return_name
	)
	_converted = CSSFontPaletteValuesRuleKeyIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSFontPaletteValuesRule) Values() (_result *CSSFontPaletteValuesRuleValueIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("values", _args[0:_end]...)
	var (
		_converted *CSSFontPaletteValuesRuleValueIterator // javascript: CSSFontPaletteValuesRuleValueIterator _what_return_name
	)
	_converted = CSSFontPaletteValuesRuleValueIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *CSSFontPaletteValuesRule) Get(key uint) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := key
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *string // javascript: DOMString _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *CSSFontPaletteValuesRule) Has(key uint) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := key
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("has", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *CSSFontPaletteValuesRule) Clear() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("clear", _args[0:_end]...)
	return
}

func (_this *CSSFontPaletteValuesRule) Delete(key uint) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := key
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("delete", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *CSSFontPaletteValuesRule) Set(key uint, value string) (_result *CSSFontPaletteValuesRule) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := key
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("set", _args[0:_end]...)
	var (
		_converted *CSSFontPaletteValuesRule // javascript: CSSFontPaletteValuesRule _what_return_name
	)
	_converted = CSSFontPaletteValuesRuleFromJS(_returned)
	_result = _converted
	return
}

// class: CSSFontPaletteValuesRuleEntryIterator
type CSSFontPaletteValuesRuleEntryIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSFontPaletteValuesRuleEntryIterator) JSValue() js.Value {
	return _this.Value_JS
}

// CSSFontPaletteValuesRuleEntryIteratorFromJS is casting a js.Wrapper into CSSFontPaletteValuesRuleEntryIterator.
func CSSFontPaletteValuesRuleEntryIteratorFromJS(value js.Wrapper) *CSSFontPaletteValuesRuleEntryIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontPaletteValuesRuleEntryIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *CSSFontPaletteValuesRuleEntryIterator) Next() (_result *CSSFontPaletteValuesRuleEntryIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *CSSFontPaletteValuesRuleEntryIteratorValue // javascript: CSSFontPaletteValuesRuleEntryIteratorValue _what_return_name
	)
	_converted = CSSFontPaletteValuesRuleEntryIteratorValueFromJS(_returned)
	_result = _converted
	return
}

// class: CSSFontPaletteValuesRuleKeyIterator
type CSSFontPaletteValuesRuleKeyIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSFontPaletteValuesRuleKeyIterator) JSValue() js.Value {
	return _this.Value_JS
}

// CSSFontPaletteValuesRuleKeyIteratorFromJS is casting a js.Wrapper into CSSFontPaletteValuesRuleKeyIterator.
func CSSFontPaletteValuesRuleKeyIteratorFromJS(value js.Wrapper) *CSSFontPaletteValuesRuleKeyIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontPaletteValuesRuleKeyIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *CSSFontPaletteValuesRuleKeyIterator) Next() (_result *CSSFontPaletteValuesRuleKeyIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *CSSFontPaletteValuesRuleKeyIteratorValue // javascript: CSSFontPaletteValuesRuleKeyIteratorValue _what_return_name
	)
	_converted = CSSFontPaletteValuesRuleKeyIteratorValueFromJS(_returned)
	_result = _converted
	return
}

// class: CSSFontPaletteValuesRuleValueIterator
type CSSFontPaletteValuesRuleValueIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *CSSFontPaletteValuesRuleValueIterator) JSValue() js.Value {
	return _this.Value_JS
}

// CSSFontPaletteValuesRuleValueIteratorFromJS is casting a js.Wrapper into CSSFontPaletteValuesRuleValueIterator.
func CSSFontPaletteValuesRuleValueIteratorFromJS(value js.Wrapper) *CSSFontPaletteValuesRuleValueIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSFontPaletteValuesRuleValueIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *CSSFontPaletteValuesRuleValueIterator) Next() (_result *CSSFontPaletteValuesRuleValueIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *CSSFontPaletteValuesRuleValueIteratorValue // javascript: CSSFontPaletteValuesRuleValueIteratorValue _what_return_name
	)
	_converted = CSSFontPaletteValuesRuleValueIteratorValueFromJS(_returned)
	_result = _converted
	return
}
