// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package scrollanimations

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/css/animations/webani"
	"github.com/gowebapi/webapi/dom"
)

// using following types:
// dom.Element
// webani.AnimationTimeline
// webani.FillMode

// source idl files:
// scroll-animations.idl

// transform files:
// scroll-animations.go.md

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

// enum: ScrollDirection
type ScrollDirection int

const (
	BlockScrollDirection ScrollDirection = iota
	InlineScrollDirection
	HorizontalScrollDirection
	VerticalScrollDirection
)

var scrollDirectionToWasmTable = []string{
	"block", "inline", "horizontal", "vertical",
}

var scrollDirectionFromWasmTable = map[string]ScrollDirection{
	"block": BlockScrollDirection, "inline": InlineScrollDirection, "horizontal": HorizontalScrollDirection, "vertical": VerticalScrollDirection,
}

// JSValue is converting this enum into a javascript object
func (this *ScrollDirection) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ScrollDirection) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(scrollDirectionToWasmTable) {
		return scrollDirectionToWasmTable[idx]
	}
	panic("unknown input value")
}

// ScrollDirectionFromJS is converting a javascript value into
// a ScrollDirection enum value.
func ScrollDirectionFromJS(value js.Value) ScrollDirection {
	key := value.String()
	conv, ok := scrollDirectionFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: ScrollTimelineAutoKeyword
type ScrollTimelineAutoKeyword int

const (
	AutoScrollTimelineAutoKeyword ScrollTimelineAutoKeyword = iota
)

var scrollTimelineAutoKeywordToWasmTable = []string{
	"auto",
}

var scrollTimelineAutoKeywordFromWasmTable = map[string]ScrollTimelineAutoKeyword{
	"auto": AutoScrollTimelineAutoKeyword,
}

// JSValue is converting this enum into a javascript object
func (this *ScrollTimelineAutoKeyword) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ScrollTimelineAutoKeyword) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(scrollTimelineAutoKeywordToWasmTable) {
		return scrollTimelineAutoKeywordToWasmTable[idx]
	}
	panic("unknown input value")
}

// ScrollTimelineAutoKeywordFromJS is converting a javascript value into
// a ScrollTimelineAutoKeyword enum value.
func ScrollTimelineAutoKeywordFromJS(value js.Value) ScrollTimelineAutoKeyword {
	key := value.String()
	conv, ok := scrollTimelineAutoKeywordFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: ScrollTimelineOptions
type ScrollTimelineOptions struct {
	ScrollSource      *dom.Element
	Orientation       ScrollDirection
	StartScrollOffset string
	EndScrollOffset   string
	TimeRange         *Union
	Fill              webani.FillMode
}

// JSValue is allocating a new javascript object and copy
// all values
func (_this *ScrollTimelineOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.ScrollSource.JSValue()
	out.Set("scrollSource", value0)
	value1 := _this.Orientation.JSValue()
	out.Set("orientation", value1)
	value2 := _this.StartScrollOffset
	out.Set("startScrollOffset", value2)
	value3 := _this.EndScrollOffset
	out.Set("endScrollOffset", value3)
	value4 := _this.TimeRange.JSValue()
	out.Set("timeRange", value4)
	value5 := _this.Fill.JSValue()
	out.Set("fill", value5)
	return out
}

// ScrollTimelineOptionsFromJS is allocating a new
// ScrollTimelineOptions object and copy all values in the value javascript object.
func ScrollTimelineOptionsFromJS(value js.Value) *ScrollTimelineOptions {
	var out ScrollTimelineOptions
	var (
		value0 *dom.Element    // javascript: Element {scrollSource ScrollSource scrollSource}
		value1 ScrollDirection // javascript: ScrollDirection {orientation Orientation orientation}
		value2 string          // javascript: DOMString {startScrollOffset StartScrollOffset startScrollOffset}
		value3 string          // javascript: DOMString {endScrollOffset EndScrollOffset endScrollOffset}
		value4 *Union          // javascript: Union {timeRange TimeRange timeRange}
		value5 webani.FillMode // javascript: FillMode {fill Fill fill}
	)
	if value.Get("scrollSource").Type() != js.TypeNull && value.Get("scrollSource").Type() != js.TypeUndefined {
		value0 = dom.ElementFromJS(value.Get("scrollSource"))
	}
	out.ScrollSource = value0
	value1 = ScrollDirectionFromJS(value.Get("orientation"))
	out.Orientation = value1
	value2 = (value.Get("startScrollOffset")).String()
	out.StartScrollOffset = value2
	value3 = (value.Get("endScrollOffset")).String()
	out.EndScrollOffset = value3
	value4 = UnionFromJS(value.Get("timeRange"))
	out.TimeRange = value4
	value5 = webani.FillModeFromJS(value.Get("fill"))
	out.Fill = value5
	return &out
}

// class: ScrollTimeline
type ScrollTimeline struct {
	webani.AnimationTimeline
}

// ScrollTimelineFromJS is casting a js.Value into ScrollTimeline.
func ScrollTimelineFromJS(value js.Value) *ScrollTimeline {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &ScrollTimeline{}
	ret.Value_JS = value
	return ret
}

// ScrollTimelineFromJS is casting from something that holds a js.Value into ScrollTimeline.
func ScrollTimelineFromWrapper(input core.Wrapper) *ScrollTimeline {
	return ScrollTimelineFromJS(input.JSValue())
}

func NewScrollTimeline(options *ScrollTimelineOptions) (_result *ScrollTimeline) {
	_klass := js.Global().Get("ScrollTimeline")
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
		_converted *ScrollTimeline // javascript: ScrollTimeline _what_return_name
	)
	_converted = ScrollTimelineFromJS(_returned)
	_result = _converted
	return
}

// ScrollSource returning attribute 'scrollSource' with
// type dom.Element (idl: Element).
func (_this *ScrollTimeline) ScrollSource() *dom.Element {
	var ret *dom.Element
	value := _this.Value_JS.Get("scrollSource")
	ret = dom.ElementFromJS(value)
	return ret
}

// Orientation returning attribute 'orientation' with
// type ScrollDirection (idl: ScrollDirection).
func (_this *ScrollTimeline) Orientation() ScrollDirection {
	var ret ScrollDirection
	value := _this.Value_JS.Get("orientation")
	ret = ScrollDirectionFromJS(value)
	return ret
}

// StartScrollOffset returning attribute 'startScrollOffset' with
// type string (idl: DOMString).
func (_this *ScrollTimeline) StartScrollOffset() string {
	var ret string
	value := _this.Value_JS.Get("startScrollOffset")
	ret = (value).String()
	return ret
}

// EndScrollOffset returning attribute 'endScrollOffset' with
// type string (idl: DOMString).
func (_this *ScrollTimeline) EndScrollOffset() string {
	var ret string
	value := _this.Value_JS.Get("endScrollOffset")
	ret = (value).String()
	return ret
}

// TimeRange returning attribute 'timeRange' with
// type Union (idl: Union).
func (_this *ScrollTimeline) TimeRange() *Union {
	var ret *Union
	value := _this.Value_JS.Get("timeRange")
	ret = UnionFromJS(value)
	return ret
}

// Fill returning attribute 'fill' with
// type webani.FillMode (idl: FillMode).
func (_this *ScrollTimeline) Fill() webani.FillMode {
	var ret webani.FillMode
	value := _this.Value_JS.Get("fill")
	ret = webani.FillModeFromJS(value)
	return ret
}
