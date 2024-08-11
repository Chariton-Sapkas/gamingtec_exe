package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	pb "gamingtec_exe/api/proto"
	"gamingtec_exe/service"
	store "gamingtec_exe/storage"
)

func main() {
	var (
		serverPort = ":5050"
	)
	log.Println("Attempting to start server now!")

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	// set user store
	userStore := store.NewUserStore()
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, service.NewUserServiceServer(userStore))

	// Register the Health Service
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	// Set the health status to SERVING
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	// Enable reflection for better tool support
	reflection.Register(grpcServer)

	// Initiate server on given port
	log.Println("gRPC server listening on " + serverPort)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
