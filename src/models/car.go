package models

type Brand struct {
	Name    string
	Country string
}

// Car it's basic
type Car struct {
	Brand  Brand
	Year   int
	Color  string
	Engine string
	Model  string
}
