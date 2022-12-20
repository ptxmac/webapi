// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package payment

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
	"github.com/gowebapi/webapi/payment/request"
	"github.com/gowebapi/webapi/serviceworker/client"
)

// using following types:
// client.PromiseNilWindowClient
// domcore.ExtendableEvent
// javascript.FrozenArray
// javascript.Object
// javascript.Promise
// javascript.PromiseBool
// javascript.PromiseFinally
// javascript.PromiseSequenceString
// javascript.PromiseVoid
// request.PaymentCurrencyAmount
// request.PaymentDetailsModifier
// request.PaymentMethodData

// source idl files:
// payment-handler.idl
// promises.idl

// transform files:
// payment-handler.go.md
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
type PromiseNilPaymentMethodChangeResponseOnFulfilledFunc func(value *PaymentMethodChangeResponse)

// PromiseNilPaymentMethodChangeResponseOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseNilPaymentMethodChangeResponseOnFulfilled js.Func

func PromiseNilPaymentMethodChangeResponseOnFulfilledToJS(callback PromiseNilPaymentMethodChangeResponseOnFulfilledFunc) *PromiseNilPaymentMethodChangeResponseOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseNilPaymentMethodChangeResponseOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *PaymentMethodChangeResponse // javascript: PaymentMethodChangeResponse value
		)
		_p0 = PaymentMethodChangeResponseFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseNilPaymentMethodChangeResponseOnFulfilledFromJS(_value js.Value) PromiseNilPaymentMethodChangeResponseOnFulfilledFunc {
	return func(value *PaymentMethodChangeResponse) {
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
type PromiseNilPaymentMethodChangeResponseOnRejectedFunc func(reason js.Value)

// PromiseNilPaymentMethodChangeResponseOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseNilPaymentMethodChangeResponseOnRejected js.Func

func PromiseNilPaymentMethodChangeResponseOnRejectedToJS(callback PromiseNilPaymentMethodChangeResponseOnRejectedFunc) *PromiseNilPaymentMethodChangeResponseOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseNilPaymentMethodChangeResponseOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromiseNilPaymentMethodChangeResponseOnRejectedFromJS(_value js.Value) PromiseNilPaymentMethodChangeResponseOnRejectedFunc {
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
type PromisePaymentHandlerResponseOnFulfilledFunc func(value *PaymentHandlerResponse)

// PromisePaymentHandlerResponseOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePaymentHandlerResponseOnFulfilled js.Func

func PromisePaymentHandlerResponseOnFulfilledToJS(callback PromisePaymentHandlerResponseOnFulfilledFunc) *PromisePaymentHandlerResponseOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromisePaymentHandlerResponseOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *PaymentHandlerResponse // javascript: PaymentHandlerResponse value
		)
		_p0 = PaymentHandlerResponseFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromisePaymentHandlerResponseOnFulfilledFromJS(_value js.Value) PromisePaymentHandlerResponseOnFulfilledFunc {
	return func(value *PaymentHandlerResponse) {
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
type PromisePaymentHandlerResponseOnRejectedFunc func(reason js.Value)

// PromisePaymentHandlerResponseOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePaymentHandlerResponseOnRejected js.Func

func PromisePaymentHandlerResponseOnRejectedToJS(callback PromisePaymentHandlerResponseOnRejectedFunc) *PromisePaymentHandlerResponseOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromisePaymentHandlerResponseOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromisePaymentHandlerResponseOnRejectedFromJS(_value js.Value) PromisePaymentHandlerResponseOnRejectedFunc {
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

// dictionary: CanMakePaymentEventInit
type CanMakePaymentEventInit struct {
	Bubbles              bool
	Cancelable           bool
	Composed             bool
	TopOrigin            string
	PaymentRequestOrigin string
	MethodData           []*request.PaymentMethodData
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *CanMakePaymentEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.TopOrigin
	out.Set("topOrigin", value3)
	value4 := _this.PaymentRequestOrigin
	out.Set("paymentRequestOrigin", value4)
	value5 := js.Global().Get("Array").New(len(_this.MethodData))
	for __idx5, __seq_in5 := range _this.MethodData {
		__seq_out5 := __seq_in5.JSValue()
		value5.SetIndex(__idx5, __seq_out5)
	}
	out.Set("methodData", value5)
	return out
}

// CanMakePaymentEventInitFromJS is allocating a new
// CanMakePaymentEventInit object and copy all values in the value javascript object.
func CanMakePaymentEventInitFromJS(value js.Value) *CanMakePaymentEventInit {
	var out CanMakePaymentEventInit
	var (
		value0 bool                         // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool                         // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool                         // javascript: boolean {composed Composed composed}
		value3 string                       // javascript: USVString {topOrigin TopOrigin topOrigin}
		value4 string                       // javascript: USVString {paymentRequestOrigin PaymentRequestOrigin paymentRequestOrigin}
		value5 []*request.PaymentMethodData // javascript: sequence<PaymentMethodData> {methodData MethodData methodData}
	)
	value0 = (value.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (value.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (value.Get("composed")).Bool()
	out.Composed = value2
	value3 = (value.Get("topOrigin")).String()
	out.TopOrigin = value3
	value4 = (value.Get("paymentRequestOrigin")).String()
	out.PaymentRequestOrigin = value4
	__length5 := value.Get("methodData").Length()
	__array5 := make([]*request.PaymentMethodData, __length5, __length5)
	for __idx5 := 0; __idx5 < __length5; __idx5++ {
		var __seq_out5 *request.PaymentMethodData
		__seq_in5 := value.Get("methodData").Index(__idx5)
		__seq_out5 = request.PaymentMethodDataFromJS(__seq_in5)
		__array5[__idx5] = __seq_out5
	}
	value5 = __array5
	out.MethodData = value5
	return &out
}

// dictionary: ImageObject
type ImageObject struct {
	Src   string
	Sizes string
	Type  string
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *ImageObject) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Src
	out.Set("src", value0)
	value1 := _this.Sizes
	out.Set("sizes", value1)
	value2 := _this.Type
	out.Set("type", value2)
	return out
}

// ImageObjectFromJS is allocating a new
// ImageObject object and copy all values in the value javascript object.
func ImageObjectFromJS(value js.Value) *ImageObject {
	var out ImageObject
	var (
		value0 string // javascript: USVString {src Src src}
		value1 string // javascript: DOMString {sizes Sizes sizes}
		value2 string // javascript: DOMString {type Type _type}
	)
	value0 = (value.Get("src")).String()
	out.Src = value0
	value1 = (value.Get("sizes")).String()
	out.Sizes = value1
	value2 = (value.Get("type")).String()
	out.Type = value2
	return &out
}

// dictionary: PaymentHandlerResponse
type PaymentHandlerResponse struct {
	MethodName string
	Details    *javascript.Object
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *PaymentHandlerResponse) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.MethodName
	out.Set("methodName", value0)
	value1 := _this.Details.JSValue()
	out.Set("details", value1)
	return out
}

// PaymentHandlerResponseFromJS is allocating a new
// PaymentHandlerResponse object and copy all values in the value javascript object.
func PaymentHandlerResponseFromJS(value js.Value) *PaymentHandlerResponse {
	var out PaymentHandlerResponse
	var (
		value0 string             // javascript: DOMString {methodName MethodName methodName}
		value1 *javascript.Object // javascript: object {details Details details}
	)
	value0 = (value.Get("methodName")).String()
	out.MethodName = value0
	value1 = javascript.ObjectFromJS(value.Get("details"))
	out.Details = value1
	return &out
}

// dictionary: PaymentInstrument
type PaymentInstrument struct {
	Name         string
	Icons        []*ImageObject
	Method       string
	Capabilities *javascript.Object
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *PaymentInstrument) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Name
	out.Set("name", value0)
	value1 := js.Global().Get("Array").New(len(_this.Icons))
	for __idx1, __seq_in1 := range _this.Icons {
		__seq_out1 := __seq_in1.JSValue()
		value1.SetIndex(__idx1, __seq_out1)
	}
	out.Set("icons", value1)
	value2 := _this.Method
	out.Set("method", value2)
	value3 := _this.Capabilities.JSValue()
	out.Set("capabilities", value3)
	return out
}

// PaymentInstrumentFromJS is allocating a new
// PaymentInstrument object and copy all values in the value javascript object.
func PaymentInstrumentFromJS(value js.Value) *PaymentInstrument {
	var out PaymentInstrument
	var (
		value0 string             // javascript: DOMString {name Name name}
		value1 []*ImageObject     // javascript: sequence<ImageObject> {icons Icons icons}
		value2 string             // javascript: DOMString {method Method method}
		value3 *javascript.Object // javascript: object {capabilities Capabilities capabilities}
	)
	value0 = (value.Get("name")).String()
	out.Name = value0
	__length1 := value.Get("icons").Length()
	__array1 := make([]*ImageObject, __length1, __length1)
	for __idx1 := 0; __idx1 < __length1; __idx1++ {
		var __seq_out1 *ImageObject
		__seq_in1 := value.Get("icons").Index(__idx1)
		__seq_out1 = ImageObjectFromJS(__seq_in1)
		__array1[__idx1] = __seq_out1
	}
	value1 = __array1
	out.Icons = value1
	value2 = (value.Get("method")).String()
	out.Method = value2
	value3 = javascript.ObjectFromJS(value.Get("capabilities"))
	out.Capabilities = value3
	return &out
}

// dictionary: PaymentMethodChangeResponse
type PaymentMethodChangeResponse struct {
	Error               string
	Total               *request.PaymentCurrencyAmount
	Modifiers           *javascript.FrozenArray
	PaymentMethodErrors *javascript.Object
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *PaymentMethodChangeResponse) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Error
	out.Set("error", value0)
	value1 := _this.Total.JSValue()
	out.Set("total", value1)
	value2 := _this.Modifiers.JSValue()
	out.Set("modifiers", value2)
	value3 := _this.PaymentMethodErrors.JSValue()
	out.Set("paymentMethodErrors", value3)
	return out
}

// PaymentMethodChangeResponseFromJS is allocating a new
// PaymentMethodChangeResponse object and copy all values in the value javascript object.
func PaymentMethodChangeResponseFromJS(value js.Value) *PaymentMethodChangeResponse {
	var out PaymentMethodChangeResponse
	var (
		value0 string                         // javascript: DOMString {error Error _error}
		value1 *request.PaymentCurrencyAmount // javascript: PaymentCurrencyAmount {total Total total}
		value2 *javascript.FrozenArray        // javascript: FrozenArray {modifiers Modifiers modifiers}
		value3 *javascript.Object             // javascript: object {paymentMethodErrors PaymentMethodErrors paymentMethodErrors}
	)
	value0 = (value.Get("error")).String()
	out.Error = value0
	value1 = request.PaymentCurrencyAmountFromJS(value.Get("total"))
	out.Total = value1
	value2 = javascript.FrozenArrayFromJS(value.Get("modifiers"))
	out.Modifiers = value2
	value3 = javascript.ObjectFromJS(value.Get("paymentMethodErrors"))
	out.PaymentMethodErrors = value3
	return &out
}

// dictionary: PaymentRequestEventInit
type PaymentRequestEventInit struct {
	Bubbles              bool
	Cancelable           bool
	Composed             bool
	TopOrigin            string
	PaymentRequestOrigin string
	PaymentRequestId     string
	MethodData           []*request.PaymentMethodData
	Total                *request.PaymentCurrencyAmount
	Modifiers            []*request.PaymentDetailsModifier
	InstrumentKey        string
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *PaymentRequestEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.TopOrigin
	out.Set("topOrigin", value3)
	value4 := _this.PaymentRequestOrigin
	out.Set("paymentRequestOrigin", value4)
	value5 := _this.PaymentRequestId
	out.Set("paymentRequestId", value5)
	value6 := js.Global().Get("Array").New(len(_this.MethodData))
	for __idx6, __seq_in6 := range _this.MethodData {
		__seq_out6 := __seq_in6.JSValue()
		value6.SetIndex(__idx6, __seq_out6)
	}
	out.Set("methodData", value6)
	value7 := _this.Total.JSValue()
	out.Set("total", value7)
	value8 := js.Global().Get("Array").New(len(_this.Modifiers))
	for __idx8, __seq_in8 := range _this.Modifiers {
		__seq_out8 := __seq_in8.JSValue()
		value8.SetIndex(__idx8, __seq_out8)
	}
	out.Set("modifiers", value8)
	value9 := _this.InstrumentKey
	out.Set("instrumentKey", value9)
	return out
}

// PaymentRequestEventInitFromJS is allocating a new
// PaymentRequestEventInit object and copy all values in the value javascript object.
func PaymentRequestEventInitFromJS(value js.Value) *PaymentRequestEventInit {
	var out PaymentRequestEventInit
	var (
		value0 bool                              // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool                              // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool                              // javascript: boolean {composed Composed composed}
		value3 string                            // javascript: USVString {topOrigin TopOrigin topOrigin}
		value4 string                            // javascript: USVString {paymentRequestOrigin PaymentRequestOrigin paymentRequestOrigin}
		value5 string                            // javascript: DOMString {paymentRequestId PaymentRequestId paymentRequestId}
		value6 []*request.PaymentMethodData      // javascript: sequence<PaymentMethodData> {methodData MethodData methodData}
		value7 *request.PaymentCurrencyAmount    // javascript: PaymentCurrencyAmount {total Total total}
		value8 []*request.PaymentDetailsModifier // javascript: sequence<PaymentDetailsModifier> {modifiers Modifiers modifiers}
		value9 string                            // javascript: DOMString {instrumentKey InstrumentKey instrumentKey}
	)
	value0 = (value.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (value.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (value.Get("composed")).Bool()
	out.Composed = value2
	value3 = (value.Get("topOrigin")).String()
	out.TopOrigin = value3
	value4 = (value.Get("paymentRequestOrigin")).String()
	out.PaymentRequestOrigin = value4
	value5 = (value.Get("paymentRequestId")).String()
	out.PaymentRequestId = value5
	__length6 := value.Get("methodData").Length()
	__array6 := make([]*request.PaymentMethodData, __length6, __length6)
	for __idx6 := 0; __idx6 < __length6; __idx6++ {
		var __seq_out6 *request.PaymentMethodData
		__seq_in6 := value.Get("methodData").Index(__idx6)
		__seq_out6 = request.PaymentMethodDataFromJS(__seq_in6)
		__array6[__idx6] = __seq_out6
	}
	value6 = __array6
	out.MethodData = value6
	value7 = request.PaymentCurrencyAmountFromJS(value.Get("total"))
	out.Total = value7
	__length8 := value.Get("modifiers").Length()
	__array8 := make([]*request.PaymentDetailsModifier, __length8, __length8)
	for __idx8 := 0; __idx8 < __length8; __idx8++ {
		var __seq_out8 *request.PaymentDetailsModifier
		__seq_in8 := value.Get("modifiers").Index(__idx8)
		__seq_out8 = request.PaymentDetailsModifierFromJS(__seq_in8)
		__array8[__idx8] = __seq_out8
	}
	value8 = __array8
	out.Modifiers = value8
	value9 = (value.Get("instrumentKey")).String()
	out.InstrumentKey = value9
	return &out
}

// class: CanMakePaymentEvent
type CanMakePaymentEvent struct {
	domcore.ExtendableEvent
}

// CanMakePaymentEventFromJS is casting a js.Value into CanMakePaymentEvent.
func CanMakePaymentEventFromJS(value js.Value) *CanMakePaymentEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &CanMakePaymentEvent{}
	ret.Value_JS = value
	return ret
}

// CanMakePaymentEventFromJS is casting from something that holds a js.Value into CanMakePaymentEvent.
func CanMakePaymentEventFromWrapper(input core.Wrapper) *CanMakePaymentEvent {
	return CanMakePaymentEventFromJS(input.JSValue())
}

func NewCanMakePaymentEvent(_type string, eventInitDict *CanMakePaymentEventInit) (_result *CanMakePaymentEvent) {
	_klass := js.Global().Get("CanMakePaymentEvent")
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
		_converted *CanMakePaymentEvent // javascript: CanMakePaymentEvent _what_return_name
	)
	_converted = CanMakePaymentEventFromJS(_returned)
	_result = _converted
	return
}

// TopOrigin returning attribute 'topOrigin' with
// type string (idl: USVString).
func (_this *CanMakePaymentEvent) TopOrigin() string {
	var ret string
	value := _this.Value_JS.Get("topOrigin")
	ret = (value).String()
	return ret
}

// PaymentRequestOrigin returning attribute 'paymentRequestOrigin' with
// type string (idl: USVString).
func (_this *CanMakePaymentEvent) PaymentRequestOrigin() string {
	var ret string
	value := _this.Value_JS.Get("paymentRequestOrigin")
	ret = (value).String()
	return ret
}

// MethodData returning attribute 'methodData' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *CanMakePaymentEvent) MethodData() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("methodData")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

func (_this *CanMakePaymentEvent) RespondWith(canMakePaymentResponse *javascript.PromiseBool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := canMakePaymentResponse.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("respondWith", _args[0:_end]...)
	return
}

// class: PaymentInstruments
type PaymentInstruments struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *PaymentInstruments) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// PaymentInstrumentsFromJS is casting a js.Value into PaymentInstruments.
func PaymentInstrumentsFromJS(value js.Value) *PaymentInstruments {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PaymentInstruments{}
	ret.Value_JS = value
	return ret
}

// PaymentInstrumentsFromJS is casting from something that holds a js.Value into PaymentInstruments.
func PaymentInstrumentsFromWrapper(input core.Wrapper) *PaymentInstruments {
	return PaymentInstrumentsFromJS(input.JSValue())
}

func (_this *PaymentInstruments) Delete(instrumentKey string) (_result *javascript.PromiseBool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := instrumentKey
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("delete", _args[0:_end]...)
	var (
		_converted *javascript.PromiseBool // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseBoolFromJS(_returned)
	_result = _converted
	return
}

func (_this *PaymentInstruments) Get(instrumentKey string) (_result *javascript.Promise) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := instrumentKey
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *PaymentInstruments) Keys() (_result *javascript.PromiseSequenceString) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("keys", _args[0:_end]...)
	var (
		_converted *javascript.PromiseSequenceString // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseSequenceStringFromJS(_returned)
	_result = _converted
	return
}

func (_this *PaymentInstruments) Has(instrumentKey string) (_result *javascript.PromiseBool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := instrumentKey
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("has", _args[0:_end]...)
	var (
		_converted *javascript.PromiseBool // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseBoolFromJS(_returned)
	_result = _converted
	return
}

func (_this *PaymentInstruments) Set(instrumentKey string, details *PaymentInstrument) (_result *javascript.PromiseVoid) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := instrumentKey
	_args[0] = _p0
	_end++
	_p1 := details.JSValue()
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("set", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}

func (_this *PaymentInstruments) Clear() (_result *javascript.PromiseVoid) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("clear", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}

// class: PaymentManager
type PaymentManager struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *PaymentManager) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// PaymentManagerFromJS is casting a js.Value into PaymentManager.
func PaymentManagerFromJS(value js.Value) *PaymentManager {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PaymentManager{}
	ret.Value_JS = value
	return ret
}

// PaymentManagerFromJS is casting from something that holds a js.Value into PaymentManager.
func PaymentManagerFromWrapper(input core.Wrapper) *PaymentManager {
	return PaymentManagerFromJS(input.JSValue())
}

// Instruments returning attribute 'instruments' with
// type PaymentInstruments (idl: PaymentInstruments).
func (_this *PaymentManager) Instruments() *PaymentInstruments {
	var ret *PaymentInstruments
	value := _this.Value_JS.Get("instruments")
	ret = PaymentInstrumentsFromJS(value)
	return ret
}

// UserHint returning attribute 'userHint' with
// type string (idl: DOMString).
func (_this *PaymentManager) UserHint() string {
	var ret string
	value := _this.Value_JS.Get("userHint")
	ret = (value).String()
	return ret
}

// SetUserHint setting attribute 'userHint' with
// type string (idl: DOMString).
func (_this *PaymentManager) SetUserHint(value string) {
	input := value
	_this.Value_JS.Set("userHint", input)
}

// class: PaymentRequestEvent
type PaymentRequestEvent struct {
	domcore.ExtendableEvent
}

// PaymentRequestEventFromJS is casting a js.Value into PaymentRequestEvent.
func PaymentRequestEventFromJS(value js.Value) *PaymentRequestEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PaymentRequestEvent{}
	ret.Value_JS = value
	return ret
}

// PaymentRequestEventFromJS is casting from something that holds a js.Value into PaymentRequestEvent.
func PaymentRequestEventFromWrapper(input core.Wrapper) *PaymentRequestEvent {
	return PaymentRequestEventFromJS(input.JSValue())
}

func NewPaymentRequestEvent(_type string, eventInitDict *PaymentRequestEventInit) (_result *PaymentRequestEvent) {
	_klass := js.Global().Get("PaymentRequestEvent")
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
		_converted *PaymentRequestEvent // javascript: PaymentRequestEvent _what_return_name
	)
	_converted = PaymentRequestEventFromJS(_returned)
	_result = _converted
	return
}

