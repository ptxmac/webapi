// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package fetch

import js "github.com/gowebapi/webapi/core/failjs"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
	"github.com/gowebapi/webapi/patch"
)

// using following types:
// domcore.AbortSignal
// javascript.Promise
// patch.ByteString
// patch.ReadableStream

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

// enum: ReferrerPolicy
type ReferrerPolicy int

const (
	EmptyString0ReferrerPolicy ReferrerPolicy = iota
	NoReferrerReferrerPolicy
	NoReferrerWhenDowngradeReferrerPolicy
	SameOriginReferrerPolicy
	OriginReferrerPolicy
	StrictOriginReferrerPolicy
	OriginWhenCrossOriginReferrerPolicy
	StrictOriginWhenCrossOriginReferrerPolicy
	UnsafeUrlReferrerPolicy
)

var referrerPolicyToWasmTable = []string{
	"", "no-referrer", "no-referrer-when-downgrade", "same-origin", "origin", "strict-origin", "origin-when-cross-origin", "strict-origin-when-cross-origin", "unsafe-url",
}

var referrerPolicyFromWasmTable = map[string]ReferrerPolicy{
	"": EmptyString0ReferrerPolicy, "no-referrer": NoReferrerReferrerPolicy, "no-referrer-when-downgrade": NoReferrerWhenDowngradeReferrerPolicy, "same-origin": SameOriginReferrerPolicy, "origin": OriginReferrerPolicy, "strict-origin": StrictOriginReferrerPolicy, "origin-when-cross-origin": OriginWhenCrossOriginReferrerPolicy, "strict-origin-when-cross-origin": StrictOriginWhenCrossOriginReferrerPolicy, "unsafe-url": UnsafeUrlReferrerPolicy,
}

// JSValue is converting this enum into a java object
func (this *ReferrerPolicy) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ReferrerPolicy) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(referrerPolicyToWasmTable) {
		return referrerPolicyToWasmTable[idx]
	}
	panic("unknown input value")
}

// ReferrerPolicyFromJS is converting a javascript value into
// a ReferrerPolicy enum value.
func ReferrerPolicyFromJS(value js.Value) ReferrerPolicy {
	key := value.String()
	conv, ok := referrerPolicyFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: RequestCache
type RequestCache int

const (
	DefaultRequestCache RequestCache = iota
	NoStoreRequestCache
	ReloadRequestCache
	NoCacheRequestCache
	ForceCacheRequestCache
	OnlyIfCachedRequestCache
)

var requestCacheToWasmTable = []string{
	"default", "no-store", "reload", "no-cache", "force-cache", "only-if-cached",
}

var requestCacheFromWasmTable = map[string]RequestCache{
	"default": DefaultRequestCache, "no-store": NoStoreRequestCache, "reload": ReloadRequestCache, "no-cache": NoCacheRequestCache, "force-cache": ForceCacheRequestCache, "only-if-cached": OnlyIfCachedRequestCache,
}

// JSValue is converting this enum into a java object
func (this *RequestCache) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this RequestCache) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(requestCacheToWasmTable) {
		return requestCacheToWasmTable[idx]
	}
	panic("unknown input value")
}

// RequestCacheFromJS is converting a javascript value into
// a RequestCache enum value.
func RequestCacheFromJS(value js.Value) RequestCache {
	key := value.String()
	conv, ok := requestCacheFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: RequestCredentials
type RequestCredentials int

const (
	OmitRequestCredentials RequestCredentials = iota
	SameOriginRequestCredentials
	IncludeRequestCredentials
)

var requestCredentialsToWasmTable = []string{
	"omit", "same-origin", "include",
}

var requestCredentialsFromWasmTable = map[string]RequestCredentials{
	"omit": OmitRequestCredentials, "same-origin": SameOriginRequestCredentials, "include": IncludeRequestCredentials,
}

// JSValue is converting this enum into a java object
func (this *RequestCredentials) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this RequestCredentials) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(requestCredentialsToWasmTable) {
		return requestCredentialsToWasmTable[idx]
	}
	panic("unknown input value")
}

