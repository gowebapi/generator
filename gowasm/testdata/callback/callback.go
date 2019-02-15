// Code generated by webidlgenerator. DO NOT EDIT.

// +build !js

package callback

import js "github.com/gowebapi/webapi/core/failjs"

// using following types:

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

// enum: Baz
type Baz int

const (
	HelloBaz Baz = iota
	WorldBaz
)

var bazToWasmTable = []string{
	"hello", "world",
}

var bazFromWasmTable = map[string]Baz{
	"hello": HelloBaz, "world": WorldBaz,
}

// JSValue is converting this enum into a java object
func (this *Baz) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this Baz) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(bazToWasmTable) {
		return bazToWasmTable[idx]
	}
	panic("unknown input value")
}

// BazFromJS is converting a javascript value into
// a Baz enum value.
func BazFromJS(value js.Value) Baz {
	key := value.String()
	conv, ok := bazFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: Test1
type Test1 func()

func Test1ToJS(callback Test1) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var ()
		// TODO: return value
		callback()
	})
	return &ret
}

func Test1FromJS(_value js.Value) Test1 {
	return func() {
		var (
			_args [0]interface{}
			_end  int
		)
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: Test2
type Test2 func(a int, b string)

func Test2ToJS(callback Test2) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 int    // javascript: long a
			_p1 string // javascript: DOMString b
		)
		_p0 = (args[0]).Int()
		_p1 = (args[1]).String()
		// TODO: return value
		callback(_p0, _p1)
	})
	return &ret
}

func Test2FromJS(_value js.Value) Test2 {
	return func(a int, b string) {
		var (
			_args [2]interface{}
			_end  int
		)
		_p0 := a
		_args[0] = _p0
		_end++
		_p1 := b
		_args[1] = _p1
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: Test3
type Test3 func(a int, b string) int

func Test3ToJS(callback Test3) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 int    // javascript: long a
			_p1 string // javascript: DOMString b
		)
		_p0 = (args[0]).Int()
		_p1 = (args[1]).String()
		// TODO: return value
		callback(_p0, _p1)
	})
	return &ret
}

func Test3FromJS(_value js.Value) Test3 {
	return func(a int, b string) (_result int) {
		var (
			_args [2]interface{}
			_end  int
		)
		_p0 := a
		_args[0] = _p0
		_end++
		_p1 := b
		_args[1] = _p1
		_end++
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted int // javascript: long
		)
		_converted = (_returned).Int()
		_result = _converted
		return
	}
}

// callback: Test4
type Test4 func(a *Foo) *Foo

func Test4ToJS(callback Test4) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 *Foo // javascript: Foo a
		)
		_p0 = FooFromJS(args[0])
		// TODO: return value
		callback(_p0)
	})
	return &ret
}

func Test4FromJS(_value js.Value) Test4 {
	return func(a *Foo) (_result *Foo) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := a.JSValue()
		_args[0] = _p0
		_end++
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted *Foo // javascript: Foo
		)
		_converted = FooFromJS(_returned)
		_result = _converted
		return
	}
}

// callback: Test5
type Test5 func(a *Foo, b *Bar, c *Bar) *Bar

func Test5ToJS(callback Test5) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 *Foo // javascript: Foo a
			_p1 *Bar // javascript: Bar b
			_p2 *Bar // javascript: Bar c
		)
		_p0 = FooFromJS(args[0])
		_p1 = BarFromJS(args[1])
		_p2 = BarFromJS(args[2])
		// TODO: return value
		callback(_p0, _p1, _p2)
	})
	return &ret
}

func Test5FromJS(_value js.Value) Test5 {
	return func(a *Foo, b *Bar, c *Bar) (_result *Bar) {
		var (
			_args [3]interface{}
			_end  int
		)
		_p0 := a.JSValue()
		_args[0] = _p0
		_end++
		_p1 := b.JSValue()
		_args[1] = _p1
		_end++
		_p2 := c.JSValue()
		_args[2] = _p2
		_end++
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted *Bar // javascript: Bar
		)
		_converted = BarFromJS(_returned)
		_result = _converted
		return
	}
}

