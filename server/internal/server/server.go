package server

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/wcygan/scylla-on-k8s/generated/go/ping/v1"
	"log"
	"time"
)

// PingService is used to implement ping.PingService.
type PingService struct {
	pb.UnimplementedPingServiceServer
}

// NewServer creates a new server.
func NewServer() *PingService {
	return &PingService{}
}

// Ping implements the Ping method from the PingService.
func (s *PingService) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	timestamp := time.Now().Unix()
	token := uuid.New().String()
	log.Printf("%s  ---  %d", token, timestamp)
	return &pb.PingResponse{}, nil
}
