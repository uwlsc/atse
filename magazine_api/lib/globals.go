package lib

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	_presignClient *s3.PresignClient
)

func GetPresignedClient() *s3.PresignClient {
	return _presignClient
}

func RegisterGlobalInfrastructure(
	log Logger,
	s3Client *s3.Client,
	presignClient *s3.PresignClient,
	awsCfg aws.Config,
) {
	log.Info("registering global presigned client")
	_presignClient = presignClient
}
