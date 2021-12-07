package functions

import "fmt"

func ShowMaps() {
	keyValue()
}

func keyValue(){
	var myArray = make(map[string]int)

	myArray["aliuska"] = 21
	myArray["abbul"] = 29

	for k, v := range myArray {
		fmt.Printf("key: %v - value: %v \n",k,v)
	}

	var value, keyExist = myArray["invaludKey"]
	fmt.Printf("value: %v - key exist: %v \n",value,keyExist)
}