package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/sugaml/pharmacy/internal/patient" // Adjust to your generated protobuf package

	"google.golang.org/grpc"
)

func main() {
	// Dial the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPatientServiceClient(conn)

	// Prepare the request
	req := &pb.PatientRequest{
		Id: "123",
	}

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the GetPatient RPC
	resp, err := client.GetPatient(ctx, req)
	if err != nil {
		log.Fatalf("Error while calling GetPatient: %v", err)
	}

	// Print the response from the server
	fmt.Printf("Patient Name: %s, Age: %d\n", resp.Name, resp.Age)
}
