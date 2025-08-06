package server

import "context"

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	result := req.A + req.func Benchmark(b *testing.B) {
		for i := 0; i < b.N; i++ {
			
		}
	}
}
