// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package webvtt

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/html/media"
)

// using following types:
// dom.DocumentFragment
// media.TextTrackCue

// source idl files:
// webvtt.idl

// transform files:
// webvtt.go.md

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

// enum: AlignSetting
type AlignSetting int

const (
	StartAlignSetting AlignSetting = iota
	CenterAlignSetting
	EndAlignSetting
	LeftAlignSetting
	RightAlignSetting
)

var alignSettingToWasmTable = []string{
	"start", "center", "end", "left", "right",
}

var alignSettingFromWasmTable = map[string]AlignSetting{
	"start": StartAlignSetting, "center": CenterAlignSetting, "end": EndAlignSetting, "left": LeftAlignSetting, "right": RightAlignSetting,
}

// JSValue is converting this enum into a javascript object
func (this *AlignSetting) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this AlignSetting) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(alignSettingToWasmTable) {
		return alignSettingToWasmTable[idx]
	}
	panic("unknown input value")
}

// AlignSettingFromJS is converting a javascript value into
// a AlignSetting enum value.
func AlignSettingFromJS(value js.Value) AlignSetting {
	key := value.String()
	conv, ok := alignSettingFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: AutoKeyword
type AutoKeyword int

const (
	AutoAutoKeyword AutoKeyword = iota
)

var autoKeywordToWasmTable = []string{
	"auto",
}

var autoKeywordFromWasmTable = map[string]AutoKeyword{
	"auto": AutoAutoKeyword,
}

// JSValue is converting this enum into a javascript object
func (this *AutoKeyword) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this AutoKeyword) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(autoKeywordToWasmTable) {
		return autoKeywordToWasmTable[idx]
	}
	panic("unknown input value")
}

// AutoKeywordFromJS is converting a javascript value into
// a AutoKeyword enum value.
func AutoKeywordFromJS(value js.Value) AutoKeyword {
	key := value.String()
	conv, ok := autoKeywordFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: DirectionSetting
type DirectionSetting int

const (
	EmptyString0DirectionSetting DirectionSetting = iota
	RlDirectionSetting
	LrDirectionSetting
)

var directionSettingToWasmTable = []string{
	"", "rl", "lr",
}

var directionSettingFromWasmTable = map[string]DirectionSetting{
	"": EmptyString0DirectionSetting, "rl": RlDirectionSetting, "lr": LrDirectionSetting,
}

// JSValue is converting this enum into a javascript object
func (this *DirectionSetting) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this DirectionSetting) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(directionSettingToWasmTable) {
		return directionSettingToWasmTable[idx]
	}
	panic("unknown input value")
}

// DirectionSettingFromJS is converting a javascript value into
// a DirectionSetting enum value.
func DirectionSettingFromJS(value js.Value) DirectionSetting {
	key := value.String()
	conv, ok := directionSettingFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: LineAlignSetting
type LineAlignSetting int

const (
	StartLineAlignSetting LineAlignSetting = iota
	CenterLineAlignSetting
	EndLineAlignSetting
)

var lineAlignSettingToWasmTable = []string{
	"start", "center", "end",
}

var lineAlignSettingFromWasmTable = map[string]LineAlignSetting{
	"start": StartLineAlignSetting, "center": CenterLineAlignSetting, "end": EndLineAlignSetting,
}

// JSValue is converting this enum into a javascript object
func (this *LineAlignSetting) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this LineAlignSetting) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(lineAlignSettingToWasmTable) {
		return lineAlignSettingToWasmTable[idx]
	}
	panic("unknown input value")
}

// LineAlignSettingFromJS is converting a javascript value into
// a LineAlignSetting enum value.
func LineAlignSettingFromJS(value js.Value) LineAlignSetting {
	key := value.String()
	conv, ok := lineAlignSettingFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: PositionAlignSetting
type PositionAlignSetting int

const (
	LineLeftPositionAlignSetting PositionAlignSetting = iota
	CenterPositionAlignSetting
	LineRightPositionAlignSetting
	AutoPositionAlignSetting
)

var positionAlignSettingToWasmTable = []string{
	"line-left", "center", "line-right", "auto",
}

var positionAlignSettingFromWasmTable = map[string]PositionAlignSetting{
	"line-left": LineLeftPositionAlignSetting, "center": CenterPositionAlignSetting, "line-right": LineRightPositionAlignSetting, "auto": AutoPositionAlignSetting,
}

// JSValue is converting this enum into a javascript object
func (this *PositionAlignSetting) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this PositionAlignSetting) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(positionAlignSettingToWasmTable) {
		return positionAlignSettingToWasmTable[idx]
	}
	panic("unknown input value")
}

