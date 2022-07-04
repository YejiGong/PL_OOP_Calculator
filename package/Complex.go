package main

import (
	"Calcul"
	"reflect"
	"strings"
)

type Complex struct {
	RealNum      interface{}
	ImaginaryNum interface{}
	Calcul.Number
}

//a+bi 꼴
func NewComplex(a string) interface{} {
	var tmp Complex
	var rStr []string
	if strings.Contains(a, "+") {
		rStr = strings.Split(a, "+")
	} else if strings.Contains(a, "-") {
		rStr = strings.Split(a, "-")
		if rStr[0] != "" {
			rStr = strings.Split(a, "-")
			rStr[1] = "-" + rStr[1]
		} else {
			rStr[0] = "-" + rStr[1]
			rStr[1] = "-" + rStr[2]
		}
	}
	if strings.Contains(rStr[0], "/") {
		tmp.RealNum = NewRational(rStr[0])
	} else if strings.Contains(rStr[0], ".") {
		tmp.RealNum = Calcul.NewFloat(rStr[0])
	} else {
		tmp.RealNum = Calcul.NewInteger(rStr[0])
	}
	rStr_ := strings.Split(rStr[1], "i")
	if strings.Contains(rStr_[0], "/") {
		tmp.ImaginaryNum = NewRational(rStr_[0])
	} else if strings.Contains(rStr_[0], ".") {
		tmp.ImaginaryNum = Calcul.NewFloat(rStr_[0])
	} else {
		if rStr_[0] == "0" {
			return tmp.RealNum
		}
		tmp.ImaginaryNum = Calcul.NewInteger(rStr_[0])
	}
	return tmp
}
func (c Complex) GetValue() Complex {
	return c
}
func (c Complex) SetValue(a Complex) Complex {
	return a
}
func (c Complex) ToString() string {
	strRealNum := reflect.ValueOf(c.RealNum).MethodByName("ToString").Call([]reflect.Value{})[0].Interface().(string)
	strImaginaryNum := reflect.ValueOf(c.ImaginaryNum).MethodByName("ToString").Call([]reflect.Value{})[0].Interface().(string)
	if strings.Contains(strImaginaryNum, "-") {
		return strRealNum + strImaginaryNum + "i"
	} else {
		return strRealNum + "+" + strImaginaryNum + "i"
	}
}
func (c Complex) Convert(a interface{}) (bool, interface{}, interface{}) {
	var new_a, new_b interface{}
	var tmp Complex
	var checker bool
	switch a.(type) {
	case Calcul.Integer:
		tmp.RealNum = reflect.ValueOf(a).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface().(Calcul.Integer)
		tmp.ImaginaryNum = Calcul.NewInteger("0")
	case Calcul.Float:
		tmp.RealNum = reflect.ValueOf(a).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface().(Calcul.Float)
		tmp.ImaginaryNum = Calcul.NewInteger("0")
	case Rational:
		tmp.RealNum = reflect.ValueOf(a).MethodByName("GetValue").Call([]reflect.Value{})[0].Interface().(Rational)
		tmp.ImaginaryNum = Calcul.NewInteger("0")
	}
	new_b = reflect.ValueOf(tmp).Interface()
	new_a = c
	return checker, new_a, new_b
}

