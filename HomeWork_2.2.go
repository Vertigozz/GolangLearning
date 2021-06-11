package main

import (
	"fmt"
	"math"
)
/*2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
Площадь круга должна вводиться пользователем с клавиатуры.*/

func main() {
	var a float64

	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&a)

	fmt.Printf("Диаметр круга равен: %.2f", math.Sqrt(a/math.Pi)*2)

	fmt.Printf("\nДлина окружности равна: %.2f", math.Sqrt(a/math.Pi)*2*math.Pi)
}
