package v1

import (
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewV1Routes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewEmployeeRoutes),
	fx.Provide(NewUploadRoutes),
	fx.Provide(NewStoryRoutes),
	fx.Provide(NewAdvertRoutes),
	fx.Provide(NewContentRoutes),
	fx.Provide(NewMagazineIssueRoutes),
	fx.Provide(NewMagazineRoutes),
	fx.Provide(NewPhotoRoutes),
)

type V1Routes struct {
	logger  lib.Logger
	handler infrastructure.Router
	routes  []infrastructure.SubRoute
}

func NewV1Routes(
	logger lib.Logger,
	handler infrastructure.Router,
	user_routes UserRoutes,
	uploadRoutes UploadRoutes,
	employee_routes EmployeeRoutes,
	story_routes StoryRoutes,
	advert_routes AdvertRoutes,
	content_routes ContentRoutes,
	issue_routes MagazineIssueRoutes,
	magazine_routes MagazineRoutes,
	photo_routes PhotoRoutes,
) V1Routes {
	return V1Routes{
		handler: handler,
		logger:  logger,
		routes: []infrastructure.SubRoute{
			user_routes,
			uploadRoutes,
			employee_routes,
			story_routes,
			advert_routes,
			content_routes,
			issue_routes,
			magazine_routes,
			photo_routes,
		},
	}
}

func (s V1Routes) Setup() {
	s.logger.Info("Setting up v1 api routes")
	api := s.handler.Group("/api/v1")
	for _, v := range s.routes {
		v.Setup(api)
	}
}
