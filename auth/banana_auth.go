package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/config"
	"net/http"
)

type secret string

func BananaAuthMiddleware(yaml *config.YAML) gin.HandlerFunc {
	var secret secret
	err := yaml.Get("secret").Populate(&secret)
	if err != nil {
		panic(err)
	}

	return func(context *gin.Context) {
		isHealthCheck := context.Request.URL.Path == "/healthcheck"
		secretIsCorrect := context.Query("secret") == string(secret)

		if !isHealthCheck && !secretIsCorrect {
			_ = context.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Authorization Required"))
		}

		context.Next()
	}
}
