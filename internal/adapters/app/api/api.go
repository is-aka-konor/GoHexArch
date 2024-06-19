package api

import (
	"GoHexArch/internal/ports"
)

type Adapter struct {
	db         ports.DbPort
	arithmetic ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arithmetic ports.ArithmeticPort) *Adapter {
	return &Adapter{
		db:         db,
		arithmetic: arithmetic,
	}
}

func (apiAdapter Adapter) GetSum(a, b int) (int, error) {
	result, err := apiAdapter.arithmetic.Sum(a, b)
	if err != nil {
		return 0, err
	}

	// Save the result to the database
	err = apiAdapter.db.AddToHistory(result, "GetSum")
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (apiAdapter Adapter) GetSub(a, b int) (int, error) {
	result, err := apiAdapter.arithmetic.Sub(a, b)
	if err != nil {
		return 0, err
	}

	// Save the result to the database
	err = apiAdapter.db.AddToHistory(result, "GetSub")
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (apiAdapter Adapter) GetMul(a, b int) (int, error) {
	result, err := apiAdapter.arithmetic.Mul(a, b)
	if err != nil {
		return 0, err
	}

	// Save the result to the database
	err = apiAdapter.db.AddToHistory(result, "GetMul")
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (apiAdapter Adapter) GetDiv(a, b int) (int, error) {
	result, err := apiAdapter.arithmetic.Div(a, b)
	if err != nil {
		return 0, err
	}
	// Save the result to the database
	err = apiAdapter.db.AddToHistory(result, "GetDiv")
	if err != nil {
		return 0, err
	}
	return result, nil
}
