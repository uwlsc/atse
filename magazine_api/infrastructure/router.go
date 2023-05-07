package infrastructure

import (
	"magazine_api/lib"
	"net/http"

	_ "magazine_api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router -> Gin Router
type Router struct {
	*gin.Engine
}

//NewRouter : all the routes are defined here
func NewRouter(env lib.Env, logger lib.Logger) Router {

	gin.DefaultWriter = logger.GetGinLogger()
	httpRouter := gin.Default()

	httpRouter.MaxMultipartMemory = env.MaxMultipartMemory

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "PDF Converter-API 📺 API Up and Running"})
	})

	httpRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return Router{
		httpRouter,
	}
}

type SubRoute interface {
	Setup(*gin.RouterGroup)
}

// Route interface
type Route interface {
	Setup()
}