// RequestCredentialsFromJS is converting a javascript value into
// a RequestCredentials enum value.
func RequestCredentialsFromJS(value js.Value) RequestCredentials {
	key := value.String()
	conv, ok := requestCredentialsFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: RequestDestination
type RequestDestination int

const (
	EmptyString0RequestDestination RequestDestination = iota
	AudioRequestDestination
	AudioworkletRequestDestination
	DocumentRequestDestination
	EmbedRequestDestination
	FontRequestDestination
	ImageRequestDestination
	ManifestRequestDestination
	ObjectRequestDestination
	PaintworkletRequestDestination
	ReportRequestDestination
	ScriptRequestDestination
	SharedworkerRequestDestination
	StyleRequestDestination
	TrackRequestDestination
	VideoRequestDestination
	WorkerRequestDestination
	XsltRequestDestination
)

var requestDestinationToWasmTable = []string{
	"", "audio", "audioworklet", "document", "embed", "font", "image", "manifest", "object", "paintworklet", "report", "script", "sharedworker", "style", "track", "video", "worker", "xslt",
}

var requestDestinationFromWasmTable = map[string]RequestDestination{
	"": EmptyString0RequestDestination, "audio": AudioRequestDestination, "audioworklet": AudioworkletRequestDestination, "document": DocumentRequestDestination, "embed": EmbedRequestDestination, "font": FontRequestDestination, "image": ImageRequestDestination, "manifest": ManifestRequestDestination, "object": ObjectRequestDestination, "paintworklet": PaintworkletRequestDestination, "report": ReportRequestDestination, "script": ScriptRequestDestination, "sharedworker": SharedworkerRequestDestination, "style": StyleRequestDestination, "track": TrackRequestDestination, "video": VideoRequestDestination, "worker": WorkerRequestDestination, "xslt": XsltRequestDestination,
}

// JSValue is converting this enum into a java object
func (this *RequestDestination) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this RequestDestination) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(requestDestinationToWasmTable) {
		return requestDestinationToWasmTable[idx]
	}
	panic("unknown input value")
}

// RequestDestinationFromJS is converting a javascript value into
// a RequestDestination enum value.
func RequestDestinationFromJS(value js.Value) RequestDestination {
	key := value.String()
	conv, ok := requestDestinationFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: RequestMode
type RequestMode int

const (
	NavigateRequestMode RequestMode = iota
	SameOriginRequestMode
	NoCorsRequestMode
	CorsRequestMode
)

var requestModeToWasmTable = []string{
	"navigate", "same-origin", "no-cors", "cors",
}

var requestModeFromWasmTable = map[string]RequestMode{
	"navigate": NavigateRequestMode, "same-origin": SameOriginRequestMode, "no-cors": NoCorsRequestMode, "cors": CorsRequestMode,
}

// JSValue is converting this enum into a java object
func (this *RequestMode) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this RequestMode) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(requestModeToWasmTable) {
		return requestModeToWasmTable[idx]
	}
	panic("unknown input value")
}

// RequestModeFromJS is converting a javascript value into
// a RequestMode enum value.
func RequestModeFromJS(value js.Value) RequestMode {
	key := value.String()
	conv, ok := requestModeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: RequestRedirect
type RequestRedirect int

const (
	FollowRequestRedirect RequestRedirect = iota
	ErrorRequestRedirect
	ManualRequestRedirect
)

var requestRedirectToWasmTable = []string{
	"follow", "error", "manual",
}

var requestRedirectFromWasmTable = map[string]RequestRedirect{
	"follow": FollowRequestRedirect, "error": ErrorRequestRedirect, "manual": ManualRequestRedirect,
}

// JSValue is converting this enum into a java object
func (this *RequestRedirect) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this RequestRedirect) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(requestRedirectToWasmTable) {
		return requestRedirectToWasmTable[idx]
	}
	panic("unknown input value")
}