//using while calculating (because RealNum, IdealNum types can be different. it's not definite)
func conversion(a interface{}, b interface{}) (interface{}, interface{}) {
	var new_a, new_b interface{}
	v_a := reflect.ValueOf(a)
	v_b := reflect.ValueOf(b)
	temp_a := v_a.MethodByName("Convert").Call([]reflect.Value{v_b})
	if temp_a[0].Interface().(bool) == true {
		new_a = reflect.ValueOf(temp_a[1].Interface()).Interface()
		new_b = reflect.ValueOf(temp_a[2].Interface()).Interface()
	} else {
		temp_b := v_b.MethodByName("Convert").Call([]reflect.Value{v_a})
		new_b = reflect.ValueOf(temp_b[1].Interface()).Interface()
		new_a = reflect.ValueOf(temp_b[2].Interface()).Interface()
	}
	return new_a, new_b
}
func (c Complex) Sum(a Complex, b Complex) interface{} {
	var tmp Complex
	if reflect.TypeOf(a.RealNum) != reflect.TypeOf(b.RealNum) {
		NewRealA, NewRealB := conversion(a.RealNum, b.RealNum)
		tmp.RealNum = reflect.ValueOf(NewRealA).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(NewRealA), reflect.ValueOf(NewRealB)})[0].Interface()
	} else {
		tmp.RealNum = reflect.ValueOf(a.RealNum).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(a.RealNum), reflect.ValueOf(b.RealNum)})[0].Interface()
	}
	if reflect.TypeOf(a.ImaginaryNum) != reflect.TypeOf(b.ImaginaryNum) {
		NewIdealA, NewIdealB := conversion(a.ImaginaryNum, b.ImaginaryNum)
		tmp.ImaginaryNum = reflect.ValueOf(NewIdealA).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(NewIdealA), reflect.ValueOf(NewIdealB)})[0].Interface()
	} else {
		tmp.ImaginaryNum = reflect.ValueOf(a.ImaginaryNum).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(a.ImaginaryNum), reflect.ValueOf(b.ImaginaryNum)})[0].Interface()
	}
	//if ImaginaryNum=0 -> return only realnum
	if reflect.ValueOf(tmp.ImaginaryNum).MethodByName("ToString").Call([]reflect.Value{})[0].Interface().(string) == "0" {
		return tmp.RealNum
	} else {
		return tmp
	}
}
func (c Complex) Sub(a Complex, b Complex) interface{} {
	var tmp Complex
	if reflect.TypeOf(a.RealNum) != reflect.TypeOf(b.RealNum) {
		NewRealA, NewRealB := conversion(a.RealNum, b.RealNum)
		tmp.RealNum = reflect.ValueOf(NewRealA).MethodByName("Sub").Call([]reflect.Value{reflect.ValueOf(NewRealA), reflect.ValueOf(NewRealB)})[0].Interface()
	} else {
		tmp.RealNum = reflect.ValueOf(a.RealNum).MethodByName("Sub").Call([]reflect.Value{reflect.ValueOf(a.RealNum), reflect.ValueOf(b.RealNum)})[0].Interface()
	}
	if reflect.TypeOf(a.ImaginaryNum) != reflect.TypeOf(b.ImaginaryNum) {
		NewIdealA, NewIdealB := conversion(a.ImaginaryNum, b.ImaginaryNum)
		tmp.ImaginaryNum = reflect.ValueOf(NewIdealA).MethodByName("Sub").Call([]reflect.Value{reflect.ValueOf(NewIdealA), reflect.ValueOf(NewIdealB)})[0].Interface()
	} else {
		tmp.ImaginaryNum = reflect.ValueOf(a.ImaginaryNum).MethodByName("Sub").Call([]reflect.Value{reflect.ValueOf(a.ImaginaryNum), reflect.ValueOf(b.ImaginaryNum)})[0].Interface()
	}
	//if ImaginaryNum=0 -> return only realnum
	if reflect.ValueOf(tmp.ImaginaryNum).MethodByName("ToString").Call([]reflect.Value{})[0].Interface().(string) == "0" {
		return tmp.RealNum
	} else {
		return tmp
	}
}