// callback: Test6
type Test6 func(a ...int)

func Test6ToJS(callback Test6) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 []int // javascript: long a
		)
		_p0 = make([]int, 0, len(args[0:]))
		for _, __in := range args[0:] {
			var __out int
			__out = (__in).Int()
			_p0 = append(_p0, __out)
		}
		// TODO: return value
		callback(_p0...)
	})
	return &ret
}

func Test6FromJS(_value js.Value) Test6 {
	return func(a ...int) {
		var (
			_args []interface{} = make([]interface{}, 0+len(a))
			_end  int
		)
		for _, __in := range a {
			__out := __in
			_args[_end] = __out
			_end++
		}
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: Test7
type Test7 func(b ...js.Value)

func Test7ToJS(callback Test7) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 []js.Value // javascript: any b
		)
		_p0 = make([]js.Value, 0, len(args[0:]))
		for _, __in := range args[0:] {
			var __out js.Value
			__out = __in
			_p0 = append(_p0, __out)
		}
		// TODO: return value
		callback(_p0...)
	})
	return &ret
}

func Test7FromJS(_value js.Value) Test7 {
	return func(b ...js.Value) {
		var (
			_args []interface{} = make([]interface{}, 0+len(b))
			_end  int
		)
		for _, __in := range b {
			__out := __in
			_args[_end] = __out
			_end++
		}
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: Test8
type Test8 func(a string, b string, c ...int) int

func Test8ToJS(callback Test8) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 string // javascript: DOMString a
			_p1 string // javascript: DOMString b
			_p2 []int  // javascript: long c
		)
		_p0 = (args[0]).String()
		_p1 = (args[1]).String()
		_p2 = make([]int, 0, len(args[2:]))
		for _, __in := range args[2:] {
			var __out int
			__out = (__in).Int()
			_p2 = append(_p2, __out)
		}
		// TODO: return value
		callback(_p0, _p1, _p2...)
	})
	return &ret
}

func Test8FromJS(_value js.Value) Test8 {
	return func(a string, b string, c ...int) (_result int) {
		var (
			_args []interface{} = make([]interface{}, 2+len(c))
			_end  int
		)
		_p0 := a
		_args[0] = _p0
		_end++
		_p1 := b
		_args[1] = _p1
		_end++
		for _, __in := range c {
			__out := __in
			_args[_end] = __out
			_end++
		}
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted int // javascript: long
		)
		_converted = (_returned).Int()
		_result = _converted
		return
	}
}

// callback: Test9
type Test9 func(a int, b ...*Bar) bool

func Test9ToJS(callback Test9) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 int    // javascript: long a
			_p1 []*Bar // javascript: Bar b
		)
		_p0 = (args[0]).Int()
		_p1 = make([]*Bar, 0, len(args[1:]))
		for _, __in := range args[1:] {
			var __out *Bar
			__out = BarFromJS(__in)
			_p1 = append(_p1, __out)
		}
		// TODO: return value
		callback(_p0, _p1...)
	})
	return &ret
}

func Test9FromJS(_value js.Value) Test9 {
	return func(a int, b ...*Bar) (_result bool) {
		var (
			_args []interface{} = make([]interface{}, 1+len(b))
			_end  int
		)
		_p0 := a
		_args[0] = _p0
		_end++
		for _, __in := range b {
			__out := __in.JSValue()
			_args[_end] = __out
			_end++
		}
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted bool // javascript: boolean
		)
		_converted = (_returned).Bool()
		_result = _converted
		return
	}
}

// callback: Test10
type Test10 func(a *Bar, c ...Baz) *Bar

