package main

import (
	"helloworld/handler"
	pb "helloworld/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("helloworld"),
	)

	// Register handler
	pb.RegisterHelloworldHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
