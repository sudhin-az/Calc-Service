package server

import "context"

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	result := req.A + req.B
	return &pb.CalcResponse{Result: result}, nil
}

func (s *)  {
	
}
