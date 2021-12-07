package functions

import "fmt"

func ShowArray()  {
	mapSlice()
}

func mapSlice()  {
	var mySlice = []string{"Soy","Abbul","Rodriguez"}

	for i := range mySlice {
		fmt.Printf("index: %v \n", i)
	}

	for _, valor := range mySlice {
		fmt.Printf("value: %v \n", valor)
	}

	for i, valor := range mySlice {
		fmt.Printf("index: %v - value: %s \n", i,valor)
	}
}

func arrayTest()  {
	var myArray [6]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	fmt.Println("-----Array--------")
	fmt.Println("data:", myArray)
	fmt.Println("length:", len(myArray))
	fmt.Println("capacity:", cap(myArray))
}

func sliceTest()  {
	var mySlice = []int{100,200,300,400,500}
	fmt.Println("-----Slice--------")

	fmt.Println("index 0:", mySlice[0])
	fmt.Println("index >=4:", mySlice[4:])
	fmt.Println("index <3", mySlice[:3])
	fmt.Println("index >=2 and <4:", mySlice[2:4])
}

func sliceAppendTest()  {
	var mySlice = []int{100,200,300,400,500}

	mySlice = append(mySlice,88)

	fmt.Println("-----SliceAppend--------")
	fmt.Println("data:", mySlice)
	fmt.Println("length:", len(mySlice))
	fmt.Println("capacity:", cap(mySlice))
}

func sliceDestructuringTest()  {
	var mySlice = []int{100,200,300,400,500}
	var newSlice = []int{22,33,44}

	mySlice = append(mySlice,newSlice...)

	fmt.Println("-----SliceDestructuring--------")
	fmt.Println("data:", mySlice)
	fmt.Println("length:", len(mySlice))
	fmt.Println("capacity:", cap(mySlice))
}