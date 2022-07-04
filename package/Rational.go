package main

import (
	"Calcul"
	"reflect"
	"strconv"
	"strings"
)

type Rational struct {
	Numerator   int
	Dinominator int
	Calcul.Number
}

func NewRational(a string) Rational {
	var tmp Rational
	rStr := strings.Split(a, "/")
	tmp.Numerator, _ = strconv.Atoi(rStr[0])
	tmp.Dinominator, _ = strconv.Atoi(rStr[1])
	return tmp
}
func (r Rational) GetValue() Rational {
	return r
}
func (r Rational) SetValue(a Rational) Rational {
	return r
}
func (r Rational) ToString() string {
	var tmp string
	tmp_Numerator := strconv.Itoa(r.Numerator)
	tmp_Dinominator := strconv.Itoa(r.Dinominator)
	tmp = tmp_Numerator + "/" + tmp_Dinominator
	return tmp
}

func (r Rational) Convert(a interface{}) (bool, interface{}, interface{}) {
	var new_a, new_b interface{}
	var checker bool = false
	switch a.(type) {
	case Calcul.Integer:
		//r -> Rational. a -> Rational
		var tmp Rational
		tmp_value := reflect.ValueOf(a).MethodByName("ToString").Call([]reflect.Value{})[0].Interface().(string)
		tmp.Numerator, _ = strconv.Atoi(tmp_value)
		if tmp.Numerator != 0 {
			tmp.Dinominator, _ = strconv.Atoi("1")
			new_b = reflect.ValueOf(tmp).Interface()
		} else {
			new_b = Calcul.NewInteger("0")
		}
		new_a = r
		checker = true
	case Calcul.Float:
		//r -> float. a -> float
		tmp_value := float64(r.Numerator) / float64(r.Dinominator)
		new_a = Calcul.NewFloat(strconv.FormatFloat(tmp_value, 'f', -1, 64))
		new_b = reflect.ValueOf(a).Interface()
		checker = true
	}
	return checker, new_a, new_b
}

func commensuration(a Rational) interface{} {
	var tmp Rational
	var gcd int = 1
	if a.Numerator > a.Dinominator {
		for i := 1; i < a.Numerator; i++ {
			if a.Numerator%i == 0 && a.Dinominator%i == 0 {
				gcd = i
			}
		}
		tmp.Numerator = a.Numerator / gcd
		tmp.Dinominator = a.Dinominator / gcd
	} else if a.Numerator == a.Dinominator {
		return Calcul.NewInteger("1")
	} else {
		for i := 1; i < a.Dinominator; i++ {
			if a.Numerator%i == 0 && a.Dinominator%i == 0 {
				gcd = i
			}
		}
		tmp.Numerator = a.Numerator / gcd
		tmp.Dinominator = a.Dinominator / gcd
	}
	if tmp.Dinominator == 1 {
		return Calcul.NewInteger(strconv.Itoa(tmp.Numerator))
	} else {
		return tmp
	}
}
func (r Rational) Sum(a Rational, b Rational) interface{} {
	var tmp Rational
	tmp.Numerator = (a.Numerator * b.Dinominator) + (b.Numerator * a.Dinominator)
	tmp.Dinominator = a.Dinominator * b.Dinominator
	//if Dinominator == 0 : it can't be
	if tmp.Dinominator == 0 {
		return nil
	} else {
	}
	//if Numerator == 0 : it just value of 0
	if tmp.Numerator == 0 {
		return Calcul.NewInteger("0")
	} else {
		return commensuration(tmp)
	}
}
func (r Rational) Sub(a Rational, b Rational) interface{} {
	var tmp Rational
	tmp.Numerator = (a.Numerator * b.Dinominator) - (b.Numerator * a.Dinominator)
	tmp.Dinominator = a.Dinominator * b.Dinominator
	//if Dinominator == 0 : it can't be
	if tmp.Dinominator == 0 {
		return nil
	} else {
	}
	//if Numerator == 0 : it just value of 0
	if tmp.Numerator == 0 {
		return Calcul.NewInteger("0")
	} else {
		return commensuration(tmp)
	}
}
func (r Rational) Mult(a Rational, b Rational) interface{} {
	var tmp Rational
	tmp.Numerator = a.Numerator * b.Numerator
	tmp.Dinominator = a.Dinominator * b.Dinominator
	//if Dinominator == 0 : it can't be
	if tmp.Dinominator == 0 {
		panic("can't calculate")
	} else {
	}
	//if Numerator == 0 : it just value of 0
	if tmp.Numerator == 0 {
		return Calcul.NewInteger("0")
	} else {
		return commensuration(tmp)
	}
}
func (r Rational) Div(a Rational, b Rational) interface{} {
	var tmp Rational
	tmp.Numerator = a.Numerator * b.Dinominator
	tmp.Dinominator = a.Dinominator * b.Numerator
	//if Dinominator == 0 : it can't be
	if tmp.Dinominator == 0 {
		panic("can't calculate")
	} else {
	}
	//if Numerator == 0 : it just value of 0
	if tmp.Numerator == 0 {
		return Calcul.NewInteger("0")
	} else {
		return commensuration(tmp)
	}
}

func (r Rational) Negative() Rational {
	r.Numerator = 0 - r.Numerator
	return r
}
