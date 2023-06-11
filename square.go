package figures

type Square struct {
	Side float64
}

func (squ *Square) area() float64 {
	return squ.Side * squ.Side
}

func (squ *Square) perimeter() float64 {
	return 4 * squ.Side
}
