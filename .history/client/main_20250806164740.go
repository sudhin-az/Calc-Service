package client

import (
	pb "grpc-calculator/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not Connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)
	performOperations(client)
}
