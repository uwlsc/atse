package services

import (
	"context"
	"io"
	"magazine_api/lib"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3BucketService bucket storage for aws
type S3BucketService struct {
	*s3.Client
	uploader *manager.Uploader
	presign  *s3.PresignClient
	logger   lib.Logger
	env      lib.Env
}

func NewS3BucketService(
	logger lib.Logger,
	client *s3.Client,
	env lib.Env,
	presign *s3.PresignClient,
	uploader *manager.Uploader,
) S3BucketService {
	return S3BucketService{
		logger:   logger,
		Client:   client,
		env:      env,
		presign:  presign,
		uploader: uploader,
	}
}

// UploadFile uploads the file
func (s S3BucketService) UploadFile(
	ctx context.Context,
	file io.Reader,
	key string,
) (*manager.UploadOutput, error) {
	input := &s3.PutObjectInput{
		Bucket: &s.env.S3BucketName,
		Key:    &key,
		Body:   file,
	}

	return s.uploader.Upload(ctx, input)
}

// GetSignedURL get the signed url for file
func (s S3BucketService) GetSignedURL(
	ctx context.Context,
	key string,
) string {
	expires := time.Now().Add(time.Minute)
	input := &s3.GetObjectInput{
		Bucket:          &s.env.S3BucketName,
		Key:             &key,
		ResponseExpires: &expires,
	}
	resp, err := s.presign.PresignGetObject(ctx, input)
	if err != nil {
		s.logger.Error("error-generating-presigned-url", err.Error())
		return ""
	}

	return resp.URL
}
