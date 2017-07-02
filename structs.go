package certificateS3

// Credentials represents a credentials information to auth S3 Bucket
type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

// File represents an object from S3 Bucket
type File struct {
	Bucket string
	Key    string
}
