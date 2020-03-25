package lambdautils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/apex/gateway"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// higher level handler adapter. ServeApiGatewayProxyRequestUsingHttpHandler() is lower if
// you need to be able to give APIGatewayProxyRequest and get back byte slice & error
type LambdaHttpHandlerAdapter struct {
	httpHandler http.Handler
}

// adapts lambda.Handler to Go's http.Handler by translating between API gateway request/responses
func NewLambdaHttpHandlerAdapter(httpHandler http.Handler) lambda.Handler {
	return &LambdaHttpHandlerAdapter{httpHandler}
}

func (m *LambdaHttpHandlerAdapter) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	e := &events.APIGatewayProxyRequest{}
	if err := json.Unmarshal(payload, e); err != nil {
		return nil, fmt.Errorf("APIGatewayProxyRequest unmarshal: %w", err)
	}

	return ServeApiGatewayProxyRequestUsingHttpHandler(
		ctx,
		e,
		m.httpHandler)
}

// github.com/akrylysov/algnhsa has similar implementation than apex/gateway, but had the
// useful bits non-exported and it used httptest for production code
func ServeApiGatewayProxyRequestUsingHttpHandler(
	ctx context.Context,
	proxyRequest *events.APIGatewayProxyRequest,
	httpHandler http.Handler,
) ([]byte, error) {
	request, err := gateway.NewRequest(ctx, *proxyRequest)
	if err != nil {
		return nil, err
	}

	response := gateway.NewResponse()

	httpHandler.ServeHTTP(response, request)

	proxyResponse := response.End()

	return json.Marshal(&proxyResponse)
}