// TopOrigin returning attribute 'topOrigin' with
// type string (idl: USVString).
func (_this *PaymentRequestEvent) TopOrigin() string {
	var ret string
	value := _this.Value_JS.Get("topOrigin")
	ret = (value).String()
	return ret
}

// PaymentRequestOrigin returning attribute 'paymentRequestOrigin' with
// type string (idl: USVString).
func (_this *PaymentRequestEvent) PaymentRequestOrigin() string {
	var ret string
	value := _this.Value_JS.Get("paymentRequestOrigin")
	ret = (value).String()
	return ret
}

// PaymentRequestId returning attribute 'paymentRequestId' with
// type string (idl: DOMString).
func (_this *PaymentRequestEvent) PaymentRequestId() string {
	var ret string
	value := _this.Value_JS.Get("paymentRequestId")
	ret = (value).String()
	return ret
}

// MethodData returning attribute 'methodData' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *PaymentRequestEvent) MethodData() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("methodData")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// Total returning attribute 'total' with
// type javascript.Object (idl: object).
func (_this *PaymentRequestEvent) Total() *javascript.Object {
	var ret *javascript.Object
	value := _this.Value_JS.Get("total")
	ret = javascript.ObjectFromJS(value)
	return ret
}

// Modifiers returning attribute 'modifiers' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *PaymentRequestEvent) Modifiers() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("modifiers")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// InstrumentKey returning attribute 'instrumentKey' with
// type string (idl: DOMString).
func (_this *PaymentRequestEvent) InstrumentKey() string {
	var ret string
	value := _this.Value_JS.Get("instrumentKey")
	ret = (value).String()
	return ret
}

