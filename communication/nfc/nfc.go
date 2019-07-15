// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package nfc

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.AbortSignal
// domcore.DOMException
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// javascript.PromiseVoid

// source idl files:
// web-nfc.idl

// transform files:
// web-nfc.go.md

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

// enum: NDEFCompatibility
type NDEFCompatibility int

const (
	CompatibilityNfcForum NDEFCompatibility = iota
	CompatibilityVendor
	CompatibilityAny
)

var nDEFCompatibilityToWasmTable = []string{
	"nfc-forum", "vendor", "any",
}

var nDEFCompatibilityFromWasmTable = map[string]NDEFCompatibility{
	"nfc-forum": CompatibilityNfcForum, "vendor": CompatibilityVendor, "any": CompatibilityAny,
}

// JSValue is converting this enum into a java object
func (this *NDEFCompatibility) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this NDEFCompatibility) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(nDEFCompatibilityToWasmTable) {
		return nDEFCompatibilityToWasmTable[idx]
	}
	panic("unknown input value")
}

// NDEFCompatibilityFromJS is converting a javascript value into
// a NDEFCompatibility enum value.
func NDEFCompatibilityFromJS(value js.Value) NDEFCompatibility {
	key := value.String()
	conv, ok := nDEFCompatibilityFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: NDEFRecordType
type NDEFRecordType int

const (
	RecordEmpty NDEFRecordType = iota
	RecordText
	RecordUrl
	RecordJson
	RecordOpaque
)

var nDEFRecordTypeToWasmTable = []string{
	"empty", "text", "url", "json", "opaque",
}

var nDEFRecordTypeFromWasmTable = map[string]NDEFRecordType{
	"empty": RecordEmpty, "text": RecordText, "url": RecordUrl, "json": RecordJson, "opaque": RecordOpaque,
}

// JSValue is converting this enum into a java object
func (this *NDEFRecordType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this NDEFRecordType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(nDEFRecordTypeToWasmTable) {
		return nDEFRecordTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// NDEFRecordTypeFromJS is converting a javascript value into
// a NDEFRecordType enum value.
func NDEFRecordTypeFromJS(value js.Value) NDEFRecordType {
	key := value.String()
	conv, ok := nDEFRecordTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: NFCPushTarget
type PushTarget int

const (
	PushTargetTag PushTarget = iota
	PushTargetPeer
	PushTargetAny
)

var nFCPushTargetToWasmTable = []string{
	"tag", "peer", "any",
}

var nFCPushTargetFromWasmTable = map[string]PushTarget{
	"tag": PushTargetTag, "peer": PushTargetPeer, "any": PushTargetAny,
}

// JSValue is converting this enum into a java object
func (this *PushTarget) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this PushTarget) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(nFCPushTargetToWasmTable) {
		return nFCPushTargetToWasmTable[idx]
	}
	panic("unknown input value")
}

// PushTargetFromJS is converting a javascript value into
// a PushTarget enum value.
func PushTargetFromJS(value js.Value) PushTarget {
	key := value.String()
	conv, ok := nFCPushTargetFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: NFCErrorEventInit
type ErrorEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Error      *domcore.DOMException
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ErrorEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Error.JSValue()
	out.Set("error", value3)
	return out
}

// ErrorEventInitFromJS is allocating a new
// ErrorEventInit object and copy all values from
// input javascript object
func ErrorEventInitFromJS(value js.Wrapper) *ErrorEventInit {
	input := value.JSValue()
	var out ErrorEventInit
	var (
		value0 bool                  // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool                  // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool                  // javascript: boolean {composed Composed composed}
		value3 *domcore.DOMException // javascript: DOMException {error Error _error}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = domcore.DOMExceptionFromJS(input.Get("error"))
	out.Error = value3
	return &out
}

// dictionary: NDEFMessage
type NDEFMessage struct {
	Records []*NDEFRecord
	Url     string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *NDEFMessage) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Records))
	for __idx0, __seq_in0 := range _this.Records {
		__seq_out0 := __seq_in0.JSValue()
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("records", value0)
	value1 := _this.Url
	out.Set("url", value1)
	return out
}

// NDEFMessageFromJS is allocating a new
// NDEFMessage object and copy all values from
// input javascript object
func NDEFMessageFromJS(value js.Wrapper) *NDEFMessage {
	input := value.JSValue()
	var out NDEFMessage
	var (
		value0 []*NDEFRecord // javascript: sequence<NDEFRecord> {records Records records}
		value1 string        // javascript: USVString {url Url url}
	)
	__length0 := input.Get("records").Length()
	__array0 := make([]*NDEFRecord, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *NDEFRecord
		__seq_in0 := input.Get("records").Index(__idx0)
		__seq_out0 = NDEFRecordFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Records = value0
	value1 = (input.Get("url")).String()
	out.Url = value1
	return &out
}

// dictionary: NDEFRecord
type NDEFRecord struct {
	RecordType NDEFRecordType
	MediaType  string
	Data       *Union
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *NDEFRecord) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.RecordType.JSValue()
	out.Set("recordType", value0)
	value1 := _this.MediaType
	out.Set("mediaType", value1)
	value2 := _this.Data.JSValue()
	out.Set("data", value2)
	return out
}

// NDEFRecordFromJS is allocating a new
// NDEFRecord object and copy all values from
// input javascript object
func NDEFRecordFromJS(value js.Wrapper) *NDEFRecord {
	input := value.JSValue()
	var out NDEFRecord
	var (
		value0 NDEFRecordType // javascript: NDEFRecordType {recordType RecordType recordType}
		value1 string         // javascript: USVString {mediaType MediaType mediaType}
		value2 *Union         // javascript: Union {data Data data}
	)
	value0 = NDEFRecordTypeFromJS(input.Get("recordType"))
	out.RecordType = value0
	value1 = (input.Get("mediaType")).String()
	out.MediaType = value1
	value2 = UnionFromJS(input.Get("data"))
	out.Data = value2
	return &out
}

// dictionary: NFCPushOptions
type PushOptions struct {
	Target        PushTarget
	Timeout       float64
	IgnoreRead    bool
	Signal        *domcore.AbortSignal
	Compatibility NDEFCompatibility
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PushOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Target.JSValue()
	out.Set("target", value0)
	value1 := _this.Timeout
	out.Set("timeout", value1)
	value2 := _this.IgnoreRead
	out.Set("ignoreRead", value2)
	value3 := _this.Signal.JSValue()
	out.Set("signal", value3)
	value4 := _this.Compatibility.JSValue()
	out.Set("compatibility", value4)
	return out
}

// PushOptionsFromJS is allocating a new
// PushOptions object and copy all values from
// input javascript object
func PushOptionsFromJS(value js.Wrapper) *PushOptions {
	input := value.JSValue()
	var out PushOptions
	var (
		value0 PushTarget           // javascript: NFCPushTarget {target Target target}
		value1 float64              // javascript: unrestricted double {timeout Timeout timeout}
		value2 bool                 // javascript: boolean {ignoreRead IgnoreRead ignoreRead}
		value3 *domcore.AbortSignal // javascript: AbortSignal {signal Signal signal}
		value4 NDEFCompatibility    // javascript: NDEFCompatibility {compatibility Compatibility compatibility}
	)
	value0 = PushTargetFromJS(input.Get("target"))
	out.Target = value0
	value1 = (input.Get("timeout")).Float()
	out.Timeout = value1
	value2 = (input.Get("ignoreRead")).Bool()
	out.IgnoreRead = value2
	if input.Get("signal").Type() != js.TypeNull && input.Get("signal").Type() != js.TypeUndefined {
		value3 = domcore.AbortSignalFromJS(input.Get("signal"))
	}
	out.Signal = value3
	value4 = NDEFCompatibilityFromJS(input.Get("compatibility"))
	out.Compatibility = value4
	return &out
}

// dictionary: NFCReaderOptions
type ReaderOptions struct {
	Url           string
	RecordType    NDEFRecordType
	MediaType     string
	Compatibility NDEFCompatibility
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ReaderOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Url
	out.Set("url", value0)
	value1 := _this.RecordType.JSValue()
	out.Set("recordType", value1)
	value2 := _this.MediaType
	out.Set("mediaType", value2)
	value3 := _this.Compatibility.JSValue()
	out.Set("compatibility", value3)
	return out
}

// ReaderOptionsFromJS is allocating a new
// ReaderOptions object and copy all values from
// input javascript object
func ReaderOptionsFromJS(value js.Wrapper) *ReaderOptions {
	input := value.JSValue()
	var out ReaderOptions
	var (
		value0 string            // javascript: USVString {url Url url}
		value1 NDEFRecordType    // javascript: NDEFRecordType {recordType RecordType recordType}
		value2 string            // javascript: USVString {mediaType MediaType mediaType}
		value3 NDEFCompatibility // javascript: NDEFCompatibility {compatibility Compatibility compatibility}
	)
	value0 = (input.Get("url")).String()
	out.Url = value0
	value1 = NDEFRecordTypeFromJS(input.Get("recordType"))
	out.RecordType = value1
	value2 = (input.Get("mediaType")).String()
	out.MediaType = value2
	value3 = NDEFCompatibilityFromJS(input.Get("compatibility"))
	out.Compatibility = value3
	return &out
}

// dictionary: NFCReadingEventInit
type ReadingEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	Message    *NDEFMessage
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ReadingEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Message.JSValue()
	out.Set("message", value3)
	return out
}

