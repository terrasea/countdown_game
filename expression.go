package main

import (
	"bytes"
	"time"
	"math/rand"
	"fmt"
	"strconv"
	"github.com/soniah/evaler"
)

func heristic(expression string, total int) (int, int) {
	result, err := evaler.Eval(expression)
	if err != nil {
		return 10000, -1
	}
	value, _ := result.Float64()

	if int(value) < 0 {
		return 100000, int(value)
	} else if total - int(value) >= 0 && int(value) - total >= 0 && total - int(value) < int(value) - total {
		return total - int(value), int(value)
	} else if int(value) - total >= 0 {
		return int(value) - total, int(value)
	} else {
		return 1000000000, int(value)
	}
	
}

func getRandOperator() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	op := r1.Intn(4)
	switch op {
	case 0:
		return "*"
	case 1:
		return "/"
	case 2:
		return "+"
	case 3:
		return "-"
	}

	return ""
}

func createExpression(numbers []int) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var selected = make([]int, 0)
	var expression bytes.Buffer
	for _, number := range numbers {
		flip := r1.Intn(2)
		if flip > 0 {
			selected = append(selected, number)
		}
	}

	for pos, number := range selected {
		op := getRandOperator()
		expression.WriteString(strconv.Itoa(number))
		if pos < len(selected) - 1 {
			expression.WriteString(op)
		}
	}

	return expression.String()
}


func main() {
	bestresult := -1
	expression := ""
	value := 0
	bestvalue := 0
	bestexpression := ""
	
	for i := 0; i < 100000; i++ {
		expression = createExpression([]int{10, 10, 10, 4, 5, 6, 7, 8, 9, 10})
		var result int
		result, value = heristic(expression, 567)
		if result < bestresult || bestresult == -1 {
			bestresult = result
			bestvalue = value
			bestexpression = expression
		}

		if bestresult == 0 {
			break
		}
	}
	fmt.Println(bestexpression, ": ", bestresult, ", ", bestvalue)
}
