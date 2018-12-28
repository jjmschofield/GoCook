package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/logger"
	"go.uber.org/zap"
)

func AuthenticationMiddleware(context *gin.Context) {
	jwtToken, tokenError := getJwtToken(context)

	if tokenError != nil {
		logger.Warn("Provided token is invalid", zap.Error(tokenError))
		context.AbortWithStatus(401)
		return
	}

	validStandardClaims, validationError := hasValidClaims(jwtToken)

	if !validStandardClaims {
		logger.Warn("Provided token has incorrect claims", zap.Error(validationError))
		context.AbortWithStatus(401)
		return
	}

	context.Set("token", jwtToken)
	context.Set("userId", GetClaim(jwtToken, "sub"))

	context.Next()
}