// RequestBillingAddress returning attribute 'requestBillingAddress' with
// type bool (idl: boolean).
func (_this *PaymentRequestEvent) RequestBillingAddress() bool {
	var ret bool
	value := _this.Value_JS.Get("requestBillingAddress")
	ret = (value).Bool()
	return ret
}

func (_this *PaymentRequestEvent) OpenWindow(url string) (_result *client.PromiseNilWindowClient) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := url
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("openWindow", _args[0:_end]...)
	var (
		_converted *client.PromiseNilWindowClient // javascript: Promise _what_return_name
	)
	_converted = client.PromiseNilWindowClientFromJS(_returned)
	_result = _converted
	return
}

func (_this *PaymentRequestEvent) ChangePaymentMethod(methodName string, methodDetails *javascript.Object) (_result *PromiseNilPaymentMethodChangeResponse) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := methodName
	_args[0] = _p0
	_end++
	if methodDetails != nil {
		_p1 := methodDetails.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("changePaymentMethod", _args[0:_end]...)
	var (
		_converted *PromiseNilPaymentMethodChangeResponse // javascript: Promise _what_return_name
	)
	_converted = PromiseNilPaymentMethodChangeResponseFromJS(_returned)
	_result = _converted
	return
}

func (_this *PaymentRequestEvent) RespondWith(handlerResponsePromise *PromisePaymentHandlerResponse) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := handlerResponsePromise.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("respondWith", _args[0:_end]...)
	return
}