// RequestRedirectFromJS is converting a javascript value into
// a RequestRedirect enum value.
func RequestRedirectFromJS(value js.Value) RequestRedirect {
	key := value.String()
	conv, ok := requestRedirectFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: ResponseType
type ResponseType int

const (
	BasicResponseType ResponseType = iota
	CorsResponseType
	DefaultResponseType
	ErrorResponseType
	OpaqueResponseType
	OpaqueredirectResponseType
)

var responseTypeToWasmTable = []string{
	"basic", "cors", "default", "error", "opaque", "opaqueredirect",
}

var responseTypeFromWasmTable = map[string]ResponseType{
	"basic": BasicResponseType, "cors": CorsResponseType, "default": DefaultResponseType, "error": ErrorResponseType, "opaque": OpaqueResponseType, "opaqueredirect": OpaqueredirectResponseType,
}

// JSValue is converting this enum into a java object
func (this *ResponseType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ResponseType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(responseTypeToWasmTable) {
		return responseTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// ResponseTypeFromJS is converting a javascript value into
// a ResponseType enum value.
func ResponseTypeFromJS(value js.Value) ResponseType {
	key := value.String()
	conv, ok := responseTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: RequestInit
type RequestInit struct {
	Method         *patch.ByteString
	Headers        [][]*patch.ByteString
	Body           *Union
	Referrer       string
	ReferrerPolicy ReferrerPolicy
	Mode           RequestMode
	Credentials    RequestCredentials
	Cache          RequestCache
	Redirect       RequestRedirect
	Integrity      string
	Keepalive      bool
	Signal         *domcore.AbortSignal
	Window         js.Value
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *RequestInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Method.JSValue()
	out.Set("method", value0)
	value1 := js.Global().Get("Array").New(len(_this.Headers))
	for __idx1, __seq_in1 := range _this.Headers {
		__seq_out1 := js.Global().Get("Array").New(len(__seq_in1))
		for __idx2, __seq_in2 := range __seq_in1 {
			__seq_out2 := __seq_in2.JSValue()
			__seq_out1.SetIndex(__idx2, __seq_out2)
		}
		value1.SetIndex(__idx1, __seq_out1)
	}
	out.Set("headers", value1)
	value2 := _this.Body.JSValue()
	out.Set("body", value2)
	value3 := _this.Referrer
	out.Set("referrer", value3)
	value4 := _this.ReferrerPolicy.JSValue()
	out.Set("referrerPolicy", value4)
	value5 := _this.Mode.JSValue()
	out.Set("mode", value5)
	value6 := _this.Credentials.JSValue()
	out.Set("credentials", value6)
	value7 := _this.Cache.JSValue()
	out.Set("cache", value7)
	value8 := _this.Redirect.JSValue()
	out.Set("redirect", value8)
	value9 := _this.Integrity
	out.Set("integrity", value9)
	value10 := _this.Keepalive
	out.Set("keepalive", value10)
	value11 := _this.Signal.JSValue()
	out.Set("signal", value11)
	value12 := _this.Window
	out.Set("window", value12)
	return out
}

// RequestInitFromJS is allocating a new
// RequestInit object and copy all values from
// input javascript object
func RequestInitFromJS(value js.Wrapper) *RequestInit {
	input := value.JSValue()
	var out RequestInit
	var (
		value0  *patch.ByteString     // javascript: ByteString {method Method method}
		value1  [][]*patch.ByteString // javascript: sequence<sequence<ByteString>> {headers Headers headers}
		value2  *Union                // javascript: Union {body Body body}
		value3  string                // javascript: USVString {referrer Referrer referrer}
		value4  ReferrerPolicy        // javascript: ReferrerPolicy {referrerPolicy ReferrerPolicy referrerPolicy}
		value5  RequestMode           // javascript: RequestMode {mode Mode mode}
		value6  RequestCredentials    // javascript: RequestCredentials {credentials Credentials credentials}
		value7  RequestCache          // javascript: RequestCache {cache Cache cache}
		value8  RequestRedirect       // javascript: RequestRedirect {redirect Redirect redirect}
		value9  string                // javascript: DOMString {integrity Integrity integrity}
		value10 bool                  // javascript: boolean {keepalive Keepalive keepalive}
		value11 *domcore.AbortSignal  // javascript: AbortSignal {signal Signal signal}
		value12 js.Value              // javascript: any {window Window window}
	)
	value0 = patch.ByteStringFromJS(input.Get("method"))
	out.Method = value0
	__length1 := input.Get("headers").Length()
	__array1 := make([][]*patch.ByteString, __length1, __length1)
	for __idx1 := 0; __idx1 < __length1; __idx1++ {
		var __seq_out1 []*patch.ByteString
		__seq_in1 := input.Get("headers").Index(__idx1)
		__length2 := __seq_in1.Length()
		__array2 := make([]*patch.ByteString, __length2, __length2)
		for __idx2 := 0; __idx2 < __length2; __idx2++ {
			var __seq_out2 *patch.ByteString
			__seq_in2 := __seq_in1.Index(__idx2)
			__seq_out2 = patch.ByteStringFromJS(__seq_in2)
			__array2[__idx2] = __seq_out2
		}
		__seq_out1 = __array2
		__array1[__idx1] = __seq_out1
	}
	value1 = __array1
	out.Headers = value1
	if input.Get("body").Type() != js.TypeNull {
		value2 = UnionFromJS(input.Get("body"))
	}
	out.Body = value2
	value3 = (input.Get("referrer")).String()
	out.Referrer = value3
	value4 = ReferrerPolicyFromJS(input.Get("referrerPolicy"))
	out.ReferrerPolicy = value4
	value5 = RequestModeFromJS(input.Get("mode"))
	out.Mode = value5
	value6 = RequestCredentialsFromJS(input.Get("credentials"))
	out.Credentials = value6
	value7 = RequestCacheFromJS(input.Get("cache"))
	out.Cache = value7
	value8 = RequestRedirectFromJS(input.Get("redirect"))
	out.Redirect = value8
	value9 = (input.Get("integrity")).String()
	out.Integrity = value9
	value10 = (input.Get("keepalive")).Bool()
	out.Keepalive = value10
	if input.Get("signal").Type() != js.TypeNull {
		value11 = domcore.AbortSignalFromJS(input.Get("signal"))
	}
	out.Signal = value11
	value12 = input.Get("window")
	out.Window = value12
	return &out
}

// dictionary: ResponseInit
type ResponseInit struct {
	Status     int
	StatusText *patch.ByteString
	Headers    [][]*patch.ByteString
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ResponseInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Status
	out.Set("status", value0)
	value1 := _this.StatusText.JSValue()
	out.Set("statusText", value1)
	value2 := js.Global().Get("Array").New(len(_this.Headers))
	for __idx2, __seq_in2 := range _this.Headers {
		__seq_out2 := js.Global().Get("Array").New(len(__seq_in2))
		for __idx3, __seq_in3 := range __seq_in2 {
			__seq_out3 := __seq_in3.JSValue()
			__seq_out2.SetIndex(__idx3, __seq_out3)
		}
		value2.SetIndex(__idx2, __seq_out2)
	}
	out.Set("headers", value2)
	return out
}

// ResponseInitFromJS is allocating a new
// ResponseInit object and copy all values from
// input javascript object
func ResponseInitFromJS(value js.Wrapper) *ResponseInit {
	input := value.JSValue()
	var out ResponseInit
	var (
		value0 int                   // javascript: unsigned short {status Status status}
		value1 *patch.ByteString     // javascript: ByteString {statusText StatusText statusText}
		value2 [][]*patch.ByteString // javascript: sequence<sequence<ByteString>> {headers Headers headers}
	)
	value0 = (input.Get("status")).Int()
	out.Status = value0
	value1 = patch.ByteStringFromJS(input.Get("statusText"))
	out.StatusText = value1
	__length2 := input.Get("headers").Length()
	__array2 := make([][]*patch.ByteString, __length2, __length2)
	for __idx2 := 0; __idx2 < __length2; __idx2++ {
		var __seq_out2 []*patch.ByteString
		__seq_in2 := input.Get("headers").Index(__idx2)
		__length3 := __seq_in2.Length()
		__array3 := make([]*patch.ByteString, __length3, __length3)
		for __idx3 := 0; __idx3 < __length3; __idx3++ {
			var __seq_out3 *patch.ByteString
			__seq_in3 := __seq_in2.Index(__idx3)
			__seq_out3 = patch.ByteStringFromJS(__seq_in3)
			__array3[__idx3] = __seq_out3
		}
		__seq_out2 = __array3
		__array2[__idx2] = __seq_out2
	}
	value2 = __array2
	out.Headers = value2
	return &out
}

// interface: Headers
type Headers struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Headers) JSValue() js.Value {
	return _this.Value_JS
}

// HeadersFromJS is casting a js.Wrapper into Headers.
func HeadersFromJS(value js.Wrapper) *Headers {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Headers{}
	ret.Value_JS = input
	return ret
}

func NewHeaders(init [][]*patch.ByteString) (_result *Headers) {
	_klass := js.Global().Get("Headers")
	var (
		_args [1]interface{}
		_end  int
	)
	if init != nil {
		_p0 := js.Global().Get("Array").New(len(init))
		for __idx0, __seq_in0 := range init {
			__seq_out0 := js.Global().Get("Array").New(len(__seq_in0))
			for __idx1, __seq_in1 := range __seq_in0 {
				__seq_out1 := __seq_in1.JSValue()
				__seq_out0.SetIndex(__idx1, __seq_out1)
			}
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Headers // javascript: Headers _what_return_name
	)
	_converted = HeadersFromJS(_returned)
	_result = _converted
	return
}

func (_this *Headers) Append(name *patch.ByteString, value *patch.ByteString) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name.JSValue()
	_args[0] = _p0
	_end++
	_p1 := value.JSValue()
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("append", _args[0:_end]...)
	return
}

func (_this *Headers) Delete(name *patch.ByteString) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("delete", _args[0:_end]...)
	return
}

func (_this *Headers) Get(name *patch.ByteString) (_result *patch.ByteString) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *patch.ByteString // javascript: ByteString _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = patch.ByteStringFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *Headers) Has(name *patch.ByteString) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name.JSValue()
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

func (_this *Headers) Set(name *patch.ByteString, value *patch.ByteString) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name.JSValue()
	_args[0] = _p0
	_end++
	_p1 := value.JSValue()
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("set", _args[0:_end]...)
	return
}

// interface: Request
type Request struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Request) JSValue() js.Value {
	return _this.Value_JS
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

func NewRequest(input *Union, init *RequestInit) (_result *Request) {
	_klass := js.Global().Get("Request")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := input.JSValue()
	_args[0] = _p0
	_end++
	if init != nil {
		_p1 := init.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Request // javascript: Request _what_return_name
	)
	_converted = RequestFromJS(_returned)
	_result = _converted
	return
}

// Method returning attribute 'method' with
// type patch.ByteString (idl: ByteString).
func (_this *Request) Method() *patch.ByteString {
	var ret *patch.ByteString
	value := _this.Value_JS.Get("method")
	ret = patch.ByteStringFromJS(value)
	return ret
}

// Url returning attribute 'url' with
// type string (idl: USVString).
func (_this *Request) Url() string {
	var ret string
	value := _this.Value_JS.Get("url")
	ret = (value).String()
	return ret
}

// Headers returning attribute 'headers' with
// type Headers (idl: Headers).
func (_this *Request) Headers() *Headers {
	var ret *Headers
	value := _this.Value_JS.Get("headers")
	ret = HeadersFromJS(value)
	return ret
}

// Destination returning attribute 'destination' with
// type RequestDestination (idl: RequestDestination).
func (_this *Request) Destination() RequestDestination {
	var ret RequestDestination
	value := _this.Value_JS.Get("destination")
	ret = RequestDestinationFromJS(value)
	return ret
}

// Referrer returning attribute 'referrer' with
// type string (idl: USVString).
func (_this *Request) Referrer() string {
	var ret string
	value := _this.Value_JS.Get("referrer")
	ret = (value).String()
	return ret
}

// ReferrerPolicy returning attribute 'referrerPolicy' with
// type ReferrerPolicy (idl: ReferrerPolicy).
func (_this *Request) ReferrerPolicy() ReferrerPolicy {
	var ret ReferrerPolicy
	value := _this.Value_JS.Get("referrerPolicy")
	ret = ReferrerPolicyFromJS(value)
	return ret
}

// Mode returning attribute 'mode' with
// type RequestMode (idl: RequestMode).
func (_this *Request) Mode() RequestMode {
	var ret RequestMode
	value := _this.Value_JS.Get("mode")
	ret = RequestModeFromJS(value)
	return ret
}

// Credentials returning attribute 'credentials' with
// type RequestCredentials (idl: RequestCredentials).
func (_this *Request) Credentials() RequestCredentials {
	var ret RequestCredentials
	value := _this.Value_JS.Get("credentials")
	ret = RequestCredentialsFromJS(value)
	return ret
}

// Cache returning attribute 'cache' with
// type RequestCache (idl: RequestCache).
func (_this *Request) Cache() RequestCache {
	var ret RequestCache
	value := _this.Value_JS.Get("cache")
	ret = RequestCacheFromJS(value)
	return ret
}

// Redirect returning attribute 'redirect' with
// type RequestRedirect (idl: RequestRedirect).
func (_this *Request) Redirect() RequestRedirect {
	var ret RequestRedirect
	value := _this.Value_JS.Get("redirect")
	ret = RequestRedirectFromJS(value)
	return ret
}

// Integrity returning attribute 'integrity' with
// type string (idl: DOMString).
func (_this *Request) Integrity() string {
	var ret string
	value := _this.Value_JS.Get("integrity")
	ret = (value).String()
	return ret
}

// Keepalive returning attribute 'keepalive' with
// type bool (idl: boolean).
func (_this *Request) Keepalive() bool {
	var ret bool
	value := _this.Value_JS.Get("keepalive")
	ret = (value).Bool()
	return ret
}

// IsReloadNavigation returning attribute 'isReloadNavigation' with
// type bool (idl: boolean).
func (_this *Request) IsReloadNavigation() bool {
	var ret bool
	value := _this.Value_JS.Get("isReloadNavigation")
	ret = (value).Bool()
	return ret
}

// IsHistoryNavigation returning attribute 'isHistoryNavigation' with
// type bool (idl: boolean).
func (_this *Request) IsHistoryNavigation() bool {
	var ret bool
	value := _this.Value_JS.Get("isHistoryNavigation")
	ret = (value).Bool()
	return ret
}

// Signal returning attribute 'signal' with
// type domcore.AbortSignal (idl: AbortSignal).
func (_this *Request) Signal() *domcore.AbortSignal {
	var ret *domcore.AbortSignal
	value := _this.Value_JS.Get("signal")
	ret = domcore.AbortSignalFromJS(value)
	return ret
}

// Body returning attribute 'body' with
// type patch.ReadableStream (idl: ReadableStream).
func (_this *Request) Body() *patch.ReadableStream {
	var ret *patch.ReadableStream
	value := _this.Value_JS.Get("body")
	if value.Type() != js.TypeNull {
		ret = patch.ReadableStreamFromJS(value)
	}
	return ret
}

// BodyUsed returning attribute 'bodyUsed' with
// type bool (idl: boolean).
func (_this *Request) BodyUsed() bool {
	var ret bool
	value := _this.Value_JS.Get("bodyUsed")
	ret = (value).Bool()
	return ret
}

func (_this *Request) Clone() (_result *Request) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("clone", _args[0:_end]...)
	var (
		_converted *Request // javascript: Request _what_return_name
	)
	_converted = RequestFromJS(_returned)
	_result = _converted
	return
}

