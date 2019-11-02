package apigatewayutils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

type adapterFn func(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)

func Adapter(fn adapterFn) *HandlerAdapter {
	return &HandlerAdapter{fn}
}

type HandlerAdapter struct {
	fn adapterFn
}

func (h *HandlerAdapter) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	var req events.APIGatewayProxyRequest
	if err := json.Unmarshal(payload, &req); err != nil {
		resp, _ := BadRequest(fmt.Sprintf("payload not valid JSON: %v", err))
		return toJson(resp)
	}

	if req.HTTPMethod == "" {
		resp, _ := BadRequest("payload not valid APIGatewayProxyRequest")
		return toJson(resp)
	}

	out, err := h.fn(ctx, req)
	if err != nil {
		return nil, err
	}

	return toJson(out)
}

func toJson(resp *events.APIGatewayProxyResponse) ([]byte, error) {
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}
