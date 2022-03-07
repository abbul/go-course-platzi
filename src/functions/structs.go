package functions

import (
	"fmt"
	"go-course-platzi/src/models"
)

func ShowStructs() {
	instanceOne()
	instanceTwo()
}

func instanceOne() {
	var myCar = models.Car{
		Brand:  models.Brand{Name: "Chevrolet", Country: "Unites States"},
		Year:   2015,
		Color:  "blue",
		Engine: "1.6",
		Model:  "sonic ltz",
	}

	fmt.Printf("brand: %s - year: %v\n", myCar.Brand.Name, myCar.Year)
}

func instanceTwo() {
	var myCar models.Car
	myCar.Brand.Name = "Peugeot"
	myCar.Year = 2020

	fmt.Printf("brand: %s - year: %v\n", myCar.Brand.Name, myCar.Year)
}