func Test10ToJS(callback Test10) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 *Bar  // javascript: Bar a
			_p1 []Baz // javascript: Baz c
		)
		_p0 = BarFromJS(args[0])
		_p1 = make([]Baz, 0, len(args[1:]))
		for _, __in := range args[1:] {
			var __out Baz
			__out = BazFromJS(__in)
			_p1 = append(_p1, __out)
		}
		// TODO: return value
		callback(_p0, _p1...)
	})
	return &ret
}

func Test10FromJS(_value js.Value) Test10 {
	return func(a *Bar, c ...Baz) (_result *Bar) {
		var (
			_args []interface{} = make([]interface{}, 1+len(c))
			_end  int
		)
		_p0 := a.JSValue()
		_args[0] = _p0
		_end++
		for _, __in := range c {
			__out := __in.JSValue()
			_args[_end] = __out
			_end++
		}
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted *Bar // javascript: Bar
		)
		_converted = BarFromJS(_returned)
		_result = _converted
		return
	}
}

// callback: Test11
type Test11 func(a *Bar, c ...*Foo) string

func Test11ToJS(callback Test11) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 *Bar   // javascript: Bar a
			_p1 []*Foo // javascript: Foo c
		)
		_p0 = BarFromJS(args[0])
		_p1 = make([]*Foo, 0, len(args[1:]))
		for _, __in := range args[1:] {
			var __out *Foo
			__out = FooFromJS(__in)
			_p1 = append(_p1, __out)
		}
		// TODO: return value
		callback(_p0, _p1...)
	})
	return &ret
}

func Test11FromJS(_value js.Value) Test11 {
	return func(a *Bar, c ...*Foo) (_result string) {
		var (
			_args []interface{} = make([]interface{}, 1+len(c))
			_end  int
		)
		_p0 := a.JSValue()
		_args[0] = _p0
		_end++
		for _, __in := range c {
			__out := __in.JSValue()
			_args[_end] = __out
			_end++
		}
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted string // javascript: DOMString
		)
		_converted = (_returned).String()
		_result = _converted
		return
	}
}

// callback: Test12
type Test12 func(a *Bar, c ...*Union) *Bar

func Test12ToJS(callback Test12) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 *Bar     // javascript: Bar a
			_p1 []*Union // javascript: Union c
		)
		_p0 = BarFromJS(args[0])
		_p1 = make([]*Union, 0, len(args[1:]))
		for _, __in := range args[1:] {
			var __out *Union
			__out = UnionFromJS(__in)
			_p1 = append(_p1, __out)
		}
		// TODO: return value
		callback(_p0, _p1...)
	})
	return &ret
}

func Test12FromJS(_value js.Value) Test12 {
	return func(a *Bar, c ...*Union) (_result *Bar) {
		var (
			_args []interface{} = make([]interface{}, 1+len(c))
			_end  int
		)
		_p0 := a.JSValue()
		_args[0] = _p0
		_end++
		for _, __in := range c {
			__out := __in.JSValue()
			_args[_end] = __out
			_end++
		}
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted *Bar // javascript: Bar
		)
		_converted = BarFromJS(_returned)
		_result = _converted
		return
	}
}

// callback: Test13
type Test13 func(b ...js.Value) js.Value

func Test13ToJS(callback Test13) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 []js.Value // javascript: any b
		)
		_p0 = make([]js.Value, 0, len(args[0:]))
		for _, __in := range args[0:] {
			var __out js.Value
			__out = __in
			_p0 = append(_p0, __out)
		}
		// TODO: return value
		callback(_p0...)
	})
	return &ret
}

func Test13FromJS(_value js.Value) Test13 {
	return func(b ...js.Value) (_result js.Value) {
		var (
			_args []interface{} = make([]interface{}, 0+len(b))
			_end  int
		)
		for _, __in := range b {
			__out := __in
			_args[_end] = __out
			_end++
		}
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted js.Value // javascript: any
		)
		_converted = _returned
		_result = _converted
		return
	}
}

// callback: Test14
type Test14 func(b ...bool) *Union

