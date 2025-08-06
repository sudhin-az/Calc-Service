package server

import "context"

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, err) {

}
