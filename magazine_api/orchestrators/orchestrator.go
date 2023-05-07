package orchestrators

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewUserOrchestrator),
	fx.Provide(NewUserProfileOrchestrator),
	fx.Provide(NewEmployeeProfileOrchestrator),
)
