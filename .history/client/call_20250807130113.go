package main

import (
	"context"
	pb "grpc-calculator/proto"
	"log"
	"time"
)

func performOperations(client pb.CalculatorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Addition
	addRes, err := client.Add(ctx, &pb.CalcRequest{A: 10, B: 5})
	if err != nil {
		log.Fatalf("Could not add: %v", err)
	}
	log.Printf("Addition Result: %d", addRes.Result)

	//Subtraction
	subRes, err := client.Subtract(ctx, &pb.CalcRequest{A: 10, B: 5})
	if err != nil {
		log.Fatalf("Could not Subtract: %v", err)
	}
	log.Printf("Subtraction Result: %d", subRes.Result)

	//Division
	divRes, err := client.Divide(ctx, &pb.CalcRequest{A: 10, B: 5})
	if err != nil {
		log.Fatalf("Could not Divide: %v", err)
	}
	log.Printf("Division Result: %d", d)
}
