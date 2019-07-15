// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package featurepolicy

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/reporting"
)

// using following types:
// reporting.ReportBody

// source idl files:
// webappsec-feature-policy.idl

// transform files:
// webappsec-feature-policy.go.md

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

// interface: FeaturePolicy
type FeaturePolicy struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FeaturePolicy) JSValue() js.Value {
	return _this.Value_JS
}

// FeaturePolicyFromJS is casting a js.Wrapper into FeaturePolicy.
func FeaturePolicyFromJS(value js.Wrapper) *FeaturePolicy {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FeaturePolicy{}
	ret.Value_JS = input
	return ret
}

func (_this *FeaturePolicy) AllowsFeature(feature string, origin *string) (_result bool) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := feature
	_args[0] = _p0
	_end++
	if origin != nil {
		_p1 := origin
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("allowsFeature", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *FeaturePolicy) Features() (_result []string) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("features", _args[0:_end]...)
	var (
		_converted []string // javascript: sequence<DOMString> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]string, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 string
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = (__seq_in0).String()
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *FeaturePolicy) AllowedFeatures() (_result []string) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("allowedFeatures", _args[0:_end]...)
	var (
		_converted []string // javascript: sequence<DOMString> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]string, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 string
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = (__seq_in0).String()
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *FeaturePolicy) GetAllowlistForFeature(feature string) (_result []string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := feature
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getAllowlistForFeature", _args[0:_end]...)
	var (
		_converted []string // javascript: sequence<DOMString> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]string, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 string
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = (__seq_in0).String()
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

// interface: FeaturePolicyViolationReportBody
type ViolationReportBody struct {
	reporting.ReportBody
}

// ViolationReportBodyFromJS is casting a js.Wrapper into ViolationReportBody.
func ViolationReportBodyFromJS(value js.Wrapper) *ViolationReportBody {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ViolationReportBody{}
	ret.Value_JS = input
	return ret
}

// FeatureId returning attribute 'featureId' with
// type string (idl: DOMString).
func (_this *ViolationReportBody) FeatureId() string {
	var ret string
	value := _this.Value_JS.Get("featureId")
	ret = (value).String()
	return ret
}

// SourceFile returning attribute 'sourceFile' with
// type string (idl: DOMString).
func (_this *ViolationReportBody) SourceFile() *string {
	var ret *string
	value := _this.Value_JS.Get("sourceFile")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// LineNumber returning attribute 'lineNumber' with
// type int (idl: long).
func (_this *ViolationReportBody) LineNumber() *int {
	var ret *int
	value := _this.Value_JS.Get("lineNumber")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).Int()
		ret = &__tmp
	}
	return ret
}

// ColumnNumber returning attribute 'columnNumber' with
// type int (idl: long).
func (_this *ViolationReportBody) ColumnNumber() *int {
	var ret *int
	value := _this.Value_JS.Get("columnNumber")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).Int()
		ret = &__tmp
	}
	return ret
}

// Disposition returning attribute 'disposition' with
// type string (idl: DOMString).
func (_this *ViolationReportBody) Disposition() string {
	var ret string
	value := _this.Value_JS.Get("disposition")
	ret = (value).String()
	return ret
}
