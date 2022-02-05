package main

import "fmt"

func main() {
	fmt.Println("ИНженерный калькулятор")
	fmt.Println("Какое действие Вы хотите выполнить ?(+, -, *, /, %)")

	var action string
	fmt.Scan(&action)
	var a float64
	var b float64
	fmt.Println("Введите первое число:")
	fmt.Scan(&a)
	fmt.Println("Введите второе число:")
	fmt.Scan(&b)
	switch {
	case action == "+":
		fmt.Println("Сумма чисел " + fmt.Sprint(a+b))

	case action == "-":
		fmt.Println("Сумма чисел " + fmt.Sprint(a-b))
	case action == "*":
		fmt.Println("Сумма чисел " + fmt.Sprint(a*b))
	case action == "/":
		fmt.Println("Сумма чисел " + fmt.Sprint(a/b))
	case action == "%":
		fmt.Println("Остаток от деления:" + fmt.Sprint(int64(a)%int64(b)))
	default:
		fmt.Println("There are no  such numbers")
	}
}
