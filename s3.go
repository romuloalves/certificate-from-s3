package certificateS3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func connectToS3Service(creds Credentials) *s3.S3 {
	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(creds.Region),
	}))

	s3Credentials := credentials.NewStaticCredentialsFromCreds(credentials.Value{
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
	})

	return s3.New(session, &aws.Config{
		Credentials: s3Credentials,
	})
}

func getContentFromS3(s3Service *s3.S3, file File) (io.ReadCloser, error) {
	ctx := context.Background()

	result, err := s3Service.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(file.Bucket),
		Key:    aws.String(file.Key),
	})

	if err != nil {
		// Cast err to awserr.Error to handle specific error codes.
		aerr, ok := err.(awserr.Error)
		if ok && aerr.Code() == s3.ErrCodeNoSuchKey {
			// Specific error code handling
		}
		return nil, err
	}

	// Make sure to close the body when done with it for S3 GetObject APIs or
	// will leak connections.
	defer result.Body.Close()

	return result.Body, nil
}
