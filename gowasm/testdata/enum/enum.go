// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package enum

import js "github.com/gowebapi/webapi/core/js"

// using following types:

// source idl files:
// enum.idl

// transform files:
//

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

// enum: Foo
type Foo int

const (
	HelloFoo Foo = iota
	WorldFoo
)

var fooToWasmTable = []string{
	"hello", "world",
}

var fooFromWasmTable = map[string]Foo{
	"hello": HelloFoo, "world": WorldFoo,
}

// JSValue is converting this enum into a java object
func (this *Foo) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this Foo) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(fooToWasmTable) {
		return fooToWasmTable[idx]
	}
	panic("unknown input value")
}

// FooFromJS is converting a javascript value into
// a Foo enum value.
func FooFromJS(value js.Value) Foo {
	key := value.String()
	conv, ok := fooFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: Bar
type Bar int

const ()

var barToWasmTable = []string{}

var barFromWasmTable = map[string]Bar{}

// JSValue is converting this enum into a java object
func (this *Bar) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this Bar) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(barToWasmTable) {
		return barToWasmTable[idx]
	}
	panic("unknown input value")
}

// BarFromJS is converting a javascript value into
// a Bar enum value.
func BarFromJS(value js.Value) Bar {
	key := value.String()
	conv, ok := barFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// interface: Test
type Test struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Test) JSValue() js.Value {
	return _this.Value_JS
}

// TestFromJS is casting a js.Wrapper into Test.
func TestFromJS(value js.Wrapper) *Test {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Test{}
	ret.Value_JS = input
	return ret
}

// Hello1 returning attribute 'hello1' with
// type Foo (idl: Foo).
func (_this *Test) Hello1() Foo {
	var ret Foo
	value := _this.Value_JS.Get("hello1")
	ret = FooFromJS(value)
	return ret
}

// SetHello1 setting attribute 'hello1' with
// type Foo (idl: Foo).
func (_this *Test) SetHello1(value Foo) {
	input := value.JSValue()
	_this.Value_JS.Set("hello1", input)
}

// Hello2 returning attribute 'hello2' with
// type Bar (idl: Bar).
func (_this *Test) Hello2() Bar {
	var ret Bar
	value := _this.Value_JS.Get("hello2")
	ret = BarFromJS(value)
	return ret
}

// SetHello2 setting attribute 'hello2' with
// type Bar (idl: Bar).
func (_this *Test) SetHello2(value Bar) {
	input := value.JSValue()
	_this.Value_JS.Set("hello2", input)
}
