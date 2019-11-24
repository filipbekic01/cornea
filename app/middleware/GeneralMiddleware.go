package middleware

import (
	"github.com/kataras/iris/v12"
)

// GeneralMiddleware .
func GeneralMiddleware(ctx iris.Context) {
	ctx.Next()
}
