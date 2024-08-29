package main

import (
	"fmt"
	"log"

	"github.com/Athlevo/Api_booking_Athlevo/api"
	"github.com/Athlevo/Api_booking_Athlevo/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.Load()

	// gRPC connection to the health service
	healthGrpcConn, err := grpc.NewClient(
		cfg.BookingSvcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to health service: %v", err)
	}
	defer healthGrpcConn.Close()

	// Create router
	router := api.NewRouter(healthGrpcConn, &cfg)

	// Start server
	fmt.Printf("API Gateway server listening on port %s\n", cfg.HTTPPort)
	if err := router.Run(cfg.HTTPPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
