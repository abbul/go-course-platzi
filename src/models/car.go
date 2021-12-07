package models

type brand struct {
	name string
	country string
}

// Car it's basic
type Car struct {
	brand brand
	year int
	color string
	engine string
	model string
}