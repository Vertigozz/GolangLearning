package main

import (
	"fmt"
)

/*Напишите приложение, которое принимает на вход набор целых чисел и отдаёт на выходе его же в отсортированном виде.
Сохраните код, он понадобится нам в дальнейших уроках.*/

func main() {
	var a, b, c, d, e, f int

	fmt.Println("Введите последовательно 5 чисел для сортировки: ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	fmt.Scanln(&e)
	fmt.Scanln(&f)

	inputNumbers := []int{a, b, c, d, f, e}

	fmt.Println("\nНеотсортированный набор чисел\n", inputNumbers)

	insertionSort(inputNumbers)
	fmt.Println("\nОтсортированный набор чисел\n", inputNumbers, "\n")

}

func insertionSort(inputNumbers []int) {
	var n = len(inputNumbers)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if inputNumbers[j-1] > inputNumbers[j] {
				inputNumbers[j-1], inputNumbers[j] = inputNumbers[j], inputNumbers[j-1]
			}
			j = j - 1
		}
	}
}