func (_this *Request) ArrayBuffer() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("arrayBuffer", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Request) Blob() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("blob", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Request) FormData() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("formData", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Request) Json() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("json", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Request) Text() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("text", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

// interface: Response
type Response struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Response) JSValue() js.Value {
	return _this.Value_JS
}

// ResponseFromJS is casting a js.Wrapper into Response.
func ResponseFromJS(value js.Wrapper) *Response {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Response{}
	ret.Value_JS = input
	return ret
}

func Error() (_result *Response) {
	_klass := js.Global().Get("Response")
	_method := _klass.Get("error")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _method.Invoke(_args[0:_end]...)
	var (
		_converted *Response // javascript: Response _what_return_name
	)
	_converted = ResponseFromJS(_returned)
	_result = _converted
	return
}

func Redirect(url string, status *int) (_result *Response) {
	_klass := js.Global().Get("Response")
	_method := _klass.Get("redirect")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := url
	_args[0] = _p0
	_end++
	if status != nil {
		_p1 := status
		_args[1] = _p1
		_end++
	}
	_returned := _method.Invoke(_args[0:_end]...)
	var (
		_converted *Response // javascript: Response _what_return_name
	)
	_converted = ResponseFromJS(_returned)
	_result = _converted
	return
}

func NewResponse(body *Union, init *ResponseInit) (_result *Response) {
	_klass := js.Global().Get("Response")
	var (
		_args [2]interface{}
		_end  int
	)
	if body != nil {
		_p0 := body.JSValue()
		_args[0] = _p0
		_end++
	}
	if init != nil {
		_p1 := init.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Response // javascript: Response _what_return_name
	)
	_converted = ResponseFromJS(_returned)
	_result = _converted
	return
}

