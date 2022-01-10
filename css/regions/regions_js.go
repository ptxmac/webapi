// Code generated by webidl-bind. DO NOT EDIT.

package regions

import "syscall/js"

import (
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/dom/domcore"
)

// using following types:
// dom.Node
// dom.Range
// domcore.EventTarget

// source idl files:
// css-regions.idl

// transform files:
// css-regions.go.md

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

// class: NamedFlow
type NamedFlow struct {
	domcore.EventTarget
}

// NamedFlowFromJS is casting a js.Value into NamedFlow.
func NamedFlowFromJS(value js.Value) *NamedFlow {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &NamedFlow{}
	ret.Value_JS = value
	return ret
}

// NamedFlowFromJS is casting from something that holds a js.Value into NamedFlow.
func NamedFlowFromWrapper(input core.Wrapper) *NamedFlow {
	return NamedFlowFromJS(input.JSValue())
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *NamedFlow) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// Overset returning attribute 'overset' with
// type bool (idl: boolean).
func (_this *NamedFlow) Overset() bool {
	var ret bool
	value := _this.Value_JS.Get("overset")
	ret = (value).Bool()
	return ret
}

// FirstEmptyRegionIndex returning attribute 'firstEmptyRegionIndex' with
// type int (idl: short).
func (_this *NamedFlow) FirstEmptyRegionIndex() int {
	var ret int
	value := _this.Value_JS.Get("firstEmptyRegionIndex")
	ret = (value).Int()
	return ret
}

func (_this *NamedFlow) GetRegions() (_result []*Region) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getRegions", _args[0:_end]...)
	var (
		_converted []*Region // javascript: sequence<Region> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*Region, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *Region
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = RegionFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *NamedFlow) GetContent() (_result []*dom.Node) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getContent", _args[0:_end]...)
	var (
		_converted []*dom.Node // javascript: sequence<Node> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*dom.Node, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *dom.Node
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = dom.NodeFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *NamedFlow) GetRegionsByContent(node *dom.Node) (_result []*Region) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := node.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getRegionsByContent", _args[0:_end]...)
	var (
		_converted []*Region // javascript: sequence<Region> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*Region, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *Region
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = RegionFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

// class: NamedFlowMap
type NamedFlowMap struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *NamedFlowMap) JSValue() js.Value {
	return _this.Value_JS
}

// NamedFlowMapFromJS is casting a js.Value into NamedFlowMap.
func NamedFlowMapFromJS(value js.Value) *NamedFlowMap {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &NamedFlowMap{}
	ret.Value_JS = value
	return ret
}

// NamedFlowMapFromJS is casting from something that holds a js.Value into NamedFlowMap.
func NamedFlowMapFromWrapper(input core.Wrapper) *NamedFlowMap {
	return NamedFlowMapFromJS(input.JSValue())
}

func (_this *NamedFlowMap) Get(flowName string) (_result *NamedFlow) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := flowName
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *NamedFlow // javascript: NamedFlow _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = NamedFlowFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NamedFlowMap) Has(flowName string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := flowName
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

func (_this *NamedFlowMap) Set(flowName string, flowValue *NamedFlow) (_result *NamedFlowMap) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := flowName
	_args[0] = _p0
	_end++
	_p1 := flowValue.JSValue()
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("set", _args[0:_end]...)
	var (
		_converted *NamedFlowMap // javascript: NamedFlowMap _what_return_name
	)
	_converted = NamedFlowMapFromJS(_returned)
	_result = _converted
	return
}

func (_this *NamedFlowMap) Delete(flowName string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := flowName
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

// class: Region
type Region struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Region) JSValue() js.Value {
	return _this.Value_JS
}

// RegionFromJS is casting a js.Value into Region.
func RegionFromJS(value js.Value) *Region {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Region{}
	ret.Value_JS = value
	return ret
}

// RegionFromJS is casting from something that holds a js.Value into Region.
func RegionFromWrapper(input core.Wrapper) *Region {
	return RegionFromJS(input.JSValue())
}

// RegionOverset returning attribute 'regionOverset' with
// type string (idl: DOMString).
func (_this *Region) RegionOverset() string {
	var ret string
	value := _this.Value_JS.Get("regionOverset")
	ret = (value).String()
	return ret
}

func (_this *Region) GetRegionFlowRanges() (_result []*dom.Range) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getRegionFlowRanges", _args[0:_end]...)
	var (
		_converted []*dom.Range // javascript: sequence<Range> _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		__length0 := _returned.Length()
		__array0 := make([]*dom.Range, __length0, __length0)
		for __idx0 := 0; __idx0 < __length0; __idx0++ {
			var __seq_out0 *dom.Range
			__seq_in0 := _returned.Index(__idx0)
			__seq_out0 = dom.RangeFromJS(__seq_in0)
			__array0[__idx0] = __seq_out0
		}
		_converted = __array0
	}
	_result = _converted
	return
}
