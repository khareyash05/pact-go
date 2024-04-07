package grpc

import (
	"context"
	"fmt"
	order "github.com/GowthamGirithar/contract-testing-demo/order-service-provider/internal/adapter/grpc"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-provider/internal/infrastructure/grpc/middleware"
	pb "github.com/GowthamGirithar/contract-testing-demo/proto/order-service"
	gmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	gctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func RunServer(ctx context.Context, port string) error {
	server := grpc.NewServer(grpc.UnaryInterceptor(
		gmiddleware.ChainUnaryServer(
			gctxtags.UnaryServerInterceptor(gctxtags.WithFieldExtractor(gctxtags.CodeGenRequestFieldExtractor)),
			middleware.LogRequest,
		),
	))
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("error in listening")
	}

	orderServer := order.NewServer()
	pb.RegisterOrderServer(server, orderServer)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for range c {
			fmt.Print("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	return server.Serve(listen)
}
