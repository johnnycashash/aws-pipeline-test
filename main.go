package main

import (
	"errors"
	"log"
  
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Request Received %s", request.RequestContext.RequestID)
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, errors.New("not welcome")
	}

	return events.APIGatewayProxyResponse{
		Body:       "Welcome to pipeline test " + request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
