package main

import (
	"fmt"
	"math"
)

func main()  {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	if _, err := fmt.Scanln(&a); err != nil {
		return
	}


	fmt.Print("Введите второе число: ")
	if _, err := fmt.Scanln(&b); err != nil {
		return
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, %): ")
	if _, err := fmt.Scanln(&op); err != nil {
		return
	}


	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Print("Делить на 0 нельзя")
		} else {
			res = a / b
		}
	case "%":
		res = math.Mod(a, b)
	default:
		fmt.Print("Для операций с числами используйте для ввода только +, -, *, /")
		break
	}

	fmt.Printf("\nРезультат выполнения операции: %.2f", res)
}
