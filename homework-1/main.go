package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func currencyExchange() {
	const rubToUSD = 0.014 // курс на 01.06.2020
	var rub string

	fmt.Println("Сколько рублей у вас на руках?")
	fmt.Scanln(&rub)
	rubFloat, err := strconv.ParseFloat(rub, 32)
	if err != nil {
		log.Fatalln("Нужно ввести сумму цифрами, например \"17890\"")
	}

	fmt.Printf("За свои рубли вы можете получить %.2f долларов\n", rubFloat*rubToUSD)
}

func triangle() {
	var cathetus1, cathetus2 string
	var hypotenuse, triangleArea, perimeter float64

	fmt.Println("Введите длину первого катета:")
	fmt.Scanln(&cathetus1)
	cathetus1Float, err := strconv.ParseFloat(cathetus1, 64)
	if err != nil {
		log.Fatalln("Длину первого катета нужно ввести в виде числа, например 13")
	}

	fmt.Println("Теперь введите длину второго катета")
	fmt.Scanln(&cathetus2)
	cathetus2Float, err := strconv.ParseFloat(cathetus2, 64)
	if err != nil {
		log.Fatalln("Число, вводите именно число")
	}

	hypotenuse = math.Sqrt(math.Pow(cathetus1Float, 2) + math.Pow(cathetus2Float, 2))

	fmt.Printf("Длина гипотенузы равна %.2f\n", hypotenuse)

	triangleArea = (cathetus1Float * cathetus2Float) / 2

	fmt.Printf("Площадь треугольника = %.2f\n", triangleArea)

	perimeter = hypotenuse + cathetus1Float + cathetus2Float

	fmt.Printf("Периметр треугольника равен %.2f\n", perimeter)
}

func deposit() {
	var savings, interestRate string
	var totalSavings float64

	fmt.Println("Сколько денег вы хотите положить на счёт?")
	fmt.Scanln(&savings)
	savingsFloat, err := strconv.ParseFloat(savings, 64)
	if err != nil {
		log.Fatalln("Количество денег вводите цифрами, например 123984")
	}

	fmt.Println("Какой годовой процент предлагает вам банк?")
	fmt.Scanln(&interestRate)
	interestRateFloat, err := strconv.ParseFloat(interestRate, 64)
	if err != nil {
		log.Fatalln("Как и сумму вклада, процент тоже надо вводить цифрами")
	}
	interestRateFloat /= 100 // переводим 14% в 0.14

	totalSavings = savingsFloat

	for i := 0; i < 5; i++ {
		totalSavings += totalSavings * interestRateFloat
	}

	fmt.Printf("С учётом процентов, через 5 лет сумма вашего депозита (при условии, что вы не будете снимать деньги всё это время) будет равна %.2f\n", totalSavings)
}

func main() {
	currencyExchange()
	triangle()
	deposit()
}
