package main

import (
	"Calcul"
	"fmt"
)

func main() {
	fmt.Println("계산하고싶은 숫자를 입력하세요")
	var Calc Calcul.Calc
	//1+1/2+1/3+1/4-1/12
	Calc.Enter(Calcul.NewInteger("1"))
	Calc.Enter(NewRational("1/2"))
	Calc.Add()
	Calc.Enter(NewRational("1/3"))
	Calc.Add()
	Calc.Enter(NewRational("1/4"))
	Calc.Add()
	Calc.Enter(NewRational("1/12"))
	Calc.Sub()
	Calc.Clear()

	//2/3-10+10.5
	Calc.Enter(NewRational("2/3"))
	Calc.Enter(Calcul.NewInteger("10"))
	Calc.Sub()
	Calc.Enter(Calcul.NewFloat("10.5"))
	Calc.Add()
	Calc.Clear()

	//1.2i+3.5i+3-1/2+2.5
	Calc.Enter(NewComplex("1+2i"))
	Calc.Enter(NewComplex("3+5i"))
	Calc.Add()
	Calc.Enter(Calcul.NewInteger("3"))
	Calc.Add()
	Calc.Enter(NewRational("1/2"))
	Calc.Sub()
	Calc.Enter(Calcul.NewFloat("2.5"))
	Calc.Add()
	Calc.Clear()
}
