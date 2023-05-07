package lib

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type SignedURL string

// UnmarshalJSON -> convert from json string
func (s *SignedURL) UnmarshalJSON(by []byte) error {
	str := ""
	_ = json.Unmarshal(by, &str)
	*s = SignedURL(str)
	return nil
}

// MarshalJSON -> convert to json string
func (s SignedURL) MarshalJSON() ([]byte, error) {
	signedURL, err := s.getObjectSignedURL()
	if err != nil {
		return []byte("\"\""), nil
	}

	str := "\"" + signedURL + "\""
	return []byte(str), nil
}

// GetObjectSignedURL -> gets the signed url for the stored object
// it is done this way to prevent import cycles from hapenning
func (s SignedURL) getObjectSignedURL() (string, error) {
	env := GetEnv()
	bucketName := env.S3BucketName
	presignedClient := GetPresignedClient()

	key := string(s)

	input := &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &key,
	}
	resp, err := presignedClient.PresignGetObject(context.Background(), input)
	if err != nil {
		GetLogger().Error("error-generating-signed-url", err.Error(), "key-", key)
		return "", nil
	}
	return resp.URL, nil
}

// GetPublicBucketURL get the public bucket url in case the
// bucket responses are public
func (s SignedURL) GetPublicBucketURL() string {
	env := GetEnv()
	return PublicBucketURL(env.S3BucketName, env.AWSRegion, string(s))
}

// PublicBucketURL public bucket url generate from key
func PublicBucketURL(bucketName, awsRegion, key string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, awsRegion, key)
}
