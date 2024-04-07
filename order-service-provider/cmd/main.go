package main

import (
	"context"
	"github.com/GowthamGirithar/contract-testing-demo/order-service-provider/internal/infrastructure/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"log"
)

func main() {
	err := grpc.RunServer(context.Background(), ":8082")
	if err != nil {
		log.Fatalln(err)
	}

}
