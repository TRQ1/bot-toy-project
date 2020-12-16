package main

import (
	"context"
	"fmt"
	"os"

	model "github.com/TRQ1/bot-toy-project/ddScreenshot/core/model"
	tools "github.com/TRQ1/bot-toy-project/ddScreenshot/tools"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	// for dev
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("Error loading .env file, %v", err))
		}
		ky := model.Key{Dd_api: "", Dd_app_api: ""}
		dd := model.Graphs{Start: 0, End: 0, Metric_name: "", Query: ""}
		tools.CurlDatadog(dd, ky)
	} else {
		lambda.Start(handleRequest)
	}
}