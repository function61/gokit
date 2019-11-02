package apigatewayutils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func Redirect(to string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusFound,
		Headers: map[string]string{
			"Location": to,
		},
		Body: fmt.Sprintf("Redirecting to %s", to),
	}, nil
}

func NotFound(msg string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       msg,
	}, nil
}

func NoContent() (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
	}, nil
}

func Created() (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
	}, nil
}

func BadRequest(msg string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       msg,
	}, nil
}

func InternalServerError(msg string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       msg,
	}, nil
}

func RespondJson(out interface{}) (*events.APIGatewayProxyResponse, error) {
	asJson, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(asJson),
	}, nil
}

func OkText(msg string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		Body: string(msg),
	}, nil
}

func Binary(body []byte, contentType string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": contentType,
		},
		Body:            base64.StdEncoding.EncodeToString(body),
		IsBase64Encoded: true,
	}, nil
}
