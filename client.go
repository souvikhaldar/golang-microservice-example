package main

import (
	"context"
	"fmt"
	"log"

	micro "github.com/micro/go-micro"
	"github.com/souvikhaldar/golang-microservice-example/sum"
)

func main() {
	service := micro.NewService(micro.Name("adder.client"))

	service.Init()

	adderClient := sum.NewAdderService("adder", service.Client())

	response, err := adderClient.Sum(context.TODO(), &sum.SumRequest{First: 1, Second: 2})
	if err != nil {
		log.Println("Err in calling the server: ", err)
		return
	}
	fmt.Println("Result of summation: ", response.Result)
}