func Test14ToJS(callback Test14) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 []bool // javascript: boolean b
		)
		_p0 = make([]bool, 0, len(args[0:]))
		for _, __in := range args[0:] {
			var __out bool
			__out = (__in).Bool()
			_p0 = append(_p0, __out)
		}
		// TODO: return value
		callback(_p0...)
	})
	return &ret
}

func Test14FromJS(_value js.Value) Test14 {
	return func(b ...bool) (_result *Union) {
		var (
			_args []interface{} = make([]interface{}, 0+len(b))
			_end  int
		)
		for _, __in := range b {
			__out := __in
			_args[_end] = __out
			_end++
		}
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted *Union // javascript: Union
		)
		_converted = UnionFromJS(_returned)
		_result = _converted
		return
	}
}

// callback: Test15
type Test15 func(c js.Value, d ...*Foo) *Foo

func Test15ToJS(callback Test15) *js.Callback {
	if callback == nil {
		return nil
	}
	ret := js.NewCallback(func(args []js.Value) {
		var (
			_p0 js.Value // javascript: any c
			_p1 []*Foo   // javascript: Foo d
		)
		_p0 = args[0]
		_p1 = make([]*Foo, 0, len(args[1:]))
		for _, __in := range args[1:] {
			var __out *Foo
			__out = FooFromJS(__in)
			_p1 = append(_p1, __out)
		}
		// TODO: return value
		callback(_p0, _p1...)
	})
	return &ret
}

func Test15FromJS(_value js.Value) Test15 {
	return func(c js.Value, d ...*Foo) (_result *Foo) {
		var (
			_args []interface{} = make([]interface{}, 1+len(d))
			_end  int
		)
		_p0 := c
		_args[0] = _p0
		_end++
		for _, __in := range d {
			__out := __in.JSValue()
			_args[_end] = __out
			_end++
		}
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted *Foo // javascript: Foo
		)
		_converted = FooFromJS(_returned)
		_result = _converted
		return
	}
}

// dictionary: Bar
type Bar struct {
	A int
	B int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Bar) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.A
	out.Set("a", value0)
	value1 := _this.B
	out.Set("b", value1)
	return out
}

// BarFromJS is allocating a new
// Bar object and copy all values from
// input javascript object
func BarFromJS(input js.Value) *Bar {
	var out Bar
	var (
		out0 int // javascript: long {a A a}
		out1 int // javascript: long {b B b}
	)
	out0 = (input.Get("a")).Int()
	out.A = out0
	out1 = (input.Get("b")).Int()
	out.B = out1
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

// FooFromJS is casting a js.Value into Foo.
func FooFromJS(input js.Value) *Foo {
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Foo{}
	ret.Value_JS = input
	return ret
}

// Test1 returning attribute 'test1' with
// type Test1 (idl: Test1).
func (_this *Foo) Test1() Test1 {
	var ret Test1
	value := _this.Value_JS.Get("test1")
	ret = Test1FromJS(value)
	return ret
}

// SetTest1 setting attribute 'test1' with
// type Test1 (idl: Test1).
func (_this *Foo) SetTest1(value *js.Callback) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("test1", input)
}

// Test2 returning attribute 'test2' with
// type Test2 (idl: Test2).
func (_this *Foo) Test2() Test2 {
	var ret Test2
	value := _this.Value_JS.Get("test2")
	ret = Test2FromJS(value)
	return ret
}

// SetTest2 setting attribute 'test2' with
// type Test2 (idl: Test2).
func (_this *Foo) SetTest2(value *js.Callback) {
	var __callback1 js.Value
	if value != nil {
		__callback1 = (*value).Value
	} else {
		__callback1 = js.Null()
	}
	input := __callback1
	_this.Value_JS.Set("test2", input)
}

// Test3 returning attribute 'test3' with
// type Test3 (idl: Test3).
func (_this *Foo) Test3() Test3 {
	var ret Test3
	value := _this.Value_JS.Get("test3")
	ret = Test3FromJS(value)
	return ret
}

