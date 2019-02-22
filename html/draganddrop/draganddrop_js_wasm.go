// Code generated by webidl-bind. DO NOT EDIT.

package draganddrop

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom"
	"github.com/gowebapi/webapi/fileapi"
	"github.com/gowebapi/webapi/javascript"
	"github.com/gowebapi/webapi/patch"
)

// using following types:
// dom.Element
// fileapi.File
// fileapi.FileList
// javascript.FrozenArray
// patch.MouseEvent

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

// callback: FunctionStringCallback
type FunctionStringCallbackFunc func(data string)

// FunctionStringCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type FunctionStringCallback js.Func

func FunctionStringCallbackToJS(callback FunctionStringCallbackFunc) *FunctionStringCallback {
	if callback == nil {
		return nil
	}
	ret := FunctionStringCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 string // javascript: DOMString data
		)
		_p0 = (args[0]).String()
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func FunctionStringCallbackFromJS(_value js.Value) FunctionStringCallbackFunc {
	return func(data string) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := data
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: DragEventInit
type DragEventInit struct {
	DataTransfer *DataTransfer
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *DragEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.DataTransfer.JSValue()
	out.Set("dataTransfer", value0)
	return out
}

// DragEventInitFromJS is allocating a new
// DragEventInit object and copy all values from
// input javascript object
func DragEventInitFromJS(value js.Wrapper) *DragEventInit {
	input := value.JSValue()
	var out DragEventInit
	var (
		value0 *DataTransfer // javascript: DataTransfer {dataTransfer DataTransfer dataTransfer}
	)
	if input.Get("dataTransfer").Type() != js.TypeNull {
		value0 = DataTransferFromJS(input.Get("dataTransfer"))
	}
	out.DataTransfer = value0
	return &out
}

// interface: DataTransfer
type DataTransfer struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DataTransfer) JSValue() js.Value {
	return _this.Value_JS
}

// DataTransferFromJS is casting a js.Wrapper into DataTransfer.
func DataTransferFromJS(value js.Wrapper) *DataTransfer {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DataTransfer{}
	ret.Value_JS = input
	return ret
}

func NewDataTransfer() (_result *DataTransfer) {
	_klass := js.Global().Get("DataTransfer")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *DataTransfer // javascript: DataTransfer _what_return_name
	)
	_converted = DataTransferFromJS(_returned)
	_result = _converted
	return
}

// DropEffect returning attribute 'dropEffect' with
// type string (idl: DOMString).
func (_this *DataTransfer) DropEffect() string {
	var ret string
	value := _this.Value_JS.Get("dropEffect")
	ret = (value).String()
	return ret
}

// SetDropEffect setting attribute 'dropEffect' with
// type string (idl: DOMString).
func (_this *DataTransfer) SetDropEffect(value string) {
	input := value
	_this.Value_JS.Set("dropEffect", input)
}

// EffectAllowed returning attribute 'effectAllowed' with
// type string (idl: DOMString).
func (_this *DataTransfer) EffectAllowed() string {
	var ret string
	value := _this.Value_JS.Get("effectAllowed")
	ret = (value).String()
	return ret
}

// SetEffectAllowed setting attribute 'effectAllowed' with
// type string (idl: DOMString).
func (_this *DataTransfer) SetEffectAllowed(value string) {
	input := value
	_this.Value_JS.Set("effectAllowed", input)
}

// Items returning attribute 'items' with
// type DataTransferItemList (idl: DataTransferItemList).
func (_this *DataTransfer) Items() *DataTransferItemList {
	var ret *DataTransferItemList
	value := _this.Value_JS.Get("items")
	ret = DataTransferItemListFromJS(value)
	return ret
}

// Types returning attribute 'types' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *DataTransfer) Types() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("types")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// Files returning attribute 'files' with
// type fileapi.FileList (idl: FileList).
func (_this *DataTransfer) Files() *fileapi.FileList {
	var ret *fileapi.FileList
	value := _this.Value_JS.Get("files")
	ret = fileapi.FileListFromJS(value)
	return ret
}

