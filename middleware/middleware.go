package middleware

import (
	"fmt"
	"net/http"
	"project-pertama/model"
	"project-pertama/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func LogMiddleware(ctx *gin.Context) {
	fmt.Println("hello ini dari middleware")
	ctx.Next()
}

func AuthMiddleware(ctx *gin.Context) {
	authorizationValue := ctx.GetHeader("Authorization")
	splittedValue := strings.Split(authorizationValue, "Bearer ")
	if len(splittedValue) == 0 {
		var r model.Response = model.Response{
			Success: false,
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}
	jwtToken := splittedValue[1]

	claims, err := util.GetJWTClaims(jwtToken)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	ctx.Set("claims", claims)

	ctx.Next()
}

func AdminMiddleware(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		var r model.Response = model.Response{
			Success: false,
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	mapClaims, ok := claims.(map[string]any)
	if !ok {
		var r model.Response = model.Response{
			Success: false,
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	isAdmin, ok := mapClaims["admin"]
	if !ok {
		var r model.Response = model.Response{
			Success: false,
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	boolIsAdmin, ok := isAdmin.(bool)
	if !ok {
		var r model.Response = model.Response{
			Success: false,
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	if !boolIsAdmin {
		var r model.Response = model.Response{
			Success: false,
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	ctx.Next()
}
