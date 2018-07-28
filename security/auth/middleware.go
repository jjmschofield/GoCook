package auth

import "github.com/gin-gonic/gin"

func IsAuthenticatedMiddleware(context *gin.Context){
	jwtToken, tokenError := getJwtToken(context)

	if(tokenError != nil){
		context.AbortWithError(401, tokenError)
		return;
	}

	validStandardClaims, validationError := hasValidClaims(jwtToken)

	if(!validStandardClaims){
		context.AbortWithError(401, validationError)
		return;
	}

	context.Next()
}