//for complex multiple, dividing
func multiple(a interface{}, b interface{}, c interface{}, d interface{}) (interface{}, interface{}, interface{}, interface{}) {
	var i, j, k, l interface{}
	//i=ac, j=bd, k=ad, l=bc
	if reflect.TypeOf(a) != reflect.TypeOf(c) {
		NewA, NewB := conversion(a, c)
		i = reflect.ValueOf(NewA).MethodByName("Mult").Call([]reflect.Value{reflect.ValueOf(NewA), reflect.ValueOf(NewB)})[0].Interface()
	} else {
		i = reflect.ValueOf(a).MethodByName("Mult").Call([]reflect.Value{reflect.ValueOf(a), reflect.ValueOf(c)})[0].Interface()
	}
	if reflect.TypeOf(b) != reflect.TypeOf(d) {
		NewA, NewB := conversion(b, d)
		j = reflect.ValueOf(NewA).MethodByName("Mult").Call([]reflect.Value{reflect.ValueOf(NewA), reflect.ValueOf(NewB)})[0].Interface()
	} else {
		j = reflect.ValueOf(b).MethodByName("Mult").Call([]reflect.Value{reflect.ValueOf(b), reflect.ValueOf(d)})[0].Interface()
	}

	if reflect.TypeOf(a) != reflect.TypeOf(d) {
		NewA, NewB := conversion(a, d)
		k = reflect.ValueOf(NewA).MethodByName("Mult").Call([]reflect.Value{reflect.ValueOf(NewA), reflect.ValueOf(NewB)})[0].Interface()
	} else {
		k = reflect.ValueOf(a).MethodByName("Mult").Call([]reflect.Value{reflect.ValueOf(a), reflect.ValueOf(d)})[0].Interface()
	}
	if reflect.TypeOf(b) != reflect.TypeOf(c) {
		NewA, NewB := conversion(b, c)
		l = reflect.ValueOf(NewA).MethodByName("Mult").Call([]reflect.Value{reflect.ValueOf(NewA), reflect.ValueOf(NewB)})[0].Interface()
	} else {
		l = reflect.ValueOf(b).MethodByName("Mult").Call([]reflect.Value{reflect.ValueOf(b), reflect.ValueOf(c)})[0].Interface()
	}
	return i, j, k, l
}
func (c Complex) Mult(a Complex, b Complex) interface{} {
	//(a+bi)*(c+di)=(ac-bd)+(ad+bc)i
	i, j, k, l := multiple(a.RealNum, a.ImaginaryNum, b.RealNum, b.ImaginaryNum)
	//ac = i, bd = j, ad =k, bc= l
	var tmp Complex
	if reflect.TypeOf(i) != reflect.TypeOf(j) {
		NewI, NewJ := conversion(i, j)
		tmp.RealNum = reflect.ValueOf(NewI).MethodByName("Sub").Call([]reflect.Value{reflect.ValueOf(NewI), reflect.ValueOf(NewJ)})[0].Interface()
	} else {
		tmp.RealNum = reflect.ValueOf(i).MethodByName("Sub").Call([]reflect.Value{reflect.ValueOf(i), reflect.ValueOf(j)})[0].Interface()
	}

	if reflect.TypeOf(k) != reflect.TypeOf(l) {
		NewK, NewL := conversion(k, l)
		tmp.ImaginaryNum = reflect.ValueOf(NewK).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(NewK), reflect.ValueOf(NewL)})[0].Interface()
	} else {
		tmp.ImaginaryNum = reflect.ValueOf(i).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(k), reflect.ValueOf(l)})[0].Interface()
	}
	//if ImaginaryNum=0 -> return only realnum
	if reflect.ValueOf(tmp.ImaginaryNum).MethodByName("ToString").Call([]reflect.Value{})[0].Interface().(string) == "0" {
		return tmp.RealNum
	} else {
		return tmp
	}
}
func (c Complex) Div(a Complex, b Complex) interface{} {
	//a+bi/c+di = (ac+bd)/(c^2+d^2) - (ad-bc)/(c^2+d^2)
	i, j, k, l := multiple(a.RealNum, a.ImaginaryNum, b.RealNum, b.ImaginaryNum)
	//ac = i, bd = j, ad =k, bc= l
	m, n, _, _ := multiple(b.RealNum, b.ImaginaryNum, b.RealNum, b.ImaginaryNum)
	//m=c^2, n=d^2
	var tmp Complex

	//calculate ac+bd, ad-bc
	if reflect.TypeOf(i) != reflect.TypeOf(j) {
		NewI, NewJ := conversion(i, j)
		tmp.RealNum = reflect.ValueOf(NewI).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(NewI), reflect.ValueOf(NewJ)})[0].Interface()
	} else {
		tmp.RealNum = reflect.ValueOf(i).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(i), reflect.ValueOf(j)})[0].Interface()
	}

	if reflect.TypeOf(k) != reflect.TypeOf(l) {
		NewK, NewL := conversion(k, l)
		tmp.ImaginaryNum = reflect.ValueOf(NewK).MethodByName("Sub").Call([]reflect.Value{reflect.ValueOf(NewK), reflect.ValueOf(NewL)})[0].Interface()
	} else {
		tmp.ImaginaryNum = reflect.ValueOf(i).MethodByName("Sub").Call([]reflect.Value{reflect.ValueOf(k), reflect.ValueOf(l)})[0].Interface()
	}

	//calculate c^2+d^2
	var tmp_divided interface{}
	if reflect.TypeOf(m) != reflect.TypeOf(n) {
		NewM, NewN := conversion(m, n)
		tmp_divided = reflect.ValueOf(NewM).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(NewM), reflect.ValueOf(NewN)})[0].Interface()
	} else {
		tmp_divided = reflect.ValueOf(i).MethodByName("Sum").Call([]reflect.Value{reflect.ValueOf(m), reflect.ValueOf(n)})[0].Interface()
	}

	if reflect.ValueOf(tmp_divided).MethodByName("ToString").Call([]reflect.Value{})[0].Interface().(string) != "0" {
		//divide RealNum, ImaginaryNum by (c^2+d^2) if c^2 + d^2 is not 0
		if reflect.TypeOf(tmp.RealNum) != reflect.TypeOf(tmp_divided) {
			newR, newD := conversion(tmp.RealNum, tmp_divided)
			tmp.RealNum = reflect.ValueOf(newR).MethodByName("Div").Call([]reflect.Value{reflect.ValueOf(newR), reflect.ValueOf(newD)})[0].Interface()
		} else {
			tmp.RealNum = reflect.ValueOf(tmp.RealNum).MethodByName("Div").Call([]reflect.Value{reflect.ValueOf(tmp.RealNum), reflect.ValueOf(tmp_divided)})[0].Interface()
		}
		if reflect.TypeOf(tmp.ImaginaryNum) != reflect.TypeOf(tmp_divided) {
			newR, newD := conversion(tmp.ImaginaryNum, tmp_divided)
			tmp.ImaginaryNum = reflect.ValueOf(newR).MethodByName("Div").Call([]reflect.Value{reflect.ValueOf(newR), reflect.ValueOf(newD)})[0].Interface()
		} else {
			tmp.ImaginaryNum = reflect.ValueOf(tmp.ImaginaryNum).MethodByName("Div").Call([]reflect.Value{reflect.ValueOf(tmp.ImaginaryNum), reflect.ValueOf(tmp_divided)})[0].Interface()
		}

		//if ImaginaryNum=0 -> return only realnum
		if reflect.ValueOf(tmp.ImaginaryNum).MethodByName("ToString").Call([]reflect.Value{})[0].Interface().(string) == "0" {
			return tmp.RealNum
		} else {
			return tmp
		}
	} else {
		panic("can't calculate")
	}
}

func (c Complex) Negative() Complex {
	c.RealNum = reflect.ValueOf(c.RealNum).MethodByName("Negative").Call([]reflect.Value{})[0].Interface()
	c.ImaginaryNum = reflect.ValueOf(c.ImaginaryNum).MethodByName("Negative").Call([]reflect.Value{})[0].Interface()
	return c
}
