package main

import (
	"GoHexArch/internal/adapters/app/api"
	"GoHexArch/internal/adapters/core/arithmetic"
	gRPC "GoHexArch/internal/adapters/framework/left/grpc"
	"GoHexArch/internal/adapters/framework/right/db"
	"GoHexArch/internal/ports"
	"fmt"
	"log"
	"os"
)

func main() {
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
	defer dbAdapter.CloseDbConnection()

	// Configure the core adapter
	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, core)
	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()

	fmt.Println("Hello, World!")
	arithmeticAdapter := arithmetic.NewAdapter()
	result, err := arithmeticAdapter.Div(10, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
