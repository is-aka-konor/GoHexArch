package ports

import (
	"GoHexArch/internal/adapters/framework/left/grpc/pb"
	"context"
)

type GRPCPort interface {
	Run()
	GetSum(cts context.Context, request *pb.OperationParameters) (*pb.Result, error)
	GetSub(cts context.Context, request *pb.OperationParameters) (*pb.Result, error)
	GetMul(cts context.Context, request *pb.OperationParameters) (*pb.Result, error)
	GetDiv(cts context.Context, request *pb.OperationParameters) (*pb.Result, error)
}