func (_this *DataTransfer) SetDragImage(image *dom.Element, x int, y int) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := image.JSValue()
	_args[0] = _p0
	_end++
	_p1 := x
	_args[1] = _p1
	_end++
	_p2 := y
	_args[2] = _p2
	_end++
	_this.Value_JS.Call("setDragImage", _args[0:_end]...)
	return
}

func (_this *DataTransfer) GetData(format string) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := format
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getData", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *DataTransfer) SetData(format string, data string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := format
	_args[0] = _p0
	_end++
	_p1 := data
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("setData", _args[0:_end]...)
	return
}

func (_this *DataTransfer) ClearData(format *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	if format != nil {
		_p0 := format
		_args[0] = _p0
		_end++
	}
	_this.Value_JS.Call("clearData", _args[0:_end]...)
	return
}

// interface: DataTransferItem
type DataTransferItem struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DataTransferItem) JSValue() js.Value {
	return _this.Value_JS
}

// DataTransferItemFromJS is casting a js.Wrapper into DataTransferItem.
func DataTransferItemFromJS(value js.Wrapper) *DataTransferItem {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DataTransferItem{}
	ret.Value_JS = input
	return ret
}

// Kind returning attribute 'kind' with
// type string (idl: DOMString).
func (_this *DataTransferItem) Kind() string {
	var ret string
	value := _this.Value_JS.Get("kind")
	ret = (value).String()
	return ret
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *DataTransferItem) Type() string {
	var ret string
	value := _this.Value_JS.Get("type")
	ret = (value).String()
	return ret
}

func (_this *DataTransferItem) GetAsString(callback *FunctionStringCallback) {
	var (
		_args [1]interface{}
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
	_this.Value_JS.Call("getAsString", _args[0:_end]...)
	return
}

func (_this *DataTransferItem) GetAsFile() (_result *fileapi.File) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getAsFile", _args[0:_end]...)
	var (
		_converted *fileapi.File // javascript: File _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = fileapi.FileFromJS(_returned)
	}
	_result = _converted
	return
}

// interface: DataTransferItemList
type DataTransferItemList struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *DataTransferItemList) JSValue() js.Value {
	return _this.Value_JS
}

// DataTransferItemListFromJS is casting a js.Wrapper into DataTransferItemList.
func DataTransferItemListFromJS(value js.Wrapper) *DataTransferItemList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DataTransferItemList{}
	ret.Value_JS = input
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *DataTransferItemList) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *DataTransferItemList) Add(data string, _type string) (_result *DataTransferItem) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := data
	_args[0] = _p0
	_end++
	_p1 := _type
	_args[1] = _p1
	_end++
	_returned := _this.Value_JS.Call("add", _args[0:_end]...)
	var (
		_converted *DataTransferItem // javascript: DataTransferItem _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = DataTransferItemFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *DataTransferItemList) Add2(data *fileapi.File) (_result *DataTransferItem) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("add", _args[0:_end]...)
	var (
		_converted *DataTransferItem // javascript: DataTransferItem _what_return_name
	)
	if _returned.Type() != js.TypeNull {
		_converted = DataTransferItemFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *DataTransferItemList) Remove(index uint) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("remove", _args[0:_end]...)
	return
}

func (_this *DataTransferItemList) Clear() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("clear", _args[0:_end]...)
	return
}

// interface: DragEvent
type DragEvent struct {
	patch.MouseEvent
}

// DragEventFromJS is casting a js.Wrapper into DragEvent.
func DragEventFromJS(value js.Wrapper) *DragEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &DragEvent{}
	ret.Value_JS = input
	return ret
}

func NewDragEvent(_type string, eventInitDict *DragEventInit) (_result *DragEvent) {
	_klass := js.Global().Get("DragEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *DragEvent // javascript: DragEvent _what_return_name
	)
	_converted = DragEventFromJS(_returned)
	_result = _converted
	return
}

// DataTransfer returning attribute 'dataTransfer' with
// type DataTransfer (idl: DataTransfer).
func (_this *DragEvent) DataTransfer() *DataTransfer {
	var ret *DataTransfer
	value := _this.Value_JS.Get("dataTransfer")
	if value.Type() != js.TypeNull {
		ret = DataTransferFromJS(value)
	}
	return ret
}
