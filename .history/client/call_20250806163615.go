package client

import (
	"context"
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
}