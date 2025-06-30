package main

import "fmt"

type Counter struct {
	value int
}

func (c Counter) IncrementValue() {
	c.value++
}

func (c *Counter) IncrementPointer() {
	c.value++
}

func main() {
	c := Counter{value: 0}

	c.IncrementValue()
	fmt.Println(c.value) // What prints here? Why?
	// 
	// 0 prints here. Because the function updates the value, but that just means
	// it updates a copy of the actual value, rather than updating the value in
	// memory, which is what the struct is holding onto (aka "pointing" to).

	c.IncrementPointer()
	fmt.Println(c.value) // What prints here? Why?
	// 
	// 1 should print here, because now the function is actually updating the value
	// held by the variable in memory. The struct is also something held in memory
	// which points to the variable.
}

// Answering the questions:
// 1. Why does IncrementValue() not change the counter?
//   - Because it updates the actual value, which is just a number, not the variable's
//     value that's held in memory. And it doesn't do anything with it due to the scope
//     of the function. As soon as the update to the number is made, the function
//     closes and it's lost.
// 
// 2. Why does IncrementPointer() change the counter?
//   - It updates the number that's held at the location the variable points to in
//     memory.
// 
// 3. When would you use each approach?
//   - You'd use the value approach when you just want to do something with the value.
//     Whatever you do is a downstream type operation. Updates made to that value don't
//     need to make their way back to the variable itself.
//     When you want to update the variable itself, then you need to pass a pointer to
//     it so that the updates are carried back up to the actual location in memory, which
//     is what the variable, and therefore struct (since the struct points to or owns
//     the variable) point to.
