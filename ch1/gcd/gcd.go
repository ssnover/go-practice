package main

import "fmt"

func main() {
	fmt.Print("Enter two numbers separated by a space x y: ")
	var num1, num2 int
	_, err := fmt.Scanf("%d", &num1)
	_, err = fmt.Scanf("%d", &num2)
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	var smaller int
	if (num1 < num2) {
		smaller = num1
	} else {
		smaller = num2
	}

	var gcd int = 0
	for i := smaller; i > 0; i-- {
		if (num1 % i == 0) && (num2 % i == 0) {
			gcd = i
			break
		}
	}

	if gcd == 0 {
		fmt.Println("There is no Greatest Common Divisor.")
	} else {
		fmt.Printf("The Greatest Common Divisor of %d and %d is %d\n", num1, num2, gcd)
	}
}