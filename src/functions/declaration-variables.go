package functions

import "fmt"

func ShowDeclarationVariables() {
	declarationVariables()
}

func declarationVariables() {
	// Variables Declaration
	var name string = "Abbul"
	var lastname string = "Rodriguez"
	var age uint8 = 29
	const dni uint32 = 95419753
	var isAdmin bool = true
	fmt.Println("User: ",name, lastname)
	fmt.Println("Age: ",age)
	fmt.Println("DNI: ",dni)
	fmt.Println("Admin: ",isAdmin)
	fmt.Println("\t")

}