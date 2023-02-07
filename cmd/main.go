package main

import (
	"GoHexArchTutorial/internal/adapters/app/api"
	"GoHexArchTutorial/internal/adapters/core/arithmetic"
	gRPC "GoHexArchTutorial/internal/adapters/framework/left/grpc"
	"GoHexArchTutorial/internal/adapters/framework/right/db"
	"GoHexArchTutorial/internal/ports"
	"log"
	"os"
)

func main() {
	var err error

	//ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort
	// end of ports section
	//Adapters section
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")
	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiale dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()
	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)
	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	// End of Adapters section

	// Running Application
	gRPCAdapter.Run()
}
