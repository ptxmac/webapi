// Code generated by webidl-bind. DO NOT EDIT.

package parser

import "syscall/js"

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/core"
	"github.com/gowebapi/webapi/dom"
)

// using following types:
// dom.Node
// webapi.Document

// source idl files:
// DOM-Parsing.idl

// transform files:
// DOM-Parsing.go.md

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

// enum: SupportedType
type SupportedType int

const (
	TextHtmlSupportedType SupportedType = iota
	TextXmlSupportedType
	ApplicationXmlSupportedType
	ApplicationXhtmlXmlSupportedType
	ImageSvgXmlSupportedType
)

var supportedTypeToWasmTable = []string{
	"text/html", "text/xml", "application/xml", "application/xhtml+xml", "image/svg+xml",
}

var supportedTypeFromWasmTable = map[string]SupportedType{
	"text/html": TextHtmlSupportedType, "text/xml": TextXmlSupportedType, "application/xml": ApplicationXmlSupportedType, "application/xhtml+xml": ApplicationXhtmlXmlSupportedType, "image/svg+xml": ImageSvgXmlSupportedType,
}

// JSValue is converting this enum into a javascript object
func (this *SupportedType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this SupportedType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(supportedTypeToWasmTable) {
		return supportedTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// SupportedTypeFromJS is converting a javascript value into
// a SupportedType enum value.
func SupportedTypeFromJS(value js.Value) SupportedType {
	key := value.String()
	conv, ok := supportedTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// class: DOMParser
type DOMParser struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *DOMParser) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// DOMParserFromJS is casting a js.Value into DOMParser.
func DOMParserFromJS(value js.Value) *DOMParser {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &DOMParser{}
	ret.Value_JS = value
	return ret
}

// DOMParserFromJS is casting from something that holds a js.Value into DOMParser.
func DOMParserFromWrapper(input core.Wrapper) *DOMParser {
	return DOMParserFromJS(input.JSValue())
}

func NewDOMParser() (_result *DOMParser) {
	_klass := js.Global().Get("DOMParser")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *DOMParser // javascript: DOMParser _what_return_name
	)
	_converted = DOMParserFromJS(_returned)
	_result = _converted
	return
}

func (_this *DOMParser) ParseFromString(str string, _type SupportedType) (_result *webapi.Document) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := str
	_args[0] = _p0
	_end++
	_p1 := _type.JSValue()
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("parseFromString", _args[0:_end]...)
	var (
		_converted *webapi.Document // javascript: Document _what_return_name
	)
	_converted = webapi.DocumentFromJS(_returned)
	_result = _converted
	return
}

// class: XMLSerializer
type XMLSerializer struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *XMLSerializer) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// XMLSerializerFromJS is casting a js.Value into XMLSerializer.
func XMLSerializerFromJS(value js.Value) *XMLSerializer {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &XMLSerializer{}
	ret.Value_JS = value
	return ret
}

// XMLSerializerFromJS is casting from something that holds a js.Value into XMLSerializer.
func XMLSerializerFromWrapper(input core.Wrapper) *XMLSerializer {
	return XMLSerializerFromJS(input.JSValue())
}

func NewXMLSerializer() (_result *XMLSerializer) {
	_klass := js.Global().Get("XMLSerializer")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *XMLSerializer // javascript: XMLSerializer _what_return_name
	)
	_converted = XMLSerializerFromJS(_returned)
	_result = _converted
	return
}

func (_this *XMLSerializer) SerializeToString(root *dom.Node) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := root.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("serializeToString", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}
