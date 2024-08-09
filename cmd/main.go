package main

import (
	"log"
	"net"
	"net/http"

	pb "gamingtec_exe/api/proto"
	store "gamingtec_exe/storage"

	"gamingtec_exe/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	userStore := store.NewUserStore()
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, service.NewUserServiceServer(userStore))

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	http.HandleFunc("/health", healthCheckHandler)
	go func() {
		log.Println("Health check server listening on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("failed to serve health check: %v", err)
		}
	}()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}
