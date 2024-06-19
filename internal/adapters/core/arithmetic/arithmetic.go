package arithmetic

import (
	"fmt"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith *Adapter) Sum(a, b int) (int, error) {
	return a + b, nil
}

func (arith *Adapter) Sub(a, b int) (int, error) {
	return a - b, nil
}

func (arith *Adapter) Mul(a, b int) (int, error) {
	return a * b, nil
}

func (arith *Adapter) Div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
