package functions

import "fmt"

func ShowLoop() {
	myLoop()
}

func myLoop() {
	// Conditional For
	for i :=0; i<= 10; i++{
		fmt.Println(i)
	}

	// For While
	counter := 0
	for counter <= 10{
		fmt.Println(counter)
		counter++
	}

	// Forever For
	counterForever := 0
	for counterForever <= 10{
		fmt.Println(counterForever)
	}
}