// Code generated by webidl-bind. DO NOT EDIT.

package presentation

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/file"
	"github.com/gowebapi/webapi/html/channel"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// channel.BinaryType
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// file.Blob
// javascript.ArrayBuffer
// javascript.FrozenArray
// javascript.Promise
// javascript.PromiseFinally

// source idl files:
// presentation-api.idl
// promises.idl

// transform files:
// presentation-api.go.md
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

// enum: PresentationConnectionCloseReason
type ConnectionCloseReason int

const (
	ReasonError ConnectionCloseReason = iota
	ReasonClosed
	ReasonWentaway
)

var presentationConnectionCloseReasonToWasmTable = []string{
	"error", "closed", "wentaway",
}

var presentationConnectionCloseReasonFromWasmTable = map[string]ConnectionCloseReason{
	"error": ReasonError, "closed": ReasonClosed, "wentaway": ReasonWentaway,
}

// JSValue is converting this enum into a javascript object
func (this *ConnectionCloseReason) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ConnectionCloseReason) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(presentationConnectionCloseReasonToWasmTable) {
		return presentationConnectionCloseReasonToWasmTable[idx]
	}
	panic("unknown input value")
}

// ConnectionCloseReasonFromJS is converting a javascript value into
// a ConnectionCloseReason enum value.
func ConnectionCloseReasonFromJS(value js.Value) ConnectionCloseReason {
	key := value.String()
	conv, ok := presentationConnectionCloseReasonFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: PresentationConnectionState
type ConnectionState int

const (
	StateConnecting ConnectionState = iota
	StateConnected
	StateClosed
	StateTerminated
)

var presentationConnectionStateToWasmTable = []string{
	"connecting", "connected", "closed", "terminated",
}

var presentationConnectionStateFromWasmTable = map[string]ConnectionState{
	"connecting": StateConnecting, "connected": StateConnected, "closed": StateClosed, "terminated": StateTerminated,
}

// JSValue is converting this enum into a javascript object
func (this *ConnectionState) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ConnectionState) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(presentationConnectionStateToWasmTable) {
		return presentationConnectionStateToWasmTable[idx]
	}
	panic("unknown input value")
}

