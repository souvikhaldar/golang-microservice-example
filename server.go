package main

import (
	"github.com/souvikhaldar/golang-microservice-example/proto/sum"
	"context"
)

type server struct{}

func (s *server) Sum(context.Context,request *sum.SumRequest,response *sum.SumResponse) error {
	response.Result = request.First + request.Second
	return nil
}

func main(){
	service := micro.NewService(
		micro.Name("adder"),
	)
}