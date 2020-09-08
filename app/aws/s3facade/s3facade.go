// Because AWS SDK has so much ceremony..
package s3facade

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/function61/gokit/os/osutil"
)

// this abstraction so we support immediate validation of credentials, unlike AWS's API
type CredentialsObtainer func() (*credentials.Credentials, error)

// giving nil "credentials" is supported, if you want the automatic procedure
func Bucket(
	bucket string,
	obtainCredentials CredentialsObtainer,
	region string,
) (*BucketContext, error) {
	config := aws.NewConfig().WithRegion(region)

	if obtainCredentials != nil {
		credentials, err := obtainCredentials()
		if err != nil {
			return nil, err
		}

		config.WithCredentials(credentials)
	}

	return BucketWithConfig(bucket, config)
}

func BucketWithConfig(bucket string, config *aws.Config) (*BucketContext, error) {
	awsSession, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	return &BucketContext{
		Name: &bucket,
		S3: s3.New(
			awsSession,
			config),
	}, nil
}

type BucketContext struct {
	Name *string // pointer because all S3 operations take it as such
	S3   *s3.S3
}

// pretty much same as AWS's credentials.NewEnvCredentials(), except this fails fast instead
// of when S3 operations are first called
func CredentialsFromEnv() (*credentials.Credentials, error) {
	accessKeyId, err := osutil.GetenvRequired("AWS_ACCESS_KEY_ID")
	if err != nil {
		return nil, err
	}

	accessKeySecret, err := osutil.GetenvRequired("AWS_SECRET_ACCESS_KEY")
	if err != nil {
		return nil, err
	}

	return credentials.NewStaticCredentials(accessKeyId, accessKeySecret, ""), nil
}

// adapter for making from AWS's Credentials
func Credentials(creds *credentials.Credentials) CredentialsObtainer {
	return func() (*credentials.Credentials, error) {
		return creds, nil
	}
}
