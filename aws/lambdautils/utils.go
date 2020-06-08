// Utils for making AWS Lambda & API Gateway usage easier
package lambdautils

import (
	"context"
	"os"
)

// AWS Lambda doesn't support giving argv, so we use an ugly hack to detect it
func InLambda() bool {
	// https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html
	return os.Getenv("LAMBDA_TASK_ROOT") != ""
}

// simple function as a Lambda handler by discarding input payload and returning empty
// payload. this is available in AWS's lambda.Start(handler interface{}) but it's not typesafe
// (runtime error if fn sig doesn't match)
type NoPayloadAdapter func(context.Context) error

func (s NoPayloadAdapter) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	return []byte{}, s(ctx)
}
