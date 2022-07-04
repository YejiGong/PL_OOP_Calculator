package Calcul

import (
	"fmt"
	"reflect"
)

//stack
type Node struct {
	Value interface{}
}

type Stack struct {
	nodes []*Node
	count int
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

//cal
type Calc struct {
	Input Stack
}

func isZero(a interface{}) bool {
	tmp := reflect.ValueOf(a).MethodByName("ToString").Call([]reflect.Value{})[0].String()
	if tmp == "0" {
		return true
	} else {
		return false
	}
}

//for calculate, if type of two value is different -> convert type a to type b or convert type b to type a
func convertValue(c *Calc) (interface{}, interface{}) {
	a := c.Input.pop().Value
	b := c.Input.pop().Value
	v_a := reflect.ValueOf(a)
	v_b := reflect.ValueOf(b)
	t_a := v_a.Type()
	t_b := v_b.Type()
	var new_a, new_b interface{}
	if v_a.IsZero() == false && v_b.IsZero() == false {
		if t_a != t_b {
			temp_a := v_a.MethodByName("Convert").Call([]reflect.Value{v_b})
			if temp_a[0].Interface().(bool) == true {
				new_a = reflect.ValueOf(temp_a[1].Interface()).Interface()
				new_b = reflect.ValueOf(temp_a[2].Interface()).Interface()
			} else {
				temp_b := v_b.MethodByName("Convert").Call([]reflect.Value{v_a})
				new_b = reflect.ValueOf(temp_b[1].Interface()).Interface()
				new_a = reflect.ValueOf(temp_b[2].Interface()).Interface()
			}
		} else {
			new_a = v_a.Interface()
			new_b = v_b.Interface()
		}
	} else {
		new_a = v_a.Interface()
		new_b = v_b.Interface()
	}
	return new_a, new_b
}

//Calc->add, subtract, mutiply, divide method
func (c *Calc) Add() {
	a, b := convertValue(c)
	tmp_a := reflect.ValueOf(a).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface()
	tmp_b := reflect.ValueOf(b).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface()
	if isZero(tmp_a) == false && isZero(tmp_b) == false {
		tmp_c := reflect.ValueOf(tmp_a).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(tmp_a), reflect.ValueOf(tmp_b)})[0].Interface()
		c.Enter(tmp_c)
	} else if isZero(tmp_a) == true {
		c.Enter(tmp_b)
	} else {
		c.Enter(tmp_a)
	}
}

func (c *Calc) Sub() {
	a, b := convertValue(c)
	tmp_a := reflect.ValueOf(a).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface()
	tmp_b := reflect.ValueOf(b).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface()
	if isZero(tmp_a) == false && isZero(tmp_b) == false {
		tmp_c := reflect.ValueOf(tmp_a).MethodByName("Sub").Call([]reflect.Value{reflect.ValueOf(tmp_b), reflect.ValueOf(tmp_a)})[0].Interface()
		c.Enter(tmp_c)
	} else if isZero(tmp_a) == true {
		c.Enter(tmp_a)
	} else {
		tmp_c := reflect.ValueOf(tmp_a).MethodByName("Negative").Call([]reflect.Value{})[0].Interface()
		c.Enter(tmp_c)
	}
}

func (c *Calc) Multi() {
	a, b := convertValue(c)
	tmp_a := reflect.ValueOf(a).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface()
	tmp_b := reflect.ValueOf(b).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface()
	if isZero(tmp_a) == false && isZero(tmp_b) == false {
		tmp_c := reflect.ValueOf(tmp_a).MethodByName("Mult").Call([]reflect.Value{reflect.ValueOf(tmp_a), reflect.ValueOf(tmp_b)})[0].Interface()
		c.Enter(tmp_c)
	} else { //if tmp_b == 0
		c.Enter(NewInteger("0"))
	}
}

func (c *Calc) Divide() {
	a, b := convertValue(c)
	tmp_a := reflect.ValueOf(a).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface()
	tmp_b := reflect.ValueOf(b).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface()
	if isZero(tmp_a) == false && isZero(tmp_b) == false {
		tmp_c := reflect.ValueOf(tmp_a).MethodByName("Div").Call([]reflect.Value{reflect.ValueOf(tmp_b), reflect.ValueOf(tmp_a)})[0].Interface()
		c.Enter(tmp_c)
	} else if isZero(tmp_a) == true {
		panic("can't calculate")
	} else { //if tmp_b == 0
		c.Enter(NewInteger("0"))
	}
}

func (c *Calc) Enter(value interface{}) {
	c.Input.Push(&Node{value})
	fmt.Println(reflect.ValueOf(value).MethodByName("ToString").Call([]reflect.Value{})[0])
}

func (c *Calc) Clear() {
	for i := 0; i < c.Input.count; i++ {
		c.Input.pop()
	}
}
