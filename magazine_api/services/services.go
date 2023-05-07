package services

import (
	"go.uber.org/fx"
)

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewCognitoAuthService),
	fx.Provide(NewS3BucketService),
	fx.Provide(NewUserService),
	fx.Provide(NewUserProfileService),
	fx.Provide(NewTransactionService),
	fx.Provide(NewStoryService),
	fx.Provide(NewAdvertService),
	fx.Provide(NewContentService),
	fx.Provide(NewMagazineIssueService),
	fx.Provide(NewMagazineService),
	fx.Provide(NewPhotoService),
)
