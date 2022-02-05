package main

import "fmt"

func main() {
	fmt.Println("Engineering Calculator")
	fmt.Println("What action do you want to perform ?(+, -, *, /, %)")

	var action string
	fmt.Scan(&action)
	var a float64
	var b float64
	fmt.Println("Enter the first number:")
	fmt.Scan(&a)
	fmt.Println("Enter the Second number:")
	fmt.Scan(&b)
	switch {
	case action == "+":
		fmt.Println("Sum of numbers " + fmt.Sprint(a+b))

	case action == "-":
		fmt.Println("Sum of numbers " + fmt.Sprint(a-b))
	case action == "*":
		fmt.Println("Sum of numbers " + fmt.Sprint(a*b))
	case action == "/":
		fmt.Println("Сумма чисел " + fmt.Sprint(a/b))
	case action == "%":
		fmt.Println("The remainder of the division:" + fmt.Sprint(int64(a)%int64(b)))
	default:
		fmt.Println("There are no  such numbers")
	}
}