// PositionAlignSettingFromJS is converting a javascript value into
// a PositionAlignSetting enum value.
func PositionAlignSettingFromJS(value js.Value) PositionAlignSetting {
	key := value.String()
	conv, ok := positionAlignSettingFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: ScrollSetting
type ScrollSetting int

const (
	EmptyString0ScrollSetting ScrollSetting = iota
	UpScrollSetting
)

var scrollSettingToWasmTable = []string{
	"", "up",
}

var scrollSettingFromWasmTable = map[string]ScrollSetting{
	"": EmptyString0ScrollSetting, "up": UpScrollSetting,
}

// JSValue is converting this enum into a javascript object
func (this *ScrollSetting) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ScrollSetting) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(scrollSettingToWasmTable) {
		return scrollSettingToWasmTable[idx]
	}
	panic("unknown input value")
}

// ScrollSettingFromJS is converting a javascript value into
// a ScrollSetting enum value.
func ScrollSettingFromJS(value js.Value) ScrollSetting {
	key := value.String()
	conv, ok := scrollSettingFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// class: VTTCue
type VTTCue struct {
	media.TextTrackCue
}

// VTTCueFromJS is casting a js.Value into VTTCue.
func VTTCueFromJS(value js.Value) *VTTCue {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &VTTCue{}
	ret.Value_JS = value
	return ret
}

// VTTCueFromJS is casting from something that holds a js.Value into VTTCue.
func VTTCueFromWrapper(input core.Wrapper) *VTTCue {
	return VTTCueFromJS(input.JSValue())
}

func NewVTTCue(startTime float64, endTime float64, text string) (_result *VTTCue) {
	_klass := js.Global().Get("VTTCue")
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := startTime
	_args[0] = _p0
	_end++
	_p1 := endTime
	_args[1] = _p1
	_end++
	_p2 := text
	_args[2] = _p2
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *VTTCue // javascript: VTTCue _what_return_name
	)
	_converted = VTTCueFromJS(_returned)
	_result = _converted
	return
}

// Region returning attribute 'region' with
// type VTTRegion (idl: VTTRegion).
func (_this *VTTCue) Region() *VTTRegion {
	var ret *VTTRegion
	value := _this.Value_JS.Get("region")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = VTTRegionFromJS(value)
	}
	return ret
}

// SetRegion setting attribute 'region' with
// type VTTRegion (idl: VTTRegion).
func (_this *VTTCue) SetRegion(value *VTTRegion) {
	input := value.JSValue()
	_this.Value_JS.Set("region", input)
}

// Vertical returning attribute 'vertical' with
// type DirectionSetting (idl: DirectionSetting).
func (_this *VTTCue) Vertical() DirectionSetting {
	var ret DirectionSetting
	value := _this.Value_JS.Get("vertical")
	ret = DirectionSettingFromJS(value)
	return ret
}

// SetVertical setting attribute 'vertical' with
// type DirectionSetting (idl: DirectionSetting).
func (_this *VTTCue) SetVertical(value DirectionSetting) {
	input := value.JSValue()
	_this.Value_JS.Set("vertical", input)
}

// SnapToLines returning attribute 'snapToLines' with
// type bool (idl: boolean).
func (_this *VTTCue) SnapToLines() bool {
	var ret bool
	value := _this.Value_JS.Get("snapToLines")
	ret = (value).Bool()
	return ret
}

// SetSnapToLines setting attribute 'snapToLines' with
// type bool (idl: boolean).
func (_this *VTTCue) SetSnapToLines(value bool) {
	input := value
	_this.Value_JS.Set("snapToLines", input)
}

// Line returning attribute 'line' with
// type Union (idl: Union).
func (_this *VTTCue) Line() *Union {
	var ret *Union
	value := _this.Value_JS.Get("line")
	ret = UnionFromJS(value)
	return ret
}

