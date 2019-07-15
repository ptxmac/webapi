// Code generated by webidl-bind. DO NOT EDIT.

package reporting

import "syscall/js"

import (
	"github.com/gowebapi/webapi/javascript/missingtypes"
)

// using following types:
// missingtypes.Date

// source idl files:
// reporting.idl

// transform files:
// reporting.go.md

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

// callback: ReportingObserverCallback
type ReportingObserverCallbackFunc func(reports []*Report, observer *ReportingObserver)

// ReportingObserverCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type ReportingObserverCallback js.Func

func ReportingObserverCallbackToJS(callback ReportingObserverCallbackFunc) *ReportingObserverCallback {
	if callback == nil {
		return nil
	}
	ret := ReportingObserverCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 []*Report          // javascript: sequence<Report> reports
			_p1 *ReportingObserver // javascript: ReportingObserver observer
		)
		__length0 := args[0].Length()
		__array0 := make([]*Report, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 *Report
			__seq_in0 := args[0].Index(__idx0)
			__seq_out0 = ReportFromJS(__seq_in0)
			__array0[__idx0] = __seq_out0
		}
		_p0 = __array0
		_p1 = ReportingObserverFromJS(args[1])
		callback(_p0, _p1)
		// returning no return value
		return nil
	}))
	return &ret
}

func ReportingObserverCallbackFromJS(_value js.Value) ReportingObserverCallbackFunc {
	return func(reports []*Report, observer *ReportingObserver) {
		var (
			_args [2]interface{}
			_end  int
		)
		_p0 := js.Global().Get("Array").New(len(reports))
		for __idx0, __seq_in0 := range reports {
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

// dictionary: ReportingObserverOptions
type ReportingObserverOptions struct {
	Types    []string
	Buffered bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ReportingObserverOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Types))
	for __idx0, __seq_in0 := range _this.Types {
		__seq_out0 := __seq_in0
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("types", value0)
	value1 := _this.Buffered
	out.Set("buffered", value1)
	return out
}

// ReportingObserverOptionsFromJS is allocating a new
// ReportingObserverOptions object and copy all values from
// input javascript object
func ReportingObserverOptionsFromJS(value js.Wrapper) *ReportingObserverOptions {
	input := value.JSValue()
	var out ReportingObserverOptions
	var (
		value0 []string // javascript: sequence<DOMString> {types Types types}
		value1 bool     // javascript: boolean {buffered Buffered buffered}
	)
	__length0 := input.Get("types").Length()
	__array0 := make([]string, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 string
		__seq_in0 := input.Get("types").Index(__idx0)
		__seq_out0 = (__seq_in0).String()
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Types = value0
	value1 = (input.Get("buffered")).Bool()
	out.Buffered = value1
	return &out
}

// interface: CrashReportBody
type CrashReportBody struct {
	ReportBody
}

// CrashReportBodyFromJS is casting a js.Wrapper into CrashReportBody.
func CrashReportBodyFromJS(value js.Wrapper) *CrashReportBody {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CrashReportBody{}
	ret.Value_JS = input
	return ret
}

// Reason returning attribute 'reason' with
// type string (idl: DOMString).
func (_this *CrashReportBody) Reason() *string {
	var ret *string
	value := _this.Value_JS.Get("reason")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// interface: DeprecationReportBody
type DeprecationReportBody struct {
	ReportBody
}

// DeprecationReportBodyFromJS is casting a js.Wrapper into DeprecationReportBody.
func DeprecationReportBodyFromJS(value js.Wrapper) *DeprecationReportBody {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DeprecationReportBody{}
	ret.Value_JS = input
	return ret
}

// Id returning attribute 'id' with
// type string (idl: DOMString).
func (_this *DeprecationReportBody) Id() string {
	var ret string
	value := _this.Value_JS.Get("id")
	ret = (value).String()
	return ret
}

// AnticipatedRemoval returning attribute 'anticipatedRemoval' with
// type missingtypes.Date (idl: Date).
func (_this *DeprecationReportBody) AnticipatedRemoval() *missingtypes.Date {
	var ret *missingtypes.Date
	value := _this.Value_JS.Get("anticipatedRemoval")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = missingtypes.DateFromJS(value)
	}
	return ret
}

// Message returning attribute 'message' with
// type string (idl: DOMString).
func (_this *DeprecationReportBody) Message() string {
	var ret string
	value := _this.Value_JS.Get("message")
	ret = (value).String()
	return ret
}

// SourceFile returning attribute 'sourceFile' with
// type string (idl: DOMString).
func (_this *DeprecationReportBody) SourceFile() *string {
	var ret *string
	value := _this.Value_JS.Get("sourceFile")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// LineNumber returning attribute 'lineNumber' with
// type uint (idl: unsigned long).
func (_this *DeprecationReportBody) LineNumber() *uint {
	var ret *uint
	value := _this.Value_JS.Get("lineNumber")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (uint)((value).Int())
		ret = &__tmp
	}
	return ret
}

// ColumnNumber returning attribute 'columnNumber' with
// type uint (idl: unsigned long).
func (_this *DeprecationReportBody) ColumnNumber() *uint {
	var ret *uint
	value := _this.Value_JS.Get("columnNumber")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (uint)((value).Int())
		ret = &__tmp
	}
	return ret
}

// interface: InterventionReportBody
type InterventionReportBody struct {
	ReportBody
}

// InterventionReportBodyFromJS is casting a js.Wrapper into InterventionReportBody.
func InterventionReportBodyFromJS(value js.Wrapper) *InterventionReportBody {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &InterventionReportBody{}
	ret.Value_JS = input
	return ret
}

// Id returning attribute 'id' with
// type string (idl: DOMString).
func (_this *InterventionReportBody) Id() string {
	var ret string
	value := _this.Value_JS.Get("id")
	ret = (value).String()
	return ret
}

// Message returning attribute 'message' with
// type string (idl: DOMString).
func (_this *InterventionReportBody) Message() string {
	var ret string
	value := _this.Value_JS.Get("message")
	ret = (value).String()
	return ret
}

// SourceFile returning attribute 'sourceFile' with
// type string (idl: DOMString).
func (_this *InterventionReportBody) SourceFile() *string {
	var ret *string
	value := _this.Value_JS.Get("sourceFile")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// LineNumber returning attribute 'lineNumber' with
// type uint (idl: unsigned long).
func (_this *InterventionReportBody) LineNumber() *uint {
	var ret *uint
	value := _this.Value_JS.Get("lineNumber")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (uint)((value).Int())
		ret = &__tmp
	}
	return ret
}

// ColumnNumber returning attribute 'columnNumber' with
// type uint (idl: unsigned long).
func (_this *InterventionReportBody) ColumnNumber() *uint {
	var ret *uint
	value := _this.Value_JS.Get("columnNumber")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (uint)((value).Int())
		ret = &__tmp
	}
	return ret
}

// interface: Report
type Report struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Report) JSValue() js.Value {
	return _this.Value_JS
}

