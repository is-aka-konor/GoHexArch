package ports

type ArithmeticPort interface {
	Sum(a, b int) (int, error)
	Sub(a, b int) (int, error)
	Mul(a, b int) (int, error)
	Div(a, b int) (int, error)
}
