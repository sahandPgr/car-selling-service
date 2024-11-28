package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sahandPgr/car-selling-service/api/helper"
	"github.com/sahandPgr/car-selling-service/config"
	"github.com/sahandPgr/car-selling-service/constatns"
	jwt_service "github.com/sahandPgr/car-selling-service/internal/services/jwt"
	serviceerrors "github.com/sahandPgr/car-selling-service/internal/services/service_errors"
)

// Authentication is a middleware for authentication
func Authentication(config *config.Config) gin.HandlerFunc {
	tokenService := jwt_service.NewJwtService(config)
	return func(ctx *gin.Context) {
		var err error
		tokenClaims := make(map[string]interface{})
		authVal := ctx.GetHeader(constatns.AuthorizationHeaderKey)
		token := strings.Split(authVal, " ")[1]
		if authVal == "" {
			err = &serviceerrors.ServiceError{EndUserMessage: serviceerrors.TokenRequired}
		} else {
			tokenClaims, err = tokenService.GetClaims(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &serviceerrors.ServiceError{EndUserMessage: serviceerrors.TokenExpired}
				default:
					err = &serviceerrors.ServiceError{EndUserMessage: serviceerrors.TokenInvalid}
				}
			}
		}

		if err != nil {
			ctx.AbortWithStatusJSON(helper.Unauthorized, helper.GetBaseHttpResponseWithError(nil, false, helper.Unauthorized, err))
			return
		}
		ctx.Set(constatns.UserIdKey, tokenClaims[constatns.UserIdKey])
		ctx.Set(constatns.FirstName, tokenClaims[constatns.FirstName])
		ctx.Set(constatns.LastName, tokenClaims[constatns.LastName])
		ctx.Set(constatns.Email, tokenClaims[constatns.Email])
		ctx.Set(constatns.PhoneNumber, tokenClaims[constatns.PhoneNumber])
		ctx.Set(constatns.Roles, tokenClaims[constatns.Roles])
		ctx.Set(constatns.ExpireTimeKey, tokenClaims[constatns.ExpireTimeKey])

		ctx.Next()
	}
}

// Authorization is a middleware for authorization
func Authorization(validRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Keys) == 0 {
			ctx.AbortWithStatusJSON(helper.SatusForbidden, helper.GetBaseHttpResponse(nil, false, helper.SatusForbidden))
			return
		}
		rolesVal := ctx.Keys[constatns.Roles]
		if rolesVal == nil {
			ctx.AbortWithStatusJSON(helper.SatusForbidden, helper.GetBaseHttpResponse(nil, false, helper.SatusForbidden))
			return
		}
		roles := rolesVal.([]string)
		rolesMap := make(map[string]int)
		for _, role := range roles {
			rolesMap[role] = 0
		}
		for _, item := range validRoles {
			if _, ok := rolesMap[item]; ok {
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(helper.SatusForbidden, helper.GetBaseHttpResponse(nil, false, helper.SatusForbidden))
	}
}
