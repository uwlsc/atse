package handlers

import (
	"magazine_api/api/serializers/requests"
	"magazine_api/lib"
	"magazine_api/orchestrators"
	"magazine_api/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	logger            lib.Logger
	service           services.UserService
	user_orchestrator orchestrators.UserOrchestrator
}

func NewUserHandler(logger lib.Logger, service services.UserService, u orchestrators.UserOrchestrator) UserHandler {
	return UserHandler{logger: logger, service: service, user_orchestrator: u}
}

// CreateUser godoc
// @Summary      Create User
// @Description  It creates an normal user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      requests.CreateUser  true  "Add user"
// @Success      200   {object}  object{data=requests.CreateUser}
// @Router       /user [post]
//
// Creates user
func (u UserHandler) CreateUser(c *gin.Context) {
	var user requests.CreateUser

	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(u.logger, c, err)
		return
	}

	users, err := u.user_orchestrator.CreateUser(user)
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": users})
}
