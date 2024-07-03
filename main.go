package main

import "fmt"

func main() {
	var a int
	var b int
	var action string

	fmt.Println("Мой калькулятор")
	fmt.Println("Введите первое число:")
	fmt.Scan(&a)

	fmt.Println("Какое действие выполнить? (+, -, *, /)")
	fmt.Scan(&action)

	fmt.Println("Введите второе число:")
	fmt.Scan(&b)

	fmt.Println("Результат:")

	if a > 10 || a < 1 || b > 10 || b < 1 {
		panic("Используйте числа от 1 до 10")
		return
	}

	switch {
	case action == "+":

		fmt.Println(a + b)
	case action == "-":

		fmt.Println(a - b)
	case action == "*":

		fmt.Println(a * b)
	case action == "/":

		fmt.Println(a / b)
	default:
		fmt.Println("Паника! Только эти действия: (+, -, *, /)")
	}

}
