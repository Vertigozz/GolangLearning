package main

import (
	"fmt"
)


/*1. Напишите приложение, рекурсивно вычисляющее заданное
из стандартного ввода число Фибоначчи.
2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.*/

func main() {
	var number int
	fmt.Println("Введите число, соответствующее желаемой длине ряда Фибоначчи: ")
	fmt.Scanln(&number)
	for i, first, second := 1, 0, 1; i <=number; i++ {
		first, second = second, first + second
		fmt.Println(first, " ")
	}
fmt.Println()
}





