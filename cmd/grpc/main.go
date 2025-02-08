package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/musadhiek-p/url-shortener/api"
	"github.com/musadhiek-p/url-shortener/internal/shortener"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterShortenerServiceServer(grpcServer, &shortener.ShortenerService{})
	fmt.Println("grpc server running on port 50051..")
	if err := grpc.Server(listener); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}
