package rpc

import (
	"context"

	"github.com/RahulMj21/go-hex/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required parameters")
	}

	answer, err := grpca.api.GetAddition(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "something went wrong")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required parameters")
	}

	answer, err := grpca.api.GetSubtraction(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "something went wrong")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required parameters")
	}

	answer, err := grpca.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "something went wrong")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required parameters")
	}

	answer, err := grpca.api.GetDivision(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "something went wrong")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}
