package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Printf("Start lambda")
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		log.Printf("Processing Lambda request %v\n", request.RequestContext)
		return events.APIGatewayProxyResponse{
			Body:       "💩",
			StatusCode: 200,
		}, nil
	})
}
