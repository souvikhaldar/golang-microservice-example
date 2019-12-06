package main

import (
	"context"
	micro "github.com/micro/go-micro"
	"github.com/souvikhaldar/golang-microservice-example/sum"
	"log"
)

type Server struct{}

func (s *Server) Sum(ctx context.Context, request *sum.SumRequest, response *sum.SumResponse) error {
	response.Result = request.First + request.Second
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("adder"),
	)
	service.Init()
	sum.RegisterAdderHandler(service.Server(), new(Server))
	log.Fatal(service.Run())

}
