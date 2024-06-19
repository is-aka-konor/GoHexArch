package ports

type APIPort interface {
	GetSum(a, b int) (int, error)
	GetSub(a, b int) (int, error)
	GetMul(a, b int) (int, error)
	GetDiv(a, b int) (int, error)
}
