package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ktoyod/go-lambda-sample/app/go/pkg/greet"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
}

type MyResponse struct {
	Message string `json:"Message"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: greet.Hello(event.Name)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
