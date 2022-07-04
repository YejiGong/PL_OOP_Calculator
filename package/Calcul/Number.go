package Calcul

import (
	"reflect"
	"strconv"
)

type Number interface {
	ToString() string //for print, make value to string
	Convert(interface{}) (bool, interface{}, interface{})
	//determining which type is appropriate when dealing parameter type and type of itself
	//conversion can be occur from parameter type to type of itself or from type of itself to parameter type
	//if it's impossible to convert, first return value will be false and sec,third return value will be zero value
	//for calculate, if a new type is defined, it have to be able to convert from all existing type.
	GetValue() interface{}                     //type of itself -> get real value
	SetValue(interface{})                      //value -> type of itself
	Sum(interface{}, interface{}) interface{}  //add two value of the type of itself
	Sub(interface{}, interface{}) interface{}  //subtract two value of the type of itself
	Mult(interface{}, interface{}) interface{} //multiply two value of the type of itself
	Div(interface{}, interface{}) interface{}  //dividing two valueof the type of itself
	Negative() interface{}                     //return negative value
}
type Integer struct {
	value int
}
type Float struct {
	value float64
}

func NewInteger(a string) Integer {
	v, _ := strconv.Atoi(a)
	var tmp Integer
	tmp.value = v
	return tmp
}
func (i Integer) SetValue(a int) Integer {
	tmp := strconv.Itoa(a)
	i = NewInteger(tmp)
	return i
}
func (i Integer) GetValue() Integer {
	tmp := strconv.Itoa(i.value)
	//value, _ := strconv.Atoi(tmp)
	return NewInteger(tmp)
}
func (i Integer) ToString() string {
	//return strconv.Itoa(i.Value())
	return strconv.Itoa(i.value)
}
func (i Integer) Convert(a interface{}) (bool, interface{}, interface{}) {
	var new_a, new_b interface{}
	var checker bool = false
	switch a.(type) {
	case Float:
		//i(Integer) -> Float, a-> Float.
		new_a = NewFloat(i.ToString())
		new_b = a
		checker = true
	}
	return checker, new_a, new_b
}
func (i Integer) Sum(a Integer, b Integer) Integer {
	var c Integer
	c.value = a.value + b.value
	return c
}
func (i Integer) Sub(a Integer, b Integer) Integer {
	var c Integer
	c.value = a.value - b.value
	return c
}
func (i Integer) Mult(a Integer, b Integer) Integer {
	var c Integer
	c.value = a.value * b.value
	return c
}
func (i Integer) Div(a Integer, b Integer) Integer {
	var c Integer
	if b.value != 0 {
		c.value = a.value / b.value
	} else {
		panic("can't calculate")
	}
	return c
}
func (i Integer) Negative() Integer {
	if i.value == 0 {
		return i
	} else {
		i.value = 0 - i.value
		return i
	}
}
func NewFloat(a string) Float {
	v, _ := strconv.ParseFloat(a, 64)
	var tmp Float
	tmp.value = v
	return tmp
}
func (f Float) SetValue(a float64) Float {
	tmp := strconv.FormatFloat(a, 'f', -1, 64)
	f = NewFloat(tmp)
	return f
}
func (f Float) GetValue() Float {
	tmp := f.ToString()
	//value, _ := strconv.ParseFloat(tmp, 64)
	return NewFloat(tmp)
}
func (f Float) ToString() string {
	return strconv.FormatFloat(f.value, 'f', -1, 64)
	//return string(f.value)
}
func (f Float) Convert(a interface{}) (bool, interface{}, interface{}) {
	var new_a, new_b interface{}
	var checker bool = false
	switch a.(type) {
	case Integer:
		tmp := reflect.ValueOf(a).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface().(Integer).ToString()
		new_a = NewFloat(tmp)
		new_b = f
		checker = true
	}
	return checker, new_a, new_b
}
func (f Float) Sum(a Float, b Float) Float {
	var c Float
	c.value = a.value + b.value
	return c
}
func (f Float) Sub(a Float, b Float) Float {
	var c Float
	c.value = a.value - b.value
	return c
}
func (f Float) Mult(a Float, b Float) Float {
	var c Float
	c.value = a.value * b.value
	return c
}
func (f Float) Div(a Float, b Float) Float {
	var c Float
	if b.value != 0 {
		c.value = a.value / b.value
	} else {
		panic("can't calculate")
	}
	return c
}
func (f Float) Negative() Float {
	if f.value == 0 {
		return f
	} else {
		f.value = 0 - f.value
		return f
	}
}
