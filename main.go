package certificateS3

import (
	"golang.org/x/crypto/ssh"
)

// GetCertificate will return a certificate signer from a private key in a S3 Bucket
func GetCertificate(credentials Credentials, file File) (ssh.Signer, error) {
	s3Service := connectToS3Service(credentials)
	body, err := getContentFromS3(s3Service, file)
	if err != nil {
		return nil, err
	}

	content, err := getContentBytes(body)
	if err != nil {
		return nil, err
	}

	// Create the Signer for this private key
	signer, err := ssh.ParsePrivateKey(content)
	if err != nil {
		return nil, err
	}

	return signer, nil
}
