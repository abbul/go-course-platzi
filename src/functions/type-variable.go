package functions
import "fmt"

func ShowTypeVariable() {
 	typeVariable()
}

func typeVariable() {

	// Default Value
	var defaultStringValue string
	var defaultIntValue int32
	var defaultUnitValue uint32
	var defaultBooleanValue bool
	fmt.Println("defaultStringValue: ",defaultStringValue)
	fmt.Println("defaultIntValue: ",defaultIntValue)
	fmt.Println("defaultUnitValue: ",defaultUnitValue)
	fmt.Println("defaultBooleanValue: ",defaultBooleanValue)
	fmt.Println("\t")
}