// ConnectionStateFromJS is converting a javascript value into
// a ConnectionState enum value.
func ConnectionStateFromJS(value js.Value) ConnectionState {
	key := value.String()
	conv, ok := presentationConnectionStateFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: PromiseTemplateOnFulfilled
type PromiseAvailabilityOnFulfilledFunc func(value *Availability)

// PromiseAvailabilityOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseAvailabilityOnFulfilled js.Func

func PromiseAvailabilityOnFulfilledToJS(callback PromiseAvailabilityOnFulfilledFunc) *PromiseAvailabilityOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseAvailabilityOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Availability // javascript: PresentationAvailability value
		)
		_p0 = AvailabilityFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseAvailabilityOnFulfilledFromJS(_value js.Value) PromiseAvailabilityOnFulfilledFunc {
	return func(value *Availability) {
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
type PromiseAvailabilityOnRejectedFunc func(reason js.Value)

// PromiseAvailabilityOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseAvailabilityOnRejected js.Func

func PromiseAvailabilityOnRejectedToJS(callback PromiseAvailabilityOnRejectedFunc) *PromiseAvailabilityOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseAvailabilityOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromiseAvailabilityOnRejectedFromJS(_value js.Value) PromiseAvailabilityOnRejectedFunc {
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
type PromiseConnectionOnFulfilledFunc func(value *Connection)

// PromiseConnectionOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseConnectionOnFulfilled js.Func

func PromiseConnectionOnFulfilledToJS(callback PromiseConnectionOnFulfilledFunc) *PromiseConnectionOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseConnectionOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Connection // javascript: PresentationConnection value
		)
		_p0 = ConnectionFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseConnectionOnFulfilledFromJS(_value js.Value) PromiseConnectionOnFulfilledFunc {
	return func(value *Connection) {
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
type PromiseConnectionOnRejectedFunc func(reason js.Value)

// PromiseConnectionOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseConnectionOnRejected js.Func

func PromiseConnectionOnRejectedToJS(callback PromiseConnectionOnRejectedFunc) *PromiseConnectionOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseConnectionOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

func PromiseConnectionOnRejectedFromJS(_value js.Value) PromiseConnectionOnRejectedFunc {
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

// dictionary: PresentationConnectionAvailableEventInit
type ConnectionAvailableEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Connection *Connection
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ConnectionAvailableEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Connection.JSValue()
	out.Set("connection", value3)
	return out
}

// ConnectionAvailableEventInitFromJS is allocating a new
// ConnectionAvailableEventInit object and copy all values from
// input javascript object
func ConnectionAvailableEventInitFromJS(value js.Wrapper) *ConnectionAvailableEventInit {
	input := value.JSValue()
	var out ConnectionAvailableEventInit
	var (
		value0 bool        // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool        // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool        // javascript: boolean {composed Composed composed}
		value3 *Connection // javascript: PresentationConnection {connection Connection connection}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = ConnectionFromJS(input.Get("connection"))
	out.Connection = value3
	return &out
}

// dictionary: PresentationConnectionCloseEventInit
type ConnectionCloseEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Reason     ConnectionCloseReason
	Message    string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ConnectionCloseEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Reason.JSValue()
	out.Set("reason", value3)
	value4 := _this.Message
	out.Set("message", value4)
	return out
}

// ConnectionCloseEventInitFromJS is allocating a new
// ConnectionCloseEventInit object and copy all values from
// input javascript object
func ConnectionCloseEventInitFromJS(value js.Wrapper) *ConnectionCloseEventInit {
	input := value.JSValue()
	var out ConnectionCloseEventInit
	var (
		value0 bool                  // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool                  // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool                  // javascript: boolean {composed Composed composed}
		value3 ConnectionCloseReason // javascript: PresentationConnectionCloseReason {reason Reason reason}
		value4 string                // javascript: DOMString {message Message message}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = ConnectionCloseReasonFromJS(input.Get("reason"))
	out.Reason = value3
	value4 = (input.Get("message")).String()
	out.Message = value4
	return &out
}

// class: PresentationAvailability
type Availability struct {
	domcore.EventTarget
}

// AvailabilityFromJS is casting a js.Wrapper into Availability.
func AvailabilityFromJS(value js.Wrapper) *Availability {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Availability{}
	ret.Value_JS = input
	return ret
}

// Value returning attribute 'value' with
// type bool (idl: boolean).
func (_this *Availability) Value() bool {
	var ret bool
	value := _this.Value_JS.Get("value")
	ret = (value).Bool()
	return ret
}

// Onchange returning attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Availability) Onchange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onchange")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnchange setting attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Availability) SetOnchange(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onchange", input)
}

// class: PresentationConnection
type Connection struct {
	domcore.EventTarget
}

// ConnectionFromJS is casting a js.Wrapper into Connection.
func ConnectionFromJS(value js.Wrapper) *Connection {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Connection{}
	ret.Value_JS = input
	return ret
}

// Id returning attribute 'id' with
// type string (idl: USVString).
func (_this *Connection) Id() string {
	var ret string
	value := _this.Value_JS.Get("id")
	ret = (value).String()
	return ret
}

// Url returning attribute 'url' with
// type string (idl: USVString).
func (_this *Connection) Url() string {
	var ret string
	value := _this.Value_JS.Get("url")
	ret = (value).String()
	return ret
}

// State returning attribute 'state' with
// type ConnectionState (idl: PresentationConnectionState).
func (_this *Connection) State() ConnectionState {
	var ret ConnectionState
	value := _this.Value_JS.Get("state")
	ret = ConnectionStateFromJS(value)
	return ret
}

// Onconnect returning attribute 'onconnect' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Connection) Onconnect() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onconnect")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnconnect setting attribute 'onconnect' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Connection) SetOnconnect(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onconnect", input)
}

// Onclose returning attribute 'onclose' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Connection) Onclose() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onclose")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnclose setting attribute 'onclose' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Connection) SetOnclose(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onclose", input)
}

// Onterminate returning attribute 'onterminate' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Connection) Onterminate() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onterminate")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnterminate setting attribute 'onterminate' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Connection) SetOnterminate(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onterminate", input)
}

// BinaryType returning attribute 'binaryType' with
// type channel.BinaryType (idl: BinaryType).
func (_this *Connection) BinaryType() channel.BinaryType {
	var ret channel.BinaryType
	value := _this.Value_JS.Get("binaryType")
	ret = channel.BinaryTypeFromJS(value)
	return ret
}

// SetBinaryType setting attribute 'binaryType' with
// type channel.BinaryType (idl: BinaryType).
func (_this *Connection) SetBinaryType(value channel.BinaryType) {
	input := value.JSValue()
	_this.Value_JS.Set("binaryType", input)
}

// Onmessage returning attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Connection) Onmessage() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onmessage")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnmessage setting attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Connection) SetOnmessage(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onmessage", input)
}

func (_this *Connection) Close() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("close", _args[0:_end]...)
	return
}

func (_this *Connection) Terminate() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("terminate", _args[0:_end]...)
	return
}

