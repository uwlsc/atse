package middlewares

import (
	"magazine_api/api/serializers/responses"
	"magazine_api/constants"
	"magazine_api/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwt"
)

type CognitoAuthMiddleware struct {
	service services.CognitoAuthService
}

func NewCognitoAuthMiddleware(service services.CognitoAuthService) CognitoAuthMiddleware {
	return CognitoAuthMiddleware{service: service}
}

func (m CognitoAuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := m.getTokenFromHeader(c)

		if err != nil {
			responses.ErrorJSON(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set(constants.Claims, token.PrivateClaims())
		c.Set(constants.UID, token.PrivateClaims()["username"])

		c.Next()
	}
}

func (m CognitoAuthMiddleware) getTokenFromHeader(c *gin.Context) (jwt.Token, error) {
	header := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	token, err := m.service.VerifyToken(idToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}
