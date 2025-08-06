package client

import (
	"context"
	"time"
)

func performOperations(client pb.CalculatorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addRes, err := client.Add(ctx, &pb.CalcRequest{})
}