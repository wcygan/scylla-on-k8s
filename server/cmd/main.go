package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	pb "github.com/wcygan/scylla-on-k8s/generated/go/ping/v1"
	"google.golang.org/grpc"
)

// server is used to implement the PingService.
type server struct {
	pb.UnimplementedPingServiceServer
}

// Ping implements the Ping method from the PingService.
func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	timestamp := time.Now().Unix()
	token := uuid.New().String()
	log.Printf("%s  ---  %d", token, timestamp)
	return &pb.PingResponse{}, nil
}

func main() {
	var port string

	var rootCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the gRPC server",
		Run: func(cmd *cobra.Command, args []string) {
			lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()
			pb.RegisterPingServiceServer(s, &server{})
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		},
	}

	rootCmd.Flags().StringVarP(&port, "port", "p", "50051", "Port to run the server on")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error while running the server: %v", err)
	}
}
