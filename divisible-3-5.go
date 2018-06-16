package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	fmt.Print("Enter number to sum to: ")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Was that an integer, I'm not so sure...")
	} else {
		var sum int
		for j := 1; j <= i; j++ {
			if (j % 3 == 0) || (j % 5 == 0) {
				sum += j
			}
		}
		fmt.Printf("The sum is: %d\n", sum)
	}
}