// Type returning attribute 'type' with
// type ResponseType (idl: ResponseType).
func (_this *Response) Type() ResponseType {
	var ret ResponseType
	value := _this.Value_JS.Get("type")
	ret = ResponseTypeFromJS(value)
	return ret
}

// Url returning attribute 'url' with
// type string (idl: USVString).
func (_this *Response) Url() string {
	var ret string
	value := _this.Value_JS.Get("url")
	ret = (value).String()
	return ret
}

// Redirected returning attribute 'redirected' with
// type bool (idl: boolean).
func (_this *Response) Redirected() bool {
	var ret bool
	value := _this.Value_JS.Get("redirected")
	ret = (value).Bool()
	return ret
}

// Status returning attribute 'status' with
// type int (idl: unsigned short).
func (_this *Response) Status() int {
	var ret int
	value := _this.Value_JS.Get("status")
	ret = (value).Int()
	return ret
}

// Ok returning attribute 'ok' with
// type bool (idl: boolean).
func (_this *Response) Ok() bool {
	var ret bool
	value := _this.Value_JS.Get("ok")
	ret = (value).Bool()
	return ret
}

// StatusText returning attribute 'statusText' with
// type patch.ByteString (idl: ByteString).
func (_this *Response) StatusText() *patch.ByteString {
	var ret *patch.ByteString
	value := _this.Value_JS.Get("statusText")
	ret = patch.ByteStringFromJS(value)
	return ret
}

