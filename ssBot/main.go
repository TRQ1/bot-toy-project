package main

import (
	"context"
	"fmt"
	"os"

	model "github.com/TRQ1/bot-toy-project/ssBot/core/model"
	tools "github.com/TRQ1/bot-toy-project/ssBot/tools"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	imageName := request.QueryStringParameters["imageName"]
	service := request.QueryStringParameters["servie"]
	board := request.QueryStringParameters["board"]
	user := request.QueryStringParameters["user"]

	target := model.UploadInfo{ImageName: imageName, Service: service, Board: board, User: user}

	tools.DownloadS3(imageName)
	tools.UploadImage(target)

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	// for dev
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("Error loading .env file, %v", err))
		}

		target := model.UploadInfo{ImageName: "newrelic.png", Service: "newrelic", Board: "mkapi"}

		tools.DownloadS3("newrelic.png")
		tools.UploadImage(target)
	} else {
		lambda.Start(handleRequest)
	}

}
