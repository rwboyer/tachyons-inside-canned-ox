package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	cc := lc.ClientContext
	id := lc.Identity

	input := struct {
		ReqCC      lambdacontext.ClientContext
		ReqID      lambdacontext.CognitoIdentity
		ReqContext *lambdacontext.LambdaContext
		ReqEvents  events.APIGatewayProxyRequest
		ReqHeaders map[string]string
	}{
		cc,
		id,
		lc,
		request,
		request.Headers,
	}
	buf := bytes.NewBuffer([]byte{})
	e := json.NewEncoder(buf)
	e.Encode(input)

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "application/json"},
		MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:              buf.String(),
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
