package main

import (
	"fmt"
	"log"
	"strconv"
)

// even функция возвращает строку с ответом на вопрос является ли число чётным
func even(num int) string {
	if num%2 == 0 {
		return "Чётное число"
	}

	return "Нечётное число"
}

// divideBy3 возвращает строку с ответом на вопрос делится число на 3 без остатка
func divideBy3(num int) string {
	if num%3 == 0 {
		return "Число делится на три без остатка"
	}

	return "Число не делится на три без остатка"
}

// fibonacci функция рекурсивно вычисляет число Фибоначчи для заданного аргумента
func fibonacci(num int) int {
	if num < 2 {
		return 1
	}
	return fibonacci(num-2) + fibonacci(num-1)
}

func main() {
	var userNumber, userNum3 string

	fmt.Println("Введите число:")
	fmt.Scanln(&userNumber)

	userInt, err := strconv.Atoi(userNumber)
	if err != nil {
		log.Fatalln("Введите число")
	}

	fmt.Println(even(userInt))

	fmt.Println("Проверим, делится ли число на 3 без остатка. Введите число:")
	fmt.Scanln(&userNum3)

	userInt, err = strconv.Atoi(userNumber)
	if err != nil {
		log.Fatalln("Введите число")
	}

	fmt.Println(divideBy3(userInt))

	fmt.Println("Выводим первые 100 чисел Фибоначчи:")

	for i := 0; i <= 100; i++ {
		fmt.Printf("%d, ", fibonacci(i))
	}
}
