package functions

import "fmt"
import _ "../models"

func ShowStructs()  {
	instanceOne()
	instanceTwo()
}

func instanceOne()  {
	var myCar Car = car{brand: "Ford",year: 2020}

	fmt.Printf("brand: %s - year: %v\n", myCar.brand, myCar.year)
}

func instanceTwo()  {
	var myCar car
	myCar.brand = "Ferrari"

	fmt.Printf("brand: %s - year: %v\n", myCar.brand, myCar.year)
}
