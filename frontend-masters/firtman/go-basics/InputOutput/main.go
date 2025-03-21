package main

import "fmt"

var name = "Frontend Masters"

func calculateTax(price float32) (float32, float32) {
	return price*0.09, price*0.02
}

func birthday(pointerAge *int) {
	if (*pointerAge>140) {
		panic("Too old to be true")
	}

	fmt.Printf("The pointer is %v and the value is %v", pointerAge, *pointerAge)
	*pointerAge++
}

func main() {
	// stateTax, _ := calculateTax(100)
	// fmt.Println(stateTax)

	defer fmt.Println("Bye!!")
	defer fmt.Println("Good ")

	age := 22
	birthday(&age)
	fmt.Println(age)

	PrintData()
}
