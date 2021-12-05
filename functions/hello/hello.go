package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type Identity struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

type User struct {
	AppMetaData  *AppMetaData  `json:"app_metadata"`
	Email        string        `json:"email"`
	Exp          int           `json:"exp"`
	Sub          string        `json:"sub"`
	UserMetadata *UserMetadata `json:"user_metadata"`
}

type AppMetaData struct {
	Provider string `json:"provider"`
}
type UserMetadata struct {
	FullName string `json:"full_name"`
}

type Bearer struct {
	Identity *Identity `json:"identity"`
	User     *User     `json:"user"`
	SiteUrl  string    `json:"site_url"`
	Alg      string    `json:"alg"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	cc := lc.ClientContext
	id := lc.Identity

	bearer := lc.ClientContext.Custom["netlify"]
	raw, err := base64.StdEncoding.DecodeString(bearer)
	if err != nil {
		log.Print(err)
	}
	log.Print(raw)
	data := Bearer{}
	var rj map[string]interface{}
	err = json.Unmarshal(raw, &rj)

	if err != nil {
		log.Print(err)
	}
	err = json.Unmarshal(raw, &data)

	if err != nil {
		log.Print(err)
	}

	input := struct {
		ReqCustom  map[string]interface{}
		ReqCC      lambdacontext.ClientContext
		ReqID      lambdacontext.CognitoIdentity
		ReqContext *lambdacontext.LambdaContext
		ReqEvents  events.APIGatewayProxyRequest
		ReqHeaders map[string]string
	}{
		rj,
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
