package grpc

import (
	"GoHexArchTutorial/internal/adapters/framework/left/grpc/pb"
	"GoHexArchTutorial/internal/ports"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	arithServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server gRPC server over port 9000: %v", err)

	}
}
