package ports

// import (
// 	"context"
// )

type GRPCPort interface {
	Run()
	GetSum()
	GetSub()
	GetMul()
	GetDiv()
}
