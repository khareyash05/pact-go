package grpc

import (
	"context"
	order "github.com/GowthamGirithar/contract-testing-demo/proto/order-service"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	order.UnimplementedOrderServer
}

func NewServer() server {
	return server{}
}

func (s server) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	err := validate(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &order.CreateOrderResponse{
		OrderNumber: req.GetOrderNumber(),
	}, nil
}

func validate(req *order.CreateOrderRequest) error {
	if req.GetCustomerEmail() == "" {
		return errors.New("Invalid email format")
	}
	return nil
}

func (s server) GetOrder(ctx context.Context, req *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	return &order.GetOrderResponse{
		OrderNumber:   req.GetOrderNumber(),
		CustomerEmail: "test@gmail.com",
	}, nil
}
