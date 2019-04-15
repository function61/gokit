// Because AWS SDK has so much ceremony..
package s3facade

// why the facade? there's so many things happening to set up the client that I always
// struggle finding the correct AWS APIs, when usually all I want to do is the same..

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Client(akid string, secret string, regionId string) (*s3.S3, error) {
	awsSession, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	manualCredential := credentials.NewStaticCredentials(
		akid,
		secret,
		"")

	s3Client := s3.New(
		awsSession,
		aws.NewConfig().WithCredentials(manualCredential).WithRegion(regionId))

	return s3Client, nil
}
