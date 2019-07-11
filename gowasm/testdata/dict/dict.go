// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package dict

import js "github.com/gowebapi/webapi/core/js"

// using following types:

// source idl files:
// dict.idl

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

// dictionary: Test1
type Test1 struct {
	A int
	B js.Value
	C []int
	D []js.Value
	E [][]int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Test1) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.A
	out.Set("a", value0)
	value1 := _this.B
	out.Set("b", value1)
	value2 := js.Global().Get("Array").New(len(_this.C))
	for __idx2, __seq_in2 := range _this.C {
		__seq_out2 := __seq_in2
		value2.SetIndex(__idx2, __seq_out2)
	}
	out.Set("c", value2)
	value3 := js.Global().Get("Array").New(len(_this.D))
	for __idx3, __seq_in3 := range _this.D {
		__seq_out3 := __seq_in3
		value3.SetIndex(__idx3, __seq_out3)
	}
	out.Set("d", value3)
	value4 := js.Global().Get("Array").New(len(_this.E))
	for __idx4, __seq_in4 := range _this.E {
		__seq_out4 := js.Global().Get("Array").New(len(__seq_in4))
		for __idx5, __seq_in5 := range __seq_in4 {
			__seq_out5 := __seq_in5
			__seq_out4.SetIndex(__idx5, __seq_out5)
		}
		value4.SetIndex(__idx4, __seq_out4)
	}
	out.Set("e", value4)
	return out
}

// Test1FromJS is allocating a new
// Test1 object and copy all values from
// input javascript object
func Test1FromJS(value js.Wrapper) *Test1 {
	input := value.JSValue()
	var out Test1
	var (
		value0 int        // javascript: long {a A a}
		value1 js.Value   // javascript: any {b B b}
		value2 []int      // javascript: sequence<long> {c C c}
		value3 []js.Value // javascript: sequence<any> {d D d}
		value4 [][]int    // javascript: sequence<sequence<long>> {e E e}
	)
	value0 = (input.Get("a")).Int()
	out.A = value0
	value1 = input.Get("b")
	out.B = value1
	__length2 := input.Get("c").Length()
	__array2 := make([]int, __length2, __length2)
	for __idx2 := 0; __idx2 < __length2; __idx2++ {
		var __seq_out2 int
		__seq_in2 := input.Get("c").Index(__idx2)
		__seq_out2 = (__seq_in2).Int()
		__array2[__idx2] = __seq_out2
	}
	value2 = __array2
	out.C = value2
	__length3 := input.Get("d").Length()
	__array3 := make([]js.Value, __length3, __length3)
	for __idx3 := 0; __idx3 < __length3; __idx3++ {
		var __seq_out3 js.Value
		__seq_in3 := input.Get("d").Index(__idx3)
		__seq_out3 = __seq_in3
		__array3[__idx3] = __seq_out3
	}
	value3 = __array3
	out.D = value3
	__length4 := input.Get("e").Length()
	__array4 := make([][]int, __length4, __length4)
	for __idx4 := 0; __idx4 < __length4; __idx4++ {
		var __seq_out4 []int
		__seq_in4 := input.Get("e").Index(__idx4)
		__length5 := __seq_in4.Length()
		__array5 := make([]int, __length5, __length5)
		for __idx5 := 0; __idx5 < __length5; __idx5++ {
			var __seq_out5 int
			__seq_in5 := __seq_in4.Index(__idx5)
			__seq_out5 = (__seq_in5).Int()
			__array5[__idx5] = __seq_out5
		}
		__seq_out4 = __array5
		__array4[__idx4] = __seq_out4
	}
	value4 = __array4
	out.E = value4
	return &out
}

// dictionary: Test2
type Test2 struct {
	E [][]int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Test2) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.E))
	for __idx0, __seq_in0 := range _this.E {
		__seq_out0 := js.Global().Get("Array").New(len(__seq_in0))
		for __idx1, __seq_in1 := range __seq_in0 {
			__seq_out1 := __seq_in1
			__seq_out0.SetIndex(__idx1, __seq_out1)
		}
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("e", value0)
	return out
}

