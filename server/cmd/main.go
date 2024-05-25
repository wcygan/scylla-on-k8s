package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/spf13/cobra"
	pb "github.com/wcygan/scylla-on-k8s/generated/go/ping/v1"
	"github.com/wcygan/scylla-on-k8s/server/internal/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("Error while running the server: %v", err)
	}
}

var RootCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		cluster := gocql.NewCluster(scyllaAddress)
		session, err := cluster.CreateSession()
		if err != nil {
			log.Fatalf("failed to create session: %v", err)
		}
		defer session.Close()

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pingService := server.NewPingService(session)
		pb.RegisterPingServiceServer(s, pingService)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

var port string
var scyllaAddress string

func init() {
	RootCmd.Flags().StringVarP(&port, "port", "p", "50051", "Port to run the server on")
	RootCmd.Flags().StringVarP(&scyllaAddress, "scylla", "s", "scylla-db", "Address to connect to ScyllaDB")
}