// class: Promise
type PromiseNilPaymentMethodChangeResponse struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *PromiseNilPaymentMethodChangeResponse) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// PromiseNilPaymentMethodChangeResponseFromJS is casting a js.Value into PromiseNilPaymentMethodChangeResponse.
func PromiseNilPaymentMethodChangeResponseFromJS(value js.Value) *PromiseNilPaymentMethodChangeResponse {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromiseNilPaymentMethodChangeResponse{}
	ret.Value_JS = value
	return ret
}

// PromiseNilPaymentMethodChangeResponseFromJS is casting from something that holds a js.Value into PromiseNilPaymentMethodChangeResponse.
func PromiseNilPaymentMethodChangeResponseFromWrapper(input core.Wrapper) *PromiseNilPaymentMethodChangeResponse {
	return PromiseNilPaymentMethodChangeResponseFromJS(input.JSValue())
}

func (_this *PromiseNilPaymentMethodChangeResponse) Then(onFulfilled *PromiseNilPaymentMethodChangeResponseOnFulfilled, onRejected *PromiseNilPaymentMethodChangeResponseOnRejected) (_result *PromiseNilPaymentMethodChangeResponse) {
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
		_converted *PromiseNilPaymentMethodChangeResponse // javascript: Promise _what_return_name
	)
	_converted = PromiseNilPaymentMethodChangeResponseFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseNilPaymentMethodChangeResponse) Catch(onRejected *PromiseNilPaymentMethodChangeResponseOnRejected) (_result *PromiseNilPaymentMethodChangeResponse) {
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
		_converted *PromiseNilPaymentMethodChangeResponse // javascript: Promise _what_return_name
	)
	_converted = PromiseNilPaymentMethodChangeResponseFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseNilPaymentMethodChangeResponse) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseNilPaymentMethodChangeResponse) {
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
		_converted *PromiseNilPaymentMethodChangeResponse // javascript: Promise _what_return_name
	)
	_converted = PromiseNilPaymentMethodChangeResponseFromJS(_returned)
	_result = _converted
	return
}