// SetLine setting attribute 'line' with
// type Union (idl: Union).
func (_this *VTTCue) SetLine(value *Union) {
	input := value.JSValue()
	_this.Value_JS.Set("line", input)
}

// LineAlign returning attribute 'lineAlign' with
// type LineAlignSetting (idl: LineAlignSetting).
func (_this *VTTCue) LineAlign() LineAlignSetting {
	var ret LineAlignSetting
	value := _this.Value_JS.Get("lineAlign")
	ret = LineAlignSettingFromJS(value)
	return ret
}

// SetLineAlign setting attribute 'lineAlign' with
// type LineAlignSetting (idl: LineAlignSetting).
func (_this *VTTCue) SetLineAlign(value LineAlignSetting) {
	input := value.JSValue()
	_this.Value_JS.Set("lineAlign", input)
}

// Position returning attribute 'position' with
// type Union (idl: Union).
func (_this *VTTCue) Position() *Union {
	var ret *Union
	value := _this.Value_JS.Get("position")
	ret = UnionFromJS(value)
	return ret
}

// SetPosition setting attribute 'position' with
// type Union (idl: Union).
func (_this *VTTCue) SetPosition(value *Union) {
	input := value.JSValue()
	_this.Value_JS.Set("position", input)
}

// PositionAlign returning attribute 'positionAlign' with
// type PositionAlignSetting (idl: PositionAlignSetting).
func (_this *VTTCue) PositionAlign() PositionAlignSetting {
	var ret PositionAlignSetting
	value := _this.Value_JS.Get("positionAlign")
	ret = PositionAlignSettingFromJS(value)
	return ret
}

// SetPositionAlign setting attribute 'positionAlign' with
// type PositionAlignSetting (idl: PositionAlignSetting).
func (_this *VTTCue) SetPositionAlign(value PositionAlignSetting) {
	input := value.JSValue()
	_this.Value_JS.Set("positionAlign", input)
}

// Size returning attribute 'size' with
// type float64 (idl: double).
func (_this *VTTCue) Size() float64 {
	var ret float64
	value := _this.Value_JS.Get("size")
	ret = (value).Float()
	return ret
}

// SetSize setting attribute 'size' with
// type float64 (idl: double).
func (_this *VTTCue) SetSize(value float64) {
	input := value
	_this.Value_JS.Set("size", input)
}

// Align returning attribute 'align' with
// type AlignSetting (idl: AlignSetting).
func (_this *VTTCue) Align() AlignSetting {
	var ret AlignSetting
	value := _this.Value_JS.Get("align")
	ret = AlignSettingFromJS(value)
	return ret
}

// SetAlign setting attribute 'align' with
// type AlignSetting (idl: AlignSetting).
func (_this *VTTCue) SetAlign(value AlignSetting) {
	input := value.JSValue()
	_this.Value_JS.Set("align", input)
}

// Text returning attribute 'text' with
// type string (idl: DOMString).
func (_this *VTTCue) Text() string {
	var ret string
	value := _this.Value_JS.Get("text")
	ret = (value).String()
	return ret
}

// SetText setting attribute 'text' with
// type string (idl: DOMString).
func (_this *VTTCue) SetText(value string) {
	input := value
	_this.Value_JS.Set("text", input)
}

func (_this *VTTCue) GetCueAsHTML() (_result *dom.DocumentFragment) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getCueAsHTML", _args[0:_end]...)
	var (
		_converted *dom.DocumentFragment // javascript: DocumentFragment _what_return_name
	)
	_converted = dom.DocumentFragmentFromJS(_returned)
	_result = _converted
	return
}

// class: VTTRegion
type VTTRegion struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *VTTRegion) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// VTTRegionFromJS is casting a js.Value into VTTRegion.
func VTTRegionFromJS(value js.Value) *VTTRegion {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &VTTRegion{}
	ret.Value_JS = value
	return ret
}

// VTTRegionFromJS is casting from something that holds a js.Value into VTTRegion.
func VTTRegionFromWrapper(input core.Wrapper) *VTTRegion {
	return VTTRegionFromJS(input.JSValue())
}