func (_this *Connection) Send(message string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := message
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("send", _args[0:_end]...)
	return
}

func (_this *Connection) Send2(data *file.Blob) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("send", _args[0:_end]...)
	return
}

func (_this *Connection) Send3(data *javascript.ArrayBuffer) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("send", _args[0:_end]...)
	return
}

func (_this *Connection) Send4(data *Union) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("send", _args[0:_end]...)
	return
}

// class: PresentationConnectionAvailableEvent
type ConnectionAvailableEvent struct {
	domcore.Event
}

// ConnectionAvailableEventFromJS is casting a js.Wrapper into ConnectionAvailableEvent.
func ConnectionAvailableEventFromJS(value js.Wrapper) *ConnectionAvailableEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ConnectionAvailableEvent{}
	ret.Value_JS = input
	return ret
}

func NewPresentationConnectionAvailableEvent(_type string, eventInitDict *ConnectionAvailableEventInit) (_result *ConnectionAvailableEvent) {
	_klass := js.Global().Get("PresentationConnectionAvailableEvent")
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
		_converted *ConnectionAvailableEvent // javascript: PresentationConnectionAvailableEvent _what_return_name
	)
	_converted = ConnectionAvailableEventFromJS(_returned)
	_result = _converted
	return
}

// Connection returning attribute 'connection' with
// type Connection (idl: PresentationConnection).
func (_this *ConnectionAvailableEvent) Connection() *Connection {
	var ret *Connection
	value := _this.Value_JS.Get("connection")
	ret = ConnectionFromJS(value)
	return ret
}

// class: PresentationConnectionCloseEvent
type ConnectionCloseEvent struct {
	domcore.Event
}

// ConnectionCloseEventFromJS is casting a js.Wrapper into ConnectionCloseEvent.
func ConnectionCloseEventFromJS(value js.Wrapper) *ConnectionCloseEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ConnectionCloseEvent{}
	ret.Value_JS = input
	return ret
}

func NewPresentationConnectionCloseEvent(_type string, eventInitDict *ConnectionCloseEventInit) (_result *ConnectionCloseEvent) {
	_klass := js.Global().Get("PresentationConnectionCloseEvent")
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
		_converted *ConnectionCloseEvent // javascript: PresentationConnectionCloseEvent _what_return_name
	)
	_converted = ConnectionCloseEventFromJS(_returned)
	_result = _converted
	return
}

// Reason returning attribute 'reason' with
// type ConnectionCloseReason (idl: PresentationConnectionCloseReason).
func (_this *ConnectionCloseEvent) Reason() ConnectionCloseReason {
	var ret ConnectionCloseReason
	value := _this.Value_JS.Get("reason")
	ret = ConnectionCloseReasonFromJS(value)
	return ret
}

// Message returning attribute 'message' with
// type string (idl: DOMString).
func (_this *ConnectionCloseEvent) Message() string {
	var ret string
	value := _this.Value_JS.Get("message")
	ret = (value).String()
	return ret
}

// class: PresentationConnectionList
type ConnectionList struct {
	domcore.EventTarget
}

// ConnectionListFromJS is casting a js.Wrapper into ConnectionList.
func ConnectionListFromJS(value js.Wrapper) *ConnectionList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ConnectionList{}
	ret.Value_JS = input
	return ret
}

// Connections returning attribute 'connections' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *ConnectionList) Connections() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("connections")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// Onconnectionavailable returning attribute 'onconnectionavailable' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *ConnectionList) Onconnectionavailable() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onconnectionavailable")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnconnectionavailable setting attribute 'onconnectionavailable' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *ConnectionList) SetOnconnectionavailable(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onconnectionavailable", input)
}

// class: Presentation
type Presentation struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Presentation) JSValue() js.Value {
	return _this.Value_JS
}

// PresentationFromJS is casting a js.Wrapper into Presentation.
func PresentationFromJS(value js.Wrapper) *Presentation {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Presentation{}
	ret.Value_JS = input
	return ret
}

// DefaultRequest returning attribute 'defaultRequest' with
// type Request (idl: PresentationRequest).
func (_this *Presentation) DefaultRequest() *Request {
	var ret *Request
	value := _this.Value_JS.Get("defaultRequest")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = RequestFromJS(value)
	}
	return ret
}

// SetDefaultRequest setting attribute 'defaultRequest' with
// type Request (idl: PresentationRequest).
func (_this *Presentation) SetDefaultRequest(value *Request) {
	input := value.JSValue()
	_this.Value_JS.Set("defaultRequest", input)
}