// Headers returning attribute 'headers' with
// type Headers (idl: Headers).
func (_this *Response) Headers() *Headers {
	var ret *Headers
	value := _this.Value_JS.Get("headers")
	ret = HeadersFromJS(value)
	return ret
}

// Trailer returning attribute 'trailer' with
// type javascript.Promise (idl: Promise).
func (_this *Response) Trailer() *javascript.Promise {
	var ret *javascript.Promise
	value := _this.Value_JS.Get("trailer")
	ret = javascript.PromiseFromJS(value)
	return ret
}

// Body returning attribute 'body' with
// type patch.ReadableStream (idl: ReadableStream).
func (_this *Response) Body() *patch.ReadableStream {
	var ret *patch.ReadableStream
	value := _this.Value_JS.Get("body")
	if value.Type() != js.TypeNull {
		ret = patch.ReadableStreamFromJS(value)
	}
	return ret
}

// BodyUsed returning attribute 'bodyUsed' with
// type bool (idl: boolean).
func (_this *Response) BodyUsed() bool {
	var ret bool
	value := _this.Value_JS.Get("bodyUsed")
	ret = (value).Bool()
	return ret
}

func (_this *Response) Clone() (_result *Response) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("clone", _args[0:_end]...)
	var (
		_converted *Response // javascript: Response _what_return_name
	)
	_converted = ResponseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Response) ArrayBuffer() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("arrayBuffer", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Response) Blob() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("blob", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Response) FormData() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("formData", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Response) Json() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("json", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}

func (_this *Response) Text() (_result *javascript.Promise) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("text", _args[0:_end]...)
	var (
		_converted *javascript.Promise // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseFromJS(_returned)
	_result = _converted
	return
}