// SetTest3 setting attribute 'test3' with
// type Test3 (idl: Test3).
func (_this *Foo) SetTest3(value *js.Callback) {
	var __callback2 js.Value
	if value != nil {
		__callback2 = (*value).Value
	} else {
		__callback2 = js.Null()
	}
	input := __callback2
	_this.Value_JS.Set("test3", input)
}

// Test4 returning attribute 'test4' with
// type Test4 (idl: Test4).
func (_this *Foo) Test4() Test4 {
	var ret Test4
	value := _this.Value_JS.Get("test4")
	ret = Test4FromJS(value)
	return ret
}

// SetTest4 setting attribute 'test4' with
// type Test4 (idl: Test4).
func (_this *Foo) SetTest4(value *js.Callback) {
	var __callback3 js.Value
	if value != nil {
		__callback3 = (*value).Value
	} else {
		__callback3 = js.Null()
	}
	input := __callback3
	_this.Value_JS.Set("test4", input)
}

// Test5 returning attribute 'test5' with
// type Test5 (idl: Test5).
func (_this *Foo) Test5() Test5 {
	var ret Test5
	value := _this.Value_JS.Get("test5")
	ret = Test5FromJS(value)
	return ret
}

// SetTest5 setting attribute 'test5' with
// type Test5 (idl: Test5).
func (_this *Foo) SetTest5(value *js.Callback) {
	var __callback4 js.Value
	if value != nil {
		__callback4 = (*value).Value
	} else {
		__callback4 = js.Null()
	}
	input := __callback4
	_this.Value_JS.Set("test5", input)
}

// Test6 returning attribute 'test6' with
// type Test6 (idl: Test6).
func (_this *Foo) Test6() Test6 {
	var ret Test6
	value := _this.Value_JS.Get("test6")
	ret = Test6FromJS(value)
	return ret
}

// SetTest6 setting attribute 'test6' with
// type Test6 (idl: Test6).
func (_this *Foo) SetTest6(value *js.Callback) {
	var __callback5 js.Value
	if value != nil {
		__callback5 = (*value).Value
	} else {
		__callback5 = js.Null()
	}
	input := __callback5
	_this.Value_JS.Set("test6", input)
}

// Test7 returning attribute 'test7' with
// type Test7 (idl: Test7).
func (_this *Foo) Test7() Test7 {
	var ret Test7
	value := _this.Value_JS.Get("test7")
	ret = Test7FromJS(value)
	return ret
}

// SetTest7 setting attribute 'test7' with
// type Test7 (idl: Test7).
func (_this *Foo) SetTest7(value *js.Callback) {
	var __callback6 js.Value
	if value != nil {
		__callback6 = (*value).Value
	} else {
		__callback6 = js.Null()
	}
	input := __callback6
	_this.Value_JS.Set("test7", input)
}

// Test8 returning attribute 'test8' with
// type Test8 (idl: Test8).
func (_this *Foo) Test8() Test8 {
	var ret Test8
	value := _this.Value_JS.Get("test8")
	ret = Test8FromJS(value)
	return ret
}

// SetTest8 setting attribute 'test8' with
// type Test8 (idl: Test8).
func (_this *Foo) SetTest8(value *js.Callback) {
	var __callback7 js.Value
	if value != nil {
		__callback7 = (*value).Value
	} else {
		__callback7 = js.Null()
	}
	input := __callback7
	_this.Value_JS.Set("test8", input)
}

// Test9 returning attribute 'test9' with
// type Test9 (idl: Test9).
func (_this *Foo) Test9() Test9 {
	var ret Test9
	value := _this.Value_JS.Get("test9")
	ret = Test9FromJS(value)
	return ret
}

