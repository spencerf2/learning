package main

import "fmt"

type Counter struct {
	value int
}

func (c Counter) GetValue() int {
	return c.value
}

func (c Counter) SetValue(newValue int) {
	c.value = newValue
}

func CreateCounter(initialValue int) Counter {
	return Counter{initialValue}
}

func main() {
	// Test your code:
	c := CreateCounter(5)
	fmt.Println(c.GetValue()) // Should print 5
	c.SetValue(10)
	fmt.Println(c.GetValue()) // Should print 10
}

// Answering the questions:
// 1. What's the difference between func (c Counter) GetValue() and func GetValue(c Counter)?
//   - The first one is a method. The second one is a function. The method performs an
//     action on the struct Counter, which is represented by the variable c.
//     The second one takes the struct as a variable to do something with it.
//     I sense there may be more to this though, since the function would then have...
//     is this where pointer reference vs value or copy comes into play? I think so.
//     The function would have a copy of the struct rather than a pointer to the original.
//     Need to confirm this though when I do problem 2.
// 
// 2. Why do we use methods instead of just functions?
//   - Because it allows us to perform related actions on the struct, rather than having
//     to pass the struct to then act upon it, which would require passing a pointer I
//     believe.
