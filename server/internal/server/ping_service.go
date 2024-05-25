package server

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	pb "github.com/wcygan/scylla-on-k8s/generated/go/ping/v1"
	"log"
	"time"
)

// PingService is used to implement ping.PingService.
type PingService struct {
	pb.UnimplementedPingServiceServer
	session *gocql.Session
}

// NewPingService creates a new server.
func NewPingService(session *gocql.Session) *PingService {
	return &PingService{session: session}
}

// Ping implements the Ping method from the PingService.
func (s *PingService) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	timestamp := time.Now().UnixMilli()
	token := uuid.New().String()
	log.Printf("%s  ---  %d", token, timestamp)

	if err := s.session.Query(`INSERT INTO events (id, timestamp) VALUES (?, ?)`, token, timestamp).Exec(); err != nil {
		log.Printf("failed to insert event: %v", err)
	}

	return &pb.PingResponse{}, nil
}
