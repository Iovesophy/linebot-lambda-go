package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	Successful = 200
	ErrSsm     = 500
	ErrReq     = 500
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	line := Line{}
	err := line.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: ErrSsm}, errors.New("ssmerror")
	}

	eve, err := ParseRequest(line.ChannelSecret, req)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: ErrReq}, errors.New("badrequest")
	}
	line.EventRouter(eve)
	return events.APIGatewayProxyResponse{Body: req.Body, StatusCode: Successful}, nil
}

func ParseRequest(channelSecret string, r events.APIGatewayProxyRequest) ([]*linebot.Event, error) {
	req := &struct {
		Events []*linebot.Event `json:"events"`
	}{}
	if err := json.Unmarshal([]byte(r.Body), req); err != nil {
		return nil, err
	}
	return req.Events, nil
}

func main() {
	lambda.Start(Handler)
}
