package functions

import "fmt"

func ShowPackageFmt() {

	// Packet FMT
	const greeting = "Hi!!"
	fmt.Println("Println: ",greeting)
	fmt.Printf("Printf: %s \n", greeting)
	fmt.Printf("Printf (when know'n type): %v \n", greeting)
	var message string = fmt.Sprintf("Sprintf: %s \n", greeting)
	fmt.Printf(message)

}