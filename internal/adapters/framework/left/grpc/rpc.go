package rpc

import (
	"GoHexArch/internal/adapters/framework/left/grpc/pb"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) GetSum(cts context.Context, request *pb.OperationParameters) (*pb.Result, error) {
	result, err := grpca.api.Sum(int(request.GetA()), int(request.GetB()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}
	return &pb.Result{Value: int32(result)}, nil
}

func (grpca Adapter) GetSub(cts context.Context, request *pb.OperationParameters) (*pb.Result, error) {
	result, err := grpca.api.Sub(int(request.GetA()), int(request.GetB()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}
	return &pb.Result{Value: int32(result)}, nil
}

func (grpca Adapter) GetMul(cts context.Context, request *pb.OperationParameters) (*pb.Result, error) {
	result, err := grpca.api.Mul(int(request.GetA()), int(request.GetB()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}
	return &pb.Result{Value: int32(result)}, nil
}

func (grpca Adapter) GetDiv(cts context.Context, request *pb.OperationParameters) (*pb.Result, error) {
	if request.GetB() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Error: Division by zero")
	}
	result, err := grpca.api.Div(int(request.GetA()), int(request.GetB()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}
	return &pb.Result{Value: int32(result)}, nil
}