// class: Promise
type PromisePaymentHandlerResponse struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *PromisePaymentHandlerResponse) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// PromisePaymentHandlerResponseFromJS is casting a js.Value into PromisePaymentHandlerResponse.
func PromisePaymentHandlerResponseFromJS(value js.Value) *PromisePaymentHandlerResponse {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromisePaymentHandlerResponse{}
	ret.Value_JS = value
	return ret
}

// PromisePaymentHandlerResponseFromJS is casting from something that holds a js.Value into PromisePaymentHandlerResponse.
func PromisePaymentHandlerResponseFromWrapper(input core.Wrapper) *PromisePaymentHandlerResponse {
	return PromisePaymentHandlerResponseFromJS(input.JSValue())
}

func (_this *PromisePaymentHandlerResponse) Then(onFulfilled *PromisePaymentHandlerResponseOnFulfilled, onRejected *PromisePaymentHandlerResponseOnRejected) (_result *PromisePaymentHandlerResponse) {
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
		_converted *PromisePaymentHandlerResponse // javascript: Promise _what_return_name
	)
	_converted = PromisePaymentHandlerResponseFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePaymentHandlerResponse) Catch(onRejected *PromisePaymentHandlerResponseOnRejected) (_result *PromisePaymentHandlerResponse) {
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
		_converted *PromisePaymentHandlerResponse // javascript: Promise _what_return_name
	)
	_converted = PromisePaymentHandlerResponseFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePaymentHandlerResponse) Finally(onFinally *javascript.PromiseFinally) (_result *PromisePaymentHandlerResponse) {
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
		_converted *PromisePaymentHandlerResponse // javascript: Promise _what_return_name
	)
	_converted = PromisePaymentHandlerResponseFromJS(_returned)
	_result = _converted
	return
}