func NewVTTRegion() (_result *VTTRegion) {
	_klass := js.Global().Get("VTTRegion")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *VTTRegion // javascript: VTTRegion _what_return_name
	)
	_converted = VTTRegionFromJS(_returned)
	_result = _converted
	return
}

// Id returning attribute 'id' with
// type string (idl: DOMString).
func (_this *VTTRegion) Id() string {
	var ret string
	value := _this.Value_JS.Get("id")
	ret = (value).String()
	return ret
}

// SetId setting attribute 'id' with
// type string (idl: DOMString).
func (_this *VTTRegion) SetId(value string) {
	input := value
	_this.Value_JS.Set("id", input)
}

// Width returning attribute 'width' with
// type float64 (idl: double).
func (_this *VTTRegion) Width() float64 {
	var ret float64
	value := _this.Value_JS.Get("width")
	ret = (value).Float()
	return ret
}

// SetWidth setting attribute 'width' with
// type float64 (idl: double).
func (_this *VTTRegion) SetWidth(value float64) {
	input := value
	_this.Value_JS.Set("width", input)
}

// Lines returning attribute 'lines' with
// type uint (idl: unsigned long).
func (_this *VTTRegion) Lines() uint {
	var ret uint
	value := _this.Value_JS.Get("lines")
	ret = (uint)((value).Int())
	return ret
}

// SetLines setting attribute 'lines' with
// type uint (idl: unsigned long).
func (_this *VTTRegion) SetLines(value uint) {
	input := value
	_this.Value_JS.Set("lines", input)
}

// RegionAnchorX returning attribute 'regionAnchorX' with
// type float64 (idl: double).
func (_this *VTTRegion) RegionAnchorX() float64 {
	var ret float64
	value := _this.Value_JS.Get("regionAnchorX")
	ret = (value).Float()
	return ret
}

// SetRegionAnchorX setting attribute 'regionAnchorX' with
// type float64 (idl: double).
func (_this *VTTRegion) SetRegionAnchorX(value float64) {
	input := value
	_this.Value_JS.Set("regionAnchorX", input)
}

// RegionAnchorY returning attribute 'regionAnchorY' with
// type float64 (idl: double).
func (_this *VTTRegion) RegionAnchorY() float64 {
	var ret float64
	value := _this.Value_JS.Get("regionAnchorY")
	ret = (value).Float()
	return ret
}

// SetRegionAnchorY setting attribute 'regionAnchorY' with
// type float64 (idl: double).
func (_this *VTTRegion) SetRegionAnchorY(value float64) {
	input := value
	_this.Value_JS.Set("regionAnchorY", input)
}

// ViewportAnchorX returning attribute 'viewportAnchorX' with
// type float64 (idl: double).
func (_this *VTTRegion) ViewportAnchorX() float64 {
	var ret float64
	value := _this.Value_JS.Get("viewportAnchorX")
	ret = (value).Float()
	return ret
}

// SetViewportAnchorX setting attribute 'viewportAnchorX' with
// type float64 (idl: double).
func (_this *VTTRegion) SetViewportAnchorX(value float64) {
	input := value
	_this.Value_JS.Set("viewportAnchorX", input)
}

// ViewportAnchorY returning attribute 'viewportAnchorY' with
// type float64 (idl: double).
func (_this *VTTRegion) ViewportAnchorY() float64 {
	var ret float64
	value := _this.Value_JS.Get("viewportAnchorY")
	ret = (value).Float()
	return ret
}

// SetViewportAnchorY setting attribute 'viewportAnchorY' with
// type float64 (idl: double).
func (_this *VTTRegion) SetViewportAnchorY(value float64) {
	input := value
	_this.Value_JS.Set("viewportAnchorY", input)
}

// Scroll returning attribute 'scroll' with
// type ScrollSetting (idl: ScrollSetting).
func (_this *VTTRegion) Scroll() ScrollSetting {
	var ret ScrollSetting
	value := _this.Value_JS.Get("scroll")
	ret = ScrollSettingFromJS(value)
	return ret
}

// SetScroll setting attribute 'scroll' with
// type ScrollSetting (idl: ScrollSetting).
func (_this *VTTRegion) SetScroll(value ScrollSetting) {
	input := value.JSValue()
	_this.Value_JS.Set("scroll", input)
}