// Test2FromJS is allocating a new
// Test2 object and copy all values from
// input javascript object
func Test2FromJS(value js.Wrapper) *Test2 {
	input := value.JSValue()
	var out Test2
	var (
		value0 [][]int // javascript: sequence<sequence<long>> {e E e}
	)
	__length0 := input.Get("e").Length()
	__array0 := make([][]int, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 []int
		__seq_in0 := input.Get("e").Index(__idx0)
		__length1 := __seq_in0.Length()
		__array1 := make([]int, __length1, __length1)
		for __idx1 := 0; __idx1 < __length1; __idx1++ {
			var __seq_out1 int
			__seq_in1 := __seq_in0.Index(__idx1)
			__seq_out1 = (__seq_in1).Int()
			__array1[__idx1] = __seq_out1
		}
		__seq_out0 = __array1
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.E = value0
	return &out
}

// dictionary: Empty
type Empty struct {
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Empty) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	return out
}

// EmptyFromJS is allocating a new
// Empty object and copy all values from
// input javascript object
func EmptyFromJS(value js.Wrapper) *Empty {
	var out Empty
	var ()
	return &out
}

// dictionary: Super
type Super struct {
	A int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Super) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.A
	out.Set("a", value0)
	return out
}

// SuperFromJS is allocating a new
// Super object and copy all values from
// input javascript object
func SuperFromJS(value js.Wrapper) *Super {
	input := value.JSValue()
	var out Super
	var (
		value0 int // javascript: long {a A a}
	)
	value0 = (input.Get("a")).Int()
	out.A = value0
	return &out
}

// dictionary: Inherit
type Inherit struct {
	A int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Inherit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.A
	out.Set("a", value0)
	return out
}

// InheritFromJS is allocating a new
// Inherit object and copy all values from
// input javascript object
func InheritFromJS(value js.Wrapper) *Inherit {
	input := value.JSValue()
	var out Inherit
	var (
		value0 int // javascript: long {a A a}
	)
	value0 = (input.Get("a")).Int()
	out.A = value0
	return &out
}

// interface: Foo
type Foo struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Foo) JSValue() js.Value {
	return _this.Value_JS
}

// FooFromJS is casting a js.Wrapper into Foo.
func FooFromJS(value js.Wrapper) *Foo {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Foo{}
	ret.Value_JS = input
	return ret
}

// Test1 returning attribute 'test1' with
// type Test1 (idl: Test1).
func (_this *Foo) Test1() *Test1 {
	var ret *Test1
	value := _this.Value_JS.Get("test1")
	ret = Test1FromJS(value)
	return ret
}

// SetTest1 setting attribute 'test1' with
// type Test1 (idl: Test1).
func (_this *Foo) SetTest1(value *Test1) {
	input := value.JSValue()
	_this.Value_JS.Set("test1", input)
}

// Test2 returning attribute 'test2' with
// type Test2 (idl: Test2).
func (_this *Foo) Test2() *Test2 {
	var ret *Test2
	value := _this.Value_JS.Get("test2")
	ret = Test2FromJS(value)
	return ret
}

// SetTest2 setting attribute 'test2' with
// type Test2 (idl: Test2).
func (_this *Foo) SetTest2(value *Test2) {
	input := value.JSValue()
	_this.Value_JS.Set("test2", input)
}

// Empty returning attribute 'empty' with
// type Empty (idl: Empty).
func (_this *Foo) Empty() *Empty {
	var ret *Empty
	value := _this.Value_JS.Get("empty")
	ret = EmptyFromJS(value)
	return ret
}

// SetEmpty setting attribute 'empty' with
// type Empty (idl: Empty).
func (_this *Foo) SetEmpty(value *Empty) {
	input := value.JSValue()
	_this.Value_JS.Set("empty", input)
}

// Test3 returning attribute 'test3' with
// type Inherit (idl: Inherit).
func (_this *Foo) Test3() *Inherit {
	var ret *Inherit
	value := _this.Value_JS.Get("test3")
	ret = InheritFromJS(value)
	return ret
}

// SetTest3 setting attribute 'test3' with
// type Inherit (idl: Inherit).
func (_this *Foo) SetTest3(value *Inherit) {
	input := value.JSValue()
	_this.Value_JS.Set("test3", input)
}
