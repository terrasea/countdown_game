package main

import (
	"fmt"
)

type Expression struct {
	operand string
	lhs float64
	rhs float64
}


func (ex *Expression) execute() float64 {
	switch ex.operand {
	case "*":
		fallthrough
	case "x":
		return ex.lhs * ex.rhs
	case "/":
		return ex.lhs / ex.rhs
	case "+":
		return ex.lhs + ex.rhs
	case "-":
		return ex.lhs - ex.rhs
	}

	return 0
}




func main() {
	numbers := []float64{1, 2, 100, 10, 23, 17}
	ex := Expression{"/", 2, 4}

	//{operand: byte('*'), lhs: 2, rhs: 2};
	//ex.operand = '*'
	//ex.lhs = 2
	//ex.rhs = 2

	fmt.Print(numbers, ex, ex.execute(), "\n")
}