// ReadingEventInitFromJS is allocating a new
// ReadingEventInit object and copy all values from
// input javascript object
func ReadingEventInitFromJS(value js.Wrapper) *ReadingEventInit {
	input := value.JSValue()
	var out ReadingEventInit
	var (
		value0 bool         // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool         // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool         // javascript: boolean {composed Composed composed}
		value3 *NDEFMessage // javascript: NDEFMessage {message Message message}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = NDEFMessageFromJS(input.Get("message"))
	out.Message = value3
	return &out
}

// interface: NFCErrorEvent
type ErrorEvent struct {
	domcore.Event
}

// ErrorEventFromJS is casting a js.Wrapper into ErrorEvent.
func ErrorEventFromJS(value js.Wrapper) *ErrorEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ErrorEvent{}
	ret.Value_JS = input
	return ret
}

func NewNFCErrorEvent(_type string, errorEventInitDict *ErrorEventInit) (_result *ErrorEvent) {
	_klass := js.Global().Get("NFCErrorEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := errorEventInitDict.JSValue()
	_args[1] = _p1
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ErrorEvent // javascript: NFCErrorEvent _what_return_name
	)
	_converted = ErrorEventFromJS(_returned)
	_result = _converted
	return
}

// Error returning attribute 'error' with
// type domcore.DOMException (idl: DOMException).
func (_this *ErrorEvent) Error() *domcore.DOMException {
	var ret *domcore.DOMException
	value := _this.Value_JS.Get("error")
	ret = domcore.DOMExceptionFromJS(value)
	return ret
}

