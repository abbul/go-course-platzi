package functions

import "fmt"

func ShowKeyword()  {
	caseDefer()
	caseContinue()
}
func caseDefer()  {
	// Defer
	defer fmt.Println("End")
	fmt.Println("Hello World")
	fmt.Println("\t")
}

func caseContinue()  {
	// Defer
	for i:=0; i <= 10; i++ {

		if i == 5 {
			continue
		}
		if i > 7 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("\t")
}
