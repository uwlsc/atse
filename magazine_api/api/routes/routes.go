package routes

import (
	v1 "magazine_api/api/routes/v1"
	"magazine_api/infrastructure"

	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	v1.Module,
	fx.Provide(NewRoutes),
)

type Routes []infrastructure.Route

// NewRoutes sets up routes
func NewRoutes(v1Routes v1.V1Routes) Routes {
	return Routes{
		v1Routes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
