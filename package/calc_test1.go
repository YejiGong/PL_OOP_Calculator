package main

import (
	"Calcul"
	"fmt"
)

func main() {
	fmt.Println("계산하고싶은 숫자를 입력하세요")
	var Calc Calcul.Calc
	Calc.Enter(Calcul.NewInteger("10"))
	Calc.Enter(Calcul.NewInteger("20"))
	Calc.Add()
	Calc.Clear()

	Calc.Enter(Calcul.NewInteger("10"))
	Calc.Enter(Calcul.NewFloat("3.14159"))
	Calc.Add()
	Calc.Enter(Calcul.NewFloat("2.71828"))
	Calc.Sub()
	Calc.Clear()
}
