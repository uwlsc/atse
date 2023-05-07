package bootstrap

import (
	"context"
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/api/routes"
	"magazine_api/cmd"
	"magazine_api/component"
	"magazine_api/infrastructure"
	"magazine_api/lib"
	"magazine_api/orchestrators"

	"magazine_api/services"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	routes.Module,
	services.Module,
	component.Module,
	infrastructure.Module,
	orchestrators.Module,
	middlewares.Module,
	cmd.Module,
	lib.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	middlewares middlewares.Middlewares,
	env lib.Env,
	router infrastructure.Router,
	routes routes.Routes,
	logger lib.Logger,
	database infrastructure.Database,
	rootCmd cmd.RootCommand,
	migration infrastructure.Migrations,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				rootCmd.Run = func(cmd *cobra.Command, args []string) {
					logger.Info(`+------+------+`)
					logger.Info(`| Magazine API |`)
					logger.Info(`+------+------+`)
					// migration.Migrate()
					middlewares.Setup()
					routes.Setup()
					if env.ServerPort == "" {
						router.Run()
					} else {
						router.Run(":" + env.ServerPort)
					}
				}
				go rootCmd.Execute()
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("Stopping Application")
				conn := database.Pool
				conn.Close()
				return nil
			},
		},
	)
}
