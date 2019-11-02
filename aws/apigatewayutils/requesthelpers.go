package apigatewayutils

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func RequiredQueryParam(key string, req *events.APIGatewayProxyRequest) (string, error) {
	val, found := req.QueryStringParameters[key]
	if !found {
		return "", fmt.Errorf("missing param: %s", key)
	}

	return val, nil
}