// interface: NFCReader
type Reader struct {
	domcore.EventTarget
}

// ReaderFromJS is casting a js.Wrapper into Reader.
func ReaderFromJS(value js.Wrapper) *Reader {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Reader{}
	ret.Value_JS = input
	return ret
}

func NewNFCReader(options *ReaderOptions) (_result *Reader) {
	_klass := js.Global().Get("NFCReader")
	var (
		_args [1]interface{}
		_end  int
	)
	if options != nil {
		_p0 := options.JSValue()
		_args[0] = _p0
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Reader // javascript: NFCReader _what_return_name
	)
	_converted = ReaderFromJS(_returned)
	_result = _converted
	return
}

// Onreading returning attribute 'onreading' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Reader) Onreading() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onreading")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnreading setting attribute 'onreading' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Reader) SetOnreading(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onreading", input)
}

// Onerror returning attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Reader) Onerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnerror setting attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Reader) SetOnerror(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onerror", input)
}

func (_this *Reader) Start() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("start", _args[0:_end]...)
	return
}

func (_this *Reader) Stop() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("stop", _args[0:_end]...)
	return
}

// interface: NFCReadingEvent
type ReadingEvent struct {
	domcore.Event
}

// ReadingEventFromJS is casting a js.Wrapper into ReadingEvent.
func ReadingEventFromJS(value js.Wrapper) *ReadingEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ReadingEvent{}
	ret.Value_JS = input
	return ret
}

func NewNFCReadingEvent(_type string, readingEventInitDict *ReadingEventInit) (_result *ReadingEvent) {
	_klass := js.Global().Get("NFCReadingEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := readingEventInitDict.JSValue()
	_args[1] = _p1
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ReadingEvent // javascript: NFCReadingEvent _what_return_name
	)
	_converted = ReadingEventFromJS(_returned)
	_result = _converted
	return
}

// Message returning attribute 'message' with
// type NDEFMessage (idl: NDEFMessage).
func (_this *ReadingEvent) Message() *NDEFMessage {
	var ret *NDEFMessage
	value := _this.Value_JS.Get("message")
	ret = NDEFMessageFromJS(value)
	return ret
}

// interface: NFCWriter
type Writer struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Writer) JSValue() js.Value {
	return _this.Value_JS
}

// WriterFromJS is casting a js.Wrapper into Writer.
func WriterFromJS(value js.Wrapper) *Writer {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Writer{}
	ret.Value_JS = input
	return ret
}

func NewNFCWriter() (_result *Writer) {
	_klass := js.Global().Get("NFCWriter")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Writer // javascript: NFCWriter _what_return_name
	)
	_converted = WriterFromJS(_returned)
	_result = _converted
	return
}

func (_this *Writer) Push(message *Union, options *PushOptions) (_result *javascript.PromiseVoid) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := message.JSValue()
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("push", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}
