package main

import (
	"context"
	pb "grpc-calculator/proto"
)

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	result := req.A + req.B
	return &pb.CalcResponse{Result: result}, nil
}

func (s *CalculatorServer) Subtract(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	result := req.A - req.B
	return &pb.CalcResponse{Result: result}, nil
}

func (s *CalculatorServer) Divide(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	result := req.A / req.B
	return  &p
}
