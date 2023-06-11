package figures

import "fmt"

type Geometry interface {
	area() float64
	perimeter() float64
}

func Measures(g Geometry) {
	fmt.Println("Messures: ", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimeter: ", g.perimeter())
}
