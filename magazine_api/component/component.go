package component

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewUserComponent),
	fx.Provide(NewUserProfileComponent),
	fx.Provide(NewTransactionComponent),
	fx.Provide(NewStoryComp),
	fx.Provide(NewAdMgmtComp),
	fx.Provide(NewContentComp),
	fx.Provide(NewIssueComp),
	fx.Provide(NewMagazineComp),
	fx.Provide(NewPhotographComp),
)
