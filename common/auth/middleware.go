package auth

import "github.com/gin-gonic/gin"

func AuthenticationMiddleware(context *gin.Context) {
	jwtToken, tokenError := getJwtToken(context)

	if tokenError != nil {
		context.AbortWithError(401, tokenError)
		return
	}

	validStandardClaims, validationError := hasValidClaims(jwtToken)

	if !validStandardClaims {
		context.AbortWithError(401, validationError)
		return
	}

	context.Set("token", jwtToken)
	context.Set("userId", GetClaim(jwtToken, "sub"))

	context.Next()
}