// SetTest9 setting attribute 'test9' with
// type Test9 (idl: Test9).
func (_this *Foo) SetTest9(value *js.Callback) {
	var __callback8 js.Value
	if value != nil {
		__callback8 = (*value).Value
	} else {
		__callback8 = js.Null()
	}
	input := __callback8
	_this.Value_JS.Set("test9", input)
}

// Test10 returning attribute 'test10' with
// type Test10 (idl: Test10).
func (_this *Foo) Test10() Test10 {
	var ret Test10
	value := _this.Value_JS.Get("test10")
	ret = Test10FromJS(value)
	return ret
}

// SetTest10 setting attribute 'test10' with
// type Test10 (idl: Test10).
func (_this *Foo) SetTest10(value *js.Callback) {
	var __callback9 js.Value
	if value != nil {
		__callback9 = (*value).Value
	} else {
		__callback9 = js.Null()
	}
	input := __callback9
	_this.Value_JS.Set("test10", input)
}

// Test11 returning attribute 'test11' with
// type Test11 (idl: Test11).
func (_this *Foo) Test11() Test11 {
	var ret Test11
	value := _this.Value_JS.Get("test11")
	ret = Test11FromJS(value)
	return ret
}

// SetTest11 setting attribute 'test11' with
// type Test11 (idl: Test11).
func (_this *Foo) SetTest11(value *js.Callback) {
	var __callback10 js.Value
	if value != nil {
		__callback10 = (*value).Value
	} else {
		__callback10 = js.Null()
	}
	input := __callback10
	_this.Value_JS.Set("test11", input)
}

// Test12 returning attribute 'test12' with
// type Test12 (idl: Test12).
func (_this *Foo) Test12() Test12 {
	var ret Test12
	value := _this.Value_JS.Get("test12")
	ret = Test12FromJS(value)
	return ret
}

// SetTest12 setting attribute 'test12' with
// type Test12 (idl: Test12).
func (_this *Foo) SetTest12(value *js.Callback) {
	var __callback11 js.Value
	if value != nil {
		__callback11 = (*value).Value
	} else {
		__callback11 = js.Null()
	}
	input := __callback11
	_this.Value_JS.Set("test12", input)
}

// Test13 returning attribute 'test13' with
// type Test13 (idl: Test13).
func (_this *Foo) Test13() Test13 {
	var ret Test13
	value := _this.Value_JS.Get("test13")
	ret = Test13FromJS(value)
	return ret
}

// SetTest13 setting attribute 'test13' with
// type Test13 (idl: Test13).
func (_this *Foo) SetTest13(value *js.Callback) {
	var __callback12 js.Value
	if value != nil {
		__callback12 = (*value).Value
	} else {
		__callback12 = js.Null()
	}
	input := __callback12
	_this.Value_JS.Set("test13", input)
}

// Test14 returning attribute 'test14' with
// type Test14 (idl: Test14).
func (_this *Foo) Test14() Test14 {
	var ret Test14
	value := _this.Value_JS.Get("test14")
	ret = Test14FromJS(value)
	return ret
}

// SetTest14 setting attribute 'test14' with
// type Test14 (idl: Test14).
func (_this *Foo) SetTest14(value *js.Callback) {
	var __callback13 js.Value
	if value != nil {
		__callback13 = (*value).Value
	} else {
		__callback13 = js.Null()
	}
	input := __callback13
	_this.Value_JS.Set("test14", input)
}

// Test15 returning attribute 'test15' with
// type Test15 (idl: Test15).
func (_this *Foo) Test15() Test15 {
	var ret Test15
	value := _this.Value_JS.Get("test15")
	ret = Test15FromJS(value)
	return ret
}

// SetTest15 setting attribute 'test15' with
// type Test15 (idl: Test15).
func (_this *Foo) SetTest15(value *js.Callback) {
	var __callback14 js.Value
	if value != nil {
		__callback14 = (*value).Value
	} else {
		__callback14 = js.Null()
	}
	input := __callback14
	_this.Value_JS.Set("test15", input)
}
