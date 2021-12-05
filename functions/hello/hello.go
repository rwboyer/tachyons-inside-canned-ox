package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	input := struct {
		ReqContext events.APIGatewayProxyRequestContext //`json: reqcontext`
		ReqHeaders map[string]string                    //`json: reqheaders`
	}{
		request.RequestContext,
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