// ReportFromJS is casting a js.Wrapper into Report.
func ReportFromJS(value js.Wrapper) *Report {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Report{}
	ret.Value_JS = input
	return ret
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *Report) Type() string {
	var ret string
	value := _this.Value_JS.Get("type")
	ret = (value).String()
	return ret
}

// Url returning attribute 'url' with
// type string (idl: DOMString).
func (_this *Report) Url() string {
	var ret string
	value := _this.Value_JS.Get("url")
	ret = (value).String()
	return ret
}

// Body returning attribute 'body' with
// type ReportBody (idl: ReportBody).
func (_this *Report) Body() *ReportBody {
	var ret *ReportBody
	value := _this.Value_JS.Get("body")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = ReportBodyFromJS(value)
	}
	return ret
}

// interface: ReportBody
type ReportBody struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *ReportBody) JSValue() js.Value {
	return _this.Value_JS
}

// ReportBodyFromJS is casting a js.Wrapper into ReportBody.
func ReportBodyFromJS(value js.Wrapper) *ReportBody {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ReportBody{}
	ret.Value_JS = input
	return ret
}

// interface: ReportingObserver
type ReportingObserver struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *ReportingObserver) JSValue() js.Value {
	return _this.Value_JS
}

// ReportingObserverFromJS is casting a js.Wrapper into ReportingObserver.
func ReportingObserverFromJS(value js.Wrapper) *ReportingObserver {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ReportingObserver{}
	ret.Value_JS = input
	return ret
}

func NewReportingObserver(callback *ReportingObserverCallback, options *ReportingObserverOptions) (_result *ReportingObserver) {
	_klass := js.Global().Get("ReportingObserver")
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
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ReportingObserver // javascript: ReportingObserver _what_return_name
	)
	_converted = ReportingObserverFromJS(_returned)
	_result = _converted
	return
}

func (_this *ReportingObserver) Observe() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("observe", _args[0:_end]...)
	return
}

func (_this *ReportingObserver) Disconnect() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("disconnect", _args[0:_end]...)
	return
}

func (_this *ReportingObserver) TakeRecords() (_result []*Report) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("takeRecords", _args[0:_end]...)
	var (
		_converted []*Report // javascript: sequence<Report> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*Report, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *Report
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = ReportFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}
