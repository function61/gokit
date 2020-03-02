// Utils for making AWS Lambda & API Gateway usage easier
package lambdautils

import (
	"os"
)

// AWS Lambda doesn't support giving argv, so we use an ugly hack to detect it
func InLambda() bool {
	// https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html
	return os.Getenv("LAMBDA_TASK_ROOT") != ""
}
