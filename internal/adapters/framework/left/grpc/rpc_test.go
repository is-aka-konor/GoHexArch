package rpc

import (
	"GoHexArch/internal/adapters/app/api"
	"GoHexArch/internal/adapters/core/arithmetic"
	"GoHexArch/internal/adapters/framework/left/grpc/pb"
	"GoHexArch/internal/adapters/framework/right/db"
	"GoHexArch/internal/ports"
	"context"
	"log"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const buffer = 1024 * 1024

var listener *bufconn.Listener

func init() {
	listener = bufconn.Listen(buffer)
	server := grpc.NewServer()
	var err error
	//ports
	var dbAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	// Configure the db adapter

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")
	dbAdapter, err = db.NewAdapter(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Could not create db adapter: %v", err)
	}
	//defer dbAdapter.CloseDbConnection()

	// Configure the core adapter
	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, core)
	gRPCAdapter = NewAdapter(appAdapter)
	pb.RegisterArithmeticServiceServer(server, gRPCAdapter)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Test server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func getGRPCConnection(t *testing.T) *grpc.ClientConn {
	conn, err := grpc.NewClient("passthrough:whatever", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial passthrough:whatever: %v", err)
	}
	return conn
}

func TestSum(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(t)
	defer conn.Close()
	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 2, B: 3}
	resp, err := client.GetSum(ctx, params)
	if err != nil {
		t.Fatalf("Sum failed: %v", err)
	}
	require.Equal(t, int32(5), resp.Value)
}

func TestSub(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(t)
	defer conn.Close()
	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 2, B: 3}
	resp, err := client.GetSub(ctx, params)
	if err != nil {
		t.Fatalf("Sub failed: %v", err)
	}
	require.Equal(t, int32(-1), resp.Value)
}

func TestMul(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(t)
	defer conn.Close()
	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 2, B: 3}
	resp, err := client.GetMul(ctx, params)
	if err != nil {
		t.Fatalf("Mul failed: %v", err)
	}
	require.Equal(t, int32(6), resp.Value)
}

func TestDiv(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(t)
	defer conn.Close()
	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 6, B: 3}
	resp, err := client.GetDiv(ctx, params)
	if err != nil {
		t.Fatalf("Div failed: %v", err)
	}
	require.Equal(t, int32(2), resp.Value)
}
