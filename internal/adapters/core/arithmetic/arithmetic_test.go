package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TesSum(t *testing.T) {
	arithmetics := NewAdapter()
	result, err := arithmetics.Sum(2, 3)
	if err != nil {
		t.Fatalf("Add operation failed, expected %v, got: %v", nil, err)
	}
	require.Equal(t, 5, result)
}

func TestSub(t *testing.T) {
	arithmetics := NewAdapter()
	result, err := arithmetics.Sub(2, 3)
	if err != nil {
		t.Fatalf("Sub operation failed, expected %v, got: %v", nil, err)
	}
	require.Equal(t, -1, result)
}

func TestMul(t *testing.T) {
	arithmetics := NewAdapter()
	result, err := arithmetics.Mul(2, 3)
	if err != nil {
		t.Fatalf("Mul operation failed, expected %v, got: %v", nil, err)
	}
	require.Equal(t, 6, result)
}

func TestDiv(t *testing.T) {
	arithmetics := NewAdapter()
	result, err := arithmetics.Div(6, 3)
	if err != nil {
		t.Fatalf("Div operation failed, expected %v, got: %v", nil, err)
	}
	require.Equal(t, 2, result)
}

func TestDivByZero(t *testing.T) {
	arithmetics := NewAdapter()
	result, err := arithmetics.Div(6, 0)
	if err == nil {
		t.Fatalf("Div operation failed, expected %v, got: %v", "division by zero", err)
	}
	require.Equal(t, 0, result)
}
