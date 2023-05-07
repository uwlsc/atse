package infrastructure

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewDatabase),
	fx.Provide(NewRouter),
	fx.Provide(NewMigrations),
	fx.Provide(NewAWSConfig),
	fx.Provide(NewS3Client),
	fx.Provide(NewCognitoClient),
	fx.Provide(NewPresignClient),
	fx.Provide(NewS3Uploader),
)