// Receiver returning attribute 'receiver' with
// type Receiver (idl: PresentationReceiver).
func (_this *Presentation) Receiver() *Receiver {
	var ret *Receiver
	value := _this.Value_JS.Get("receiver")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = ReceiverFromJS(value)
	}
	return ret
}

// class: Promise
type PromiseAvailability struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseAvailability) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseAvailabilityFromJS is casting a js.Wrapper into PromiseAvailability.
func PromiseAvailabilityFromJS(value js.Wrapper) *PromiseAvailability {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromiseAvailability{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseAvailability) Then(onFulfilled *PromiseAvailabilityOnFulfilled, onRejected *PromiseAvailabilityOnRejected) (_result *PromiseAvailability) {
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
		_converted *PromiseAvailability // javascript: Promise _what_return_name
	)
	_converted = PromiseAvailabilityFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseAvailability) Catch(onRejected *PromiseAvailabilityOnRejected) (_result *PromiseAvailability) {
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
		_converted *PromiseAvailability // javascript: Promise _what_return_name
	)
	_converted = PromiseAvailabilityFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseAvailability) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseAvailability) {
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
		_converted *PromiseAvailability // javascript: Promise _what_return_name
	)
	_converted = PromiseAvailabilityFromJS(_returned)
	_result = _converted
	return
}

// class: Promise
type PromiseConnection struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseConnection) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseConnectionFromJS is casting a js.Wrapper into PromiseConnection.
func PromiseConnectionFromJS(value js.Wrapper) *PromiseConnection {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromiseConnection{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseConnection) Then(onFulfilled *PromiseConnectionOnFulfilled, onRejected *PromiseConnectionOnRejected) (_result *PromiseConnection) {
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
		_converted *PromiseConnection // javascript: Promise _what_return_name
	)
	_converted = PromiseConnectionFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseConnection) Catch(onRejected *PromiseConnectionOnRejected) (_result *PromiseConnection) {
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
		_converted *PromiseConnection // javascript: Promise _what_return_name
	)
	_converted = PromiseConnectionFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseConnection) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseConnection) {
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
		_converted *PromiseConnection // javascript: Promise _what_return_name
	)
	_converted = PromiseConnectionFromJS(_returned)
	_result = _converted
	return
}

// class: PresentationReceiver
type Receiver struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Receiver) JSValue() js.Value {
	return _this.Value_JS
}

// ReceiverFromJS is casting a js.Wrapper into Receiver.
func ReceiverFromJS(value js.Wrapper) *Receiver {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Receiver{}
	ret.Value_JS = input
	return ret
}

// ConnectionList returning attribute 'connectionList' with
// type javascript.Promise (idl: Promise).
func (_this *Receiver) ConnectionList() *javascript.Promise {
	var ret *javascript.Promise
	value := _this.Value_JS.Get("connectionList")
	ret = javascript.PromiseFromJS(value)
	return ret
}

// class: PresentationRequest
type Request struct {
	domcore.EventTarget
}

// RequestFromJS is casting a js.Wrapper into Request.
func RequestFromJS(value js.Wrapper) *Request {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Request{}
	ret.Value_JS = input
	return ret
}

func NewPresentationRequest(urls []string) (_result *Request) {
	_klass := js.Global().Get("PresentationRequest")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := js.Global().Get("Array").New(len(urls))
	for __idx0, __seq_in0 := range urls {
		__seq_out0 := __seq_in0
		_p0.SetIndex(__idx0, __seq_out0)
	}
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Request // javascript: PresentationRequest _what_return_name
	)
	_converted = RequestFromJS(_returned)
	_result = _converted
	return
}

// Onconnectionavailable returning attribute 'onconnectionavailable' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Request) Onconnectionavailable() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onconnectionavailable")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnconnectionavailable setting attribute 'onconnectionavailable' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Request) SetOnconnectionavailable(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onconnectionavailable", input)
}

func (_this *Request) Start() (_result *PromiseConnection) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("start", _args[0:_end]...)
	var (
		_converted *PromiseConnection // javascript: Promise _what_return_name
	)
	_converted = PromiseConnectionFromJS(_returned)
	_result = _converted
	return
}

func (_this *Request) Reconnect(presentationId string) (_result *PromiseConnection) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := presentationId
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("reconnect", _args[0:_end]...)
	var (
		_converted *PromiseConnection // javascript: Promise _what_return_name
	)
	_converted = PromiseConnectionFromJS(_returned)
	_result = _converted
	return
}

func (_this *Request) GetAvailability() (_result *PromiseAvailability) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getAvailability", _args[0:_end]...)
	var (
		_converted *PromiseAvailability // javascript: Promise _what_return_name
	)
	_converted = PromiseAvailabilityFromJS(_returned)
	_result = _converted
	return